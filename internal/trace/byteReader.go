package trace

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type byteReader interface {

	/**
	 * 这个接口，定义了可以读取指定的字节
	 * 至于是从哪读取的，它不管
	 */
	readBytes(len int) []byte

	/**
	 * 是否还有更多的数据
	 */
	hasMore() bool
}

/**
 * 字节读取实现，从tcpdump输出的文件中读取
 */
type tcpDumpFileByteReader struct {
	f         string        // 文件的路径
	fd        *os.File      // 文件的指针
	bufReader *bufio.Reader // 按行读取文件的reader

	buf  []byte // 每次读取512字节的数据
	spos int    // 读取到哪个字节了
	smax int    // 缓存中一共有多少字节

	end bool // 是否还有数据

	lineNum int // 当前读到文件的哪一行了
}

const MAX_BUF_LEN int = 1024

func NewTcpDumpFileByteReader(f string) *tcpDumpFileByteReader {

	reader := &tcpDumpFileByteReader{}
	reader.f = f
	reader.lineNum = 1
	reader.spos = 0
	reader.end = false
	temp := make([]byte, 0, MAX_BUF_LEN)
	reader.buf = temp

	fd, err := os.Open(f)
	if err != nil {
		panic("读取文件失败:" + f)
	}
	reader.fd = fd
	reader.bufReader = bufio.NewReader(reader.fd)

	return reader
}

func (reader *tcpDumpFileByteReader) hasMore() bool {

	return !reader.end
}

func (reader *tcpDumpFileByteReader) readBytes(len int) []byte {

	// 确保能读到足够的内容
	count := 0
	for reader.smax-reader.spos < len {

		if count >= 100 {
			//读取了100次，还没有读到足够的内容，可能确实是没有了，直接抛出异常
			reader.end = true
			break
		}
		reader.readMoreFromFile()
		count++
	}
	if reader.smax-reader.spos >= len {
		//说明够 ，直接返回
		reader.spos = reader.spos + len
		return reader.buf[reader.spos : reader.spos+len]
	}
	reader.end = true
	panic("数据不够了")
}

func (t *tcpDumpFileByteReader) close() {

	if t.fd != nil {
		t.fd.Close()
	}
}

func (reader *tcpDumpFileByteReader) readMoreFromFile() {

	// 一共读取的字节数量
	totalReadCount := 0
	//从文件里读取更多的数据 存储到缓存中去
	br := reader.bufReader
	for {
		data, _, eof := br.ReadLine()
		if eof == io.EOF {
			// 读取到文件的末尾了 这里需要直接结束了
			break
		}
		line := string(data) //这里转换有一点性能损耗
		//去除开头的空格
		line = deleteSpace(line)
		if line == "" {
			continue
		}
		// 如果是时间开头 ，那么就忽略这一行
		if timeLine(line) {
			continue
		}
		// 数据都是 0x开头 ，而且有冒号
		if tcpDumpData(line) {
			//去掉头部
			line = line[7:]
			//去掉空格
			line = deleteSpace(line)
			fmt.Println("--------：" + line)
			// 读取数据 因为没有中文，所以不用考虑编码问题
			for i := 0; i < len(line); i = i + 2 {
				if line[i] == ' ' {
					i++
					continue
				}
				numStr := line[i : i+2] //读2个16进制的数据，理论上应该不会out of index
				// 将这个16进制的字符串，转换为数字 int16 是2个字节
				var target uint16
				target = string2Uint16(numStr)
				// 转为字节 小端
				b := make([]byte, 16)
				binary.BigEndian.PutUint16(b, target)
				reader.buf = append(reader.buf, b...)
				totalReadCount = totalReadCount + 2
			}
		}
	}
	if totalReadCount <= 0 {
		reader.end = true
	}
	//
	fmt.Println("----123:" + string(reader.buf))
	reader.smax = reader.smax + totalReadCount
}

func tcpDumpData(line string) bool {

	return len(line) > 7 && line[0:2] == "0x" && line[6:7] == ":"
}

func timeLine(line string) bool {

	return len(line) > 8 && line[2] == ':' && line[5] == ':'
}

func deleteSpace(line string) string {

	start := 0
	for _, ch := range line {

		if ch != ' ' && ch != '\t' {
			return line[start:]
		}
		start++
	}
	return ""
}
