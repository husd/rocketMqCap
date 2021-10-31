package trace

import "github.com/google/gopacket/layers"

/**
 * RocketMQ 的基本通信协议
 * @author hushengdong
 */
//固定是8个字节
const mqprotocol_fix_len = 8
const max int = 4 // 4

type tcp_direction_type int

const (
	req  tcp_direction_type = 1 //请求包
	resp tcp_direction_type = 2 //响应包
)

type rocketMQProtocol struct {
	length            int // 消息长度：总长度，四个字节存储，占用一个int类型
	serializationType int // 序列化类型 1个字节
	headerLength      int // 三个字节 表示消息头长度

	header      []byte // 消息头数据：经过序列化后的消息头数据
	messageBody []byte // 消息主体数据：消息主体的二进制字节数据内容

	direction tcp_direction_type
}

type rocketMQHeader struct {
	Code      int
	Language  string
	Version   int
	Opaque    int
	Flag      int
	Remark    string
	ExtFields map[string]string
}

func newRocketMQHeader() *rocketMQHeader {

	res := &rocketMQHeader{}

	return res
}

func readMQProtocol(ch chan *layers.TCP, f func(*rocketMQProtocol, string, tcp_direction_type), name string, capPort int) {

outer:
	mq := &rocketMQProtocol{}

	fixOk := false
	headerOk := false
	bodyOk := false

	fixArray := [8]byte{}
	fixIndex := 0
	headerIndex := 0
	bodyIndex := 0
	bodyLength := 0

	for {
		if tcp, ok := <-ch; ok {
			data := tcp.LayerPayload()
			pos := 0
			dataLen := len(data)
			for pos < dataLen {
				//这里要处理，一次性把固定长度的数据，都读取出来
				for fixIndex < mqprotocol_fix_len && pos < dataLen {
					fixArray[fixIndex] = data[pos]
					fixIndex++
					pos++
				}
				if !fixOk && fixIndex == mqprotocol_fix_len {
					//读到了固定的8个长度
					mq.length = bytesToInt(data[0:max])
					mq.serializationType = int(data[max])
					mq.headerLength = bytes3ToInt(data[max+1 : mqprotocol_fix_len])
					bodyLength = mq.length - mqprotocol_fix_len - mq.headerLength + 4 // 找找原因为什么要+4
					mq.header = make([]byte, mq.headerLength)
					mq.messageBody = make([]byte, bodyLength)
					fixOk = true
				}
				//处理header数据
				for fixOk && !headerOk && headerIndex < mq.headerLength && pos < dataLen {
					mq.header[headerIndex] = data[pos]
					headerIndex++
					pos++
				}
				if fixOk && !headerOk && headerIndex == mq.headerLength {
					headerOk = true
				}
				//处理body数据
				for headerOk && bodyIndex < bodyLength && pos < dataLen {
					mq.messageBody[bodyIndex] = data[pos]
					bodyIndex++
					pos++
				}
				if bodyIndex == bodyLength {
					bodyOk = true
				}
				if fixOk && headerOk && bodyOk {

					dstPort := int(tcp.DstPort)
					dire := resp
					if dstPort == capPort {
						dire = req
					}

					f(mq, name, dire)
					//这里已经要注意，数组里可能还有数据呢，所以要继续处理
					goto outer
				}
			}
		}
	}
	close(ch)
}
