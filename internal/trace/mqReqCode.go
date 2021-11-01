package trace

/**
 *
 * @author hushengdong
 */
var reqCodeMsgMap [317]string = [317]string{}

func init() {

	reqCodeMsgMap[10] = "Broker 发送消息"
	reqCodeMsgMap[11] = "Broker 订阅消息"
	reqCodeMsgMap[12] = "Broker 查询消息"
	reqCodeMsgMap[13] = "Broker 查询Broker Offset"
	reqCodeMsgMap[14] = "Broker 查询Consumer Offset"
	reqCodeMsgMap[15] = "Broker 更新Consumer Offset"
	reqCodeMsgMap[17] = "Broker 更新或者增加一个Topic"
	reqCodeMsgMap[21] = "Broker 获取所有Topic的配置（Slave和Namesrv都会向Master请求此配置）"
	reqCodeMsgMap[22] = "Broker 获取所有Topic配置（Slave和Namesrv都会向Master请求此配置）"
	reqCodeMsgMap[23] = "Broker 获取所有Topic名称列表"
	reqCodeMsgMap[25] = "Broker 更新Broker上的配置"
	reqCodeMsgMap[26] = "Broker 获取Broker上的配置"
	reqCodeMsgMap[27] = "Broker 触发Broker删除文件"
	reqCodeMsgMap[28] = "Broker 获取Broker运行时信息"
	reqCodeMsgMap[29] = "Broker 根据时间查询队列的Offset"
	reqCodeMsgMap[30] = "Broker 查询队列最大Offset"
	reqCodeMsgMap[31] = "Broker 查询队列最小Offset"
	reqCodeMsgMap[32] = "Broker 查询队列最早消息对应时间"
	reqCodeMsgMap[33] = "Broker 根据消息ID来查询消息"
	reqCodeMsgMap[34] = "Broker Client向Client发送心跳，并注册自身"
	reqCodeMsgMap[35] = "Broker Client注销"
	reqCodeMsgMap[36] = "Broker Consumer将处理不了的消息发回服务器"
	reqCodeMsgMap[37] = "Broker Commit或者Rollback事务"
	reqCodeMsgMap[38] = "Broker 获取ConsumerId列表通过GroupName"
	reqCodeMsgMap[39] = "Broker 主动向Producer回查事务状态"
	reqCodeMsgMap[40] = "Broker Broker通知Consumer列表变化"
	reqCodeMsgMap[41] = "Broker Consumer向Master锁定队列"
	reqCodeMsgMap[42] = "Broker Consumer向Master解锁队列"
	reqCodeMsgMap[43] = "Broker 获取所有Consumer Offset"
	reqCodeMsgMap[45] = "Broker 获取所有定时进度"
	reqCodeMsgMap[100] = "Namesrv 向Namesrv追加KV配置"
	reqCodeMsgMap[101] = "Namesrv 从Namesrv获取KV配置"
	reqCodeMsgMap[102] = "Namesrv 从Namesrv获取KV配置"
	reqCodeMsgMap[103] = "Namesrv 注册一个Broker，数据都是持久化的，如果存在则覆盖配置"
	reqCodeMsgMap[104] = "Namesrv 卸载一个Broker，数据都是持久化的"
	reqCodeMsgMap[105] = "Namesrv 根据Topic获取Broker Name、队列数(包含读队列与写队列)"
	reqCodeMsgMap[106] = "Namesrv 获取注册到Name Server的所有Broker集群信息"
	reqCodeMsgMap[200] = "更新并创建消费组"
	reqCodeMsgMap[201] = "获取所有的消费组配置"
	reqCodeMsgMap[202] = "获取TOPIC信息"
	reqCodeMsgMap[203] = "获取消费者的连接链表"
	reqCodeMsgMap[204] = "获取生产者的连接链表"
	reqCodeMsgMap[205] = "禁用broker的读写，相当于摘掉这个broker"
	reqCodeMsgMap[206] = "从Name Server获取完整Topic列表"
	reqCodeMsgMap[207] = "从Broker删除订阅组"
	reqCodeMsgMap[208] = "从Broker获取消费状态（进度）"
	reqCodeMsgMap[209] = "Suspend Consumer消费过程"
	reqCodeMsgMap[210] = "Resume Consumer消费过程"
	reqCodeMsgMap[211] = "consumer重置Consumer Offset"
	reqCodeMsgMap[212] = "broker重置Consumer Offset"
	reqCodeMsgMap[213] = "调整Consumer线程池数量"
	reqCodeMsgMap[214] = "查询消息被哪些消费组消费"
	reqCodeMsgMap[215] = "从Broker删除Topic配置"
	reqCodeMsgMap[216] = "从Namesrv删除Topic配置"
	reqCodeMsgMap[217] = "Namesrv 通过 project 获取所有的 server ip 信息"
	reqCodeMsgMap[218] = "Namesrv 删除指定 project group 下的所有 server ip 信息"
	reqCodeMsgMap[219] = "通过NameSpace获取所有的KV List"
	reqCodeMsgMap[220] = "重置consumer客户端的offset"
	reqCodeMsgMap[221] = "客户端订阅消息"
	reqCodeMsgMap[222] = "通知 broker 调用 offset 重置处理"
	reqCodeMsgMap[223] = "通知 broker 调用客户端订阅消息处理"
	reqCodeMsgMap[300] = "Broker 查询topic被谁消费"
	reqCodeMsgMap[224] = "获取指定集群下的所有 topic"
	reqCodeMsgMap[301] = "向Broker注册Filter Server"
	reqCodeMsgMap[302] = "向Filter Server注册Class"
	reqCodeMsgMap[303] = "根据 topic 和 group 获取消息的时间跨度"
	reqCodeMsgMap[304] = "从nameserver获取所有系统内置 Topic 列表"
	reqCodeMsgMap[305] = "从broker获取所有系统内置 Topic 列表"
	reqCodeMsgMap[306] = "清理失效队列"
	reqCodeMsgMap[307] = "通过Broker查询Consumer内存数据"
	reqCodeMsgMap[308] = "查找被修正 offset (转发组件）"
	reqCodeMsgMap[309] = "通过Broker直接向某个Consumer发送一条消息，并立刻消费，返回结果给broker，再返回给调用方"
	reqCodeMsgMap[310] = "Broker 发送消息，优化网络数据包"
	reqCodeMsgMap[311] = "单元化相关 topic"
	reqCodeMsgMap[312] = "获取含有单元化订阅组的 Topic 列表"
	reqCodeMsgMap[313] = "获取含有单元化订阅组的非单元化 Topic 列表"
	reqCodeMsgMap[314] = "克隆某一个组的消费进度到新的组"
	reqCodeMsgMap[315] = "查看Broker上的各种统计信息"
	reqCodeMsgMap[316] = "* 发送死信队列 */"
}

func getReqCodeMsg(code int) (string, bool) {

	if code >= 0 && code < 316 {
		msg := reqCodeMsgMap[code]
		if len(msg) > 0 {
			return msg, true
		} else {
			return "", false
		}
	}
	return "", false
}
