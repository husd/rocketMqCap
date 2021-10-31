package trace

import (
	"fmt"
	"github.com/google/gopacket/layers"
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
		for k, ch := range channelMap {
			if _, ok := threadMap[k]; !ok {
				go readMQProtocol(ch, handleMqMessage, fmt.Sprintf("源IP:端口 %s 监听目的端口: %d", k, port))
				threadMap[k] = true
			}
		}
	}
}

var channelMap = make(map[string]chan *layers.TCP)
var threadMap = make(map[string]bool)

func captureByPort(port int) {

	r := NewSend2Broker("", uint16(port))
	captureByRule(r)
}
