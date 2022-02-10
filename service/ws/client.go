package ws

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	// Clients 所有客户端
	Clients *ClientMap
)

func init() {
	Clients = &ClientMap{}
}

// ClientMap 客户端对象集合
type ClientMap struct {
	data sync.Map
}

// Store 保存客户端连接对象
// key 客户端地址 ip:port
// conn 连接对象
func (c *ClientMap) Store(key string, conn *websocket.Conn) {
	c.data.Store(key, conn)
}

// Send2All 向所有客户端发消息
func (c *ClientMap) Send2All(content string) {
	c.data.Range(func(key, value interface{}) bool {
		err := value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(content))
		if err != nil {
			logrus.Info(err)
		}

		return err == nil
	})
}
