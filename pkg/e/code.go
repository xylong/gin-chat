package e

const (
	WsParseMsg     = 10001 // 解析内容信息
	WsRecord       = 10002 // 发送信息，请求历史记录
	WsRecordEnd    = 10003 // 没有更多记录了
	WsOnlineReply  = 10004 // 在线应答成功
	WsOfflineReply = 10004 // 离线应答成功
	WsMsgLimit     = 10005 // 请求受到限制
)
