package trace

import "github.com/google/gopacket/layers"

/**
 * RocketMQ 的基本通信协议
 * @author hushengdong
 */
//固定是8个字节
const mqprotocol_fix_len = 8
const max int = 4 // 4

type rocketMQProtocol struct {
	length            int // 消息长度：总长度，四个字节存储，占用一个int类型
	serializationType int // 序列化类型 1个字节
	headerLength      int // 三个字节 表示消息头长度

	header      []byte // 消息头数据：经过序列化后的消息头数据
	messageBody []byte // 消息主体数据：消息主体的二进制字节数据内容
}

func readMQProtocol(ch chan *layers.TCP, f func(*rocketMQProtocol, string), name string) {

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
					f(mq, name)
					//这里已经要注意，数组里可能还有数据呢，所以要继续处理
					goto outer
				}
			}
		}
	}
	close(ch)
}
