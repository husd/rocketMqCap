# view-package
为了研究rocketMQ原理，所以写的一个抓包分析工具，可以解析nameserver和broker等之间的数据交互。

# 使用办法

- 编译 go build husd.com/trace/main
- 运行 main.exe --port=57916

```shell
2021-11-01 11:27:52  ------------------------[ 源ip:port 127.0.0.1:9876 ~ 目的ip:port 127.0.0.1:59300 监听端口: 59300 ]--------------------
2021-11-01 11:27:52  [ 请求成功 ][响应][ 0 ] 消息头数据 : {"code":0,"extFields":{},"flag":1,"language":"JAVA","opaque":292,"version":79}
2021-11-01 11:27:52  [消息主体数据] : 
2021-11-01 11:27:52  ------------------------[ 源ip:port 127.0.0.1:59300 ~ 目的ip:port 127.0.0.1:9876 监听端口: 59300 ]--------------------
2021-11-01 11:27:52  [ Namesrv 注册一个Broker，数据都是持久化的，如果存在则覆盖配置 ][请求][ 103 ] 消息头数据 : {"code":103,"extFields":{"brokerId":"0","clusterName":"DefaultCluster","","haServerAddr":"10.2.144.15:10912","brokerName":"hushengdong"},"flag":0,"language":"JAVA","opaque":292,"version":79}
2021-11-01 11:27:52  [消息主体数据] : {"filterServerList":[],"topicConfigSerializeWrapper":{"dataVersion":{"counter":3,"timestatmp":1635732910223},"topicConfigTable":{"SCHEDULE_TOPIC_X"order":false,"perm":6,"readQueueNums":18,"topicFilterType":"SINGLE_TAG","topicName":"SCHEDULE_TOPIC_XXXX","topicSysFlag":0,"writeQueueNums":18},"TopicTest":{"order":false,"perm":6,"readQueueNums":4,"topicFilterType":"SINGLE_TAG","topicName":"TopicTest","topicSysFlag":0,"writeQueueNums":4},"SELF_TEST_TOPIC":{"order":false,"perm":6,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"SELF_TEST_TOPIC","topicSysFlag":0,"writeQueueNums":1},"DefaultCluster":{"order":false,"perm":7,"readQueueNums":16,"topicFilterType":"SINGLE_TAG","topicName":"DefaultCluster","topicSysFlag":0,"writeQueueNums":16},"DefaultCluster_REPLY_TOPIC":{"order":false,"perm":6,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"DefaultCluster_REPLY_TOPIC","topicSysFlag":0,"writeQueueNums":1},"RMQ_SYS_TRANS_HALF_TOPIC":{"order":false,"perm":6,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"RMQ_SYS_TRANS_HALF_TOPIC","topicSysFlag":0,"writeQueueNums":1},"hushengdong":{"order":false,"perm":7,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"hushengdong","topicSysFlag":0,"writeQueueNums":1},"TBW102":{"order":false,"perm":7,"readQueueNums":8,"topicFilterType":"SINGLE_TAG","topicName":"TBW102","topicSysFlag":0,"writeQueueNums":8},"BenchmarkTest":{"order":false,"perm":6,"readQueueNums":1024,"topicFilterType":"SINGLE_TAG","topicName":"BenchmarkTest","topicSysFlag":0,"writeQueueNums":1024},"OFFSET_MOVED_EVENT":{"order":false,"perm":6,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"OFFSET_MOVED_EVENT","topicSysFlag":0,"writeQueueNums":1},"%RETRY%please_rename_unique_group_name_4":{"order":false,"perm":6,"readQueueNums":1,"topicFilterType":"SINGLE_TAG","topicName":"%RETRY%please_rename_unique_group_name_4","topicSysFlag":0,"writeQueueNums":1}}}}
```
阅读rocketMQ的官方文档，可以知道协议是： 消息长度 + 序列化类型&消息头长度 + 消息头数据 + 消息主体数据 这5部分组成。
具体的解释，可以参考RequestCode里的值，比如返回消息的code:0 表示成功，其它表示失败。
请求消息的code表示各种业务关系，例如：Broker 发送消息 的 code = 10

由于是开启了多线程处理的报文，所以有时候会出现响应报文比请求报文先打印出来的情况，mq的协议是通过opaque这个字段来对请求报文
和响应报文进行1对1匹配的，所以监听的时候，可以通过这个opaque来确认那2个报文是匹配的。

# 其它

- 目前仅支持监听1个端口，如果需要同时监听2个端口，可以开2个终端。
- 默认监听的端口是9876，是nameServer的端口。
