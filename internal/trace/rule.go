package trace

/**
 * 定制的过滤的规则和handle
 * @author hushengdong
 */
type rule struct {
	name   string // 具体是那个监听
	filter string // 这个监听信息的过滤规则

	ruleHandle func(msg *[]byte)
}

/**
 * broker发送给nameserver的route info
 */
func NewBroker2NameServerRouteInfo(ip string, port uint16) *rule {

	r := &rule{}

	r.name = "broker发送给nameserver"
	r.filter = filterOfTcpAndDstPort(port)
	r.ruleHandle = broker2nameserverRouteInfo

	return r
}
