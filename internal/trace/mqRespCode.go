package trace

/**
 *
 * @author hushengdong
 */
var respCodeMsgMap [208]string = [208]string{}

func init() {

	respCodeMsgMap[0] = "请求成功"

	respCodeMsgMap[10] = "Broker 刷盘超时"
	respCodeMsgMap[11] = "Broker 同步双写，Slave不可用"
	respCodeMsgMap[12] = "Broker 同步双写，等待Slave应答超时"
	respCodeMsgMap[13] = "Broker 消息非法"
	respCodeMsgMap[14] = "Broker, Namesrv 服务不可用，可能是正在关闭或者权限问题"
	respCodeMsgMap[15] = "Broker, Namesrv 版本号不支持"
	respCodeMsgMap[16] = "Broker, Namesrv 无权限执行此操作，可能是发、收、或者其他操作"
	respCodeMsgMap[17] = "Broker, Topic不存在"
	respCodeMsgMap[18] = "Broker, Topic已经存在，创建Topic"
	respCodeMsgMap[19] = "Broker 拉消息未找到（请求的Offset等于最大Offset，最大Offset无对应消息）"
	respCodeMsgMap[20] = "Broker 可能被过滤，或者误通知等"
	respCodeMsgMap[21] = "Broker 拉消息请求的Offset不合法，太小或太大"
	respCodeMsgMap[22] = "Broker 查询消息未找到"
	respCodeMsgMap[23] = "Broker 订阅关系解析失败"
	respCodeMsgMap[24] = "Broker 订阅关系不存在"
	respCodeMsgMap[25] = "Broker 订阅关系不是最新的"
	respCodeMsgMap[26] = "Broker 订阅组不存在"
	respCodeMsgMap[200] = "Producer 事务应该被提交"
	respCodeMsgMap[201] = "Producer 事务应该被回滚"
	respCodeMsgMap[202] = "Producer 事务状态未知"
	respCodeMsgMap[203] = "Producer ProducerGroup错误"
	respCodeMsgMap[204] = "单元化消息，需要设置 buyerId"
	respCodeMsgMap[205] = "单元化消息，非本单元消息"
	respCodeMsgMap[206] = "Consumer不在线"
	respCodeMsgMap[207] = "Consumer消费消息超时"
}

func getRespCodeMsg(code int) (string, bool) {

	if code >= 0 && code < 208 {
		msg := respCodeMsgMap[code]
		if len(msg) > 0 {
			return msg, true
		} else {
			return "", false
		}
	}
	return "未知代码", false
}
