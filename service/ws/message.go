package ws

import (
	"encoding/json"
)

// Message 消息
type Message struct {
	Type int    // 消息类型
	Data []byte // 消息内容
}

func NewMessage(msgType int, data []byte) *Message {
	return &Message{Type: msgType, Data: data}
}

// parseToCommand 将json消息解析未Command
func (message *Message) parseToCommand() error {
	cmd := &Command{}
	if err := json.Unmarshal(message.Data, cmd); err != nil {
		return err
	}

	return cmd.Parse()
}
