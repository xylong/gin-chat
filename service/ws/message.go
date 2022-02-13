package ws

import (
	"encoding/json"
	"gin-chat/model"
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
func (message *Message) parseToCommand() (*model.Response, error) {
	cmd := &Command{}
	if err := json.Unmarshal(message.Data, cmd); err != nil {
		return nil, err
	}

	return cmd.Parse()
}
