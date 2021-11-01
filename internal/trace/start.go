package trace

import (
	"github.com/google/gopacket/layers"
	"sync"
)

/**
 *
 * @author hushengdong
 */
func StartTrace(port int) {

	//broker发送给nameserver的route info
	captureByPort(port)
	//captureSend2Broker()
	for {
		channelMap.Range(func(key, value interface{}) bool {
			if _, ok := threadMap.Load(key); !ok {
				go readMQProtocol(value.(chan *layers.TCP), handleMqMessage, port, key.(string))
				threadMap.Store(key, true)
			}
			return true
		})
	}
}

var channelMap = sync.Map{}
var threadMap = sync.Map{}

func captureByPort(port int) {

	r := NewSend2Broker("", uint16(port))
	captureByRule(r)
}
