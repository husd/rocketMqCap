package trace

import (
	"fmt"
)

/**
 * 解析tcp报文
 * @author hushengdong
 */
var broker2nameserverChannel chan []byte

func init() {

	broker2nameserverChannel = make(chan []byte, 1024)
}

func broker2nameserverRouteInfo(msg *[]byte) {

	broker2nameserverChannel <- *msg
}

func handleBroker2NameServer(mq *rocketMQProtocol) {

	fmt.Println("------------------------[broker发送给nameserver]--------------------")
	fmt.Println("消息长度 4字节:", mq.length)
	fmt.Println("序列化类型 1字节:", mq.serializationType)
	fmt.Println("消息头长度 3字节:", mq.headerLength)
	fmt.Println("消息头数据 :", string(mq.header))
	fmt.Println("消息主体数据 :", string(mq.messageBody))
}
