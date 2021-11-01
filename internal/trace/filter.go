package trace

import "fmt"

/**
 * 过滤器
 * 过滤IP： 10.1.1.3
 * 过滤CIDR： 128.3/16
 * 过滤端口： port 53
 * 过滤主机和端口： host 8.8.8.8 and udp port 53
 * 过滤网段和端口： net 199.16.156.0/22 and port
 * 过滤非本机 Web 流量： (port 80 and port 443) and not host 192.168.0.1
 * @author hushengdong
 */

//定义过滤器 仅仅是一个描述就可以了
func filterOfTcpAndDstPort(port uint16) string {

	return fmt.Sprintf("tcp and dst port %v", port)
}

func filterOfTcpAndPort(port uint16) string {

	return fmt.Sprintf("tcp and port %v", port)
}
