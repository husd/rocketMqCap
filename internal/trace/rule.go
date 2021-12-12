package trace

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket/layers"
	"sync"
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
	r.filter = filterOfTcpAndPort(port)
	r.ruleHandle = doRuleHandle
	r.port = int(port)

	return r
}

func doRuleHandle(key string, tcp *layers.TCP) {

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

var lock sync.Mutex

func handleMqMessage(mq *rocketMQProtocol, desc string) {

	lock.Lock()
	defer lock.Unlock()
	now := now()
	fmt.Println(now, "------------------------[", desc, "]--------------------")
	//fmt.Println(now, " 消息长度 4字节:", mq.length)
	//fmt.Println(now, " 序列化类型 1字节:", mq.serializationType)
	//fmt.Println(now, " 消息头长度 3字节:", mq.headerLength)
	header := string(mq.header)
	mqHeader := newRocketMQHeader()
	err := json.Unmarshal(mq.header, mqHeader)
	if err != nil {
		fmt.Println(now, "消息头数据失败 :", header)
	} else {
		msg, dire := getCodeMsg(mqHeader)
		if dire == req && mqHeader.Code == 34 {
			fmt.Println(now, "[", mqHeader.Code, "][请求][", msg, "] 消息头数据 :", header, " 消息体: ", string(mq.messageBody))
		} else if dire == resp {
			//fmt.Println(now, "[", mqHeader.Code, "][响应][", msg, "] 消息头数据 :", header)
		}
	}
	//fmt.Println(now, "[消息主体数据] :", string(mq.messageBody))
}

func getCodeMsg(header *rocketMQHeader) (string, tcp_direction_type) {

	code := header.Code
	msgReq, okReq := getReqCodeMsg(code)
	msgResp, okResp := getRespCodeMsg(code)
	//都有，看备注信息是否为空了
	if okReq && okResp {
		//有备注的，一般是响应报文
		if len(header.Remark) > 0 {
			return msgResp, resp
		} else {
			return msgReq, req
		}
	} else if okResp {
		return msgResp, resp
	} else if okReq {
		return msgReq, req
	}
	return msgReq, req
}
