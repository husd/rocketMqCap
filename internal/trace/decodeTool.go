package trace

import "fmt"

func DecodeTcpBytes(f string) {

	r := NewTcpDumpFileByteReader(f)

	//readMqMsg(r,handleMqMessage,"read tcpdump file:")
	for r.hasMore() {
		fmt.Println("----123:" + string(r.readBytes(512)))
	}
}

func readMqMsg(reader byteReader, f func(*rocketMQProtocol, string), desc string) {

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

	for reader.hasMore() {
		data := reader.readBytes(512)
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
				bodyLength = mq.length - 4 - mq.headerLength
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
				f(mq, desc)
				//这里已经要注意，数组里可能还有数据呢，所以要继续处理
				goto outer
			}
		}
	}
}
