package trace

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket/layers"
)

/**
 * 定制的过滤的规则和handle
 * @author hushengdong
 */
type rule struct {
	name   string // 具体是那个监听
	filter string // 这个监听信息的过滤规则
	port   int    // 监听的端口号

	ruleHandle func(srcIp string, tcp *layers.TCP)
}

/**
 * broker发送给nameserver的route info
 */
func NewSend2Broker(ip string, port uint16) *rule {

	r := &rule{}

	r.name = "发送给broker"
	r.filter = filterOfTcpAndDstPort(port)
	r.ruleHandle = doRuleHandle
	r.port = int(port)

	return r
}

func doRuleHandle(srcIp string, tcp *layers.TCP) {

	key := srcIp + ":" + tcp.SrcPort.String()
	var channel chan *layers.TCP
	if value, ok := channelMap.Load(key); !ok {
		channel = make(chan *layers.TCP, 1024)
		channelMap.Store(key, channel)
	} else {
		channel = value.(chan *layers.TCP)
	}
	ch := channel
	ch <- tcp
}

func handleMqMessage(mq *rocketMQProtocol, name string, dire tcp_direction_type) {

	now := now()
	fmt.Println(now, "------------------------[", name, "]--------------------")
	//fmt.Println(now, " 消息长度 4字节:", mq.length)
	//fmt.Println(now, " 序列化类型 1字节:", mq.serializationType)
	//fmt.Println(now, " 消息头长度 3字节:", mq.headerLength)
	header := string(mq.header)
	mqHeader := newRocketMQHeader()
	err := json.Unmarshal(mq.header, mqHeader)
	if err != nil {
		fmt.Println(now, "消息头数据 :", header)
	} else {
		msg := ""
		if dire == req {
			msg = getReqCodeMsg(mqHeader.Code)
			fmt.Println(now, "[", msg, "][请求](", mqHeader.Code, ") 消息头数据 :", header)
		} else if dire == resp {
			msg = getRespCodeMsg(mqHeader.Code)
			fmt.Println(now, "[", msg, "][响应](", mqHeader.Code, ") 消息头数据 :", header)
		}
	}
	fmt.Println(now, " 消息主体数据 :", string(mq.messageBody))
}
