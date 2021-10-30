package trace

/**
 *
 * @author hushengdong
 */
func StartTrace() {

	//broker发送给nameserver的route info
	captureBroker2NameServer()
	for {
	}
}

func captureBroker2NameServer() {

	r := NewBroker2NameServerRouteInfo("", uint16(9876))
	go readMQProtocol(broker2nameserverChannel, handleBroker2NameServer)
	captureByRule(r)
}
