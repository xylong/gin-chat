package ws

// Message 消息
type Message struct {
	Type int    // 消息类型
	Data []byte // 消息内容
}

func NewMessage(msgType int, data []byte) *Message {
	return &Message{Type: msgType, Data: data}
}
