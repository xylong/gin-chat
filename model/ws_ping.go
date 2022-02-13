package model

// Ping 客户端ping消息
// 心跳连接
type Ping struct {
}

func (p *Ping) ParseAction(action string) (*Response, error) {
	return NewResponse("ping", "pong"), nil
}
