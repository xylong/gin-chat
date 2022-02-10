package ws

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	// Clients 所有客户端
	Clients *ClientMap
)

func init() {
	Clients = &ClientMap{}
}

type
// ClientMap 客户端对象集合
ClientMap struct {
	data sync.Map
}

// NewClient 创建客户端
func NewClient(conn *websocket.Conn) *Client {
	return &Client{conn: conn}
}

// Store 保存客户端连接对象
// key 客户端地址 ip:port
// conn 连接对象
func (c *ClientMap) Store(conn *websocket.Conn) {
	client := NewClient(conn)
	client.last = time.Now().Unix()

	c.data.Store(conn.RemoteAddr().String(), client)
	go client.heartbeat(time.Minute * 1)
}

// Send2All 向所有客户端发消息
func (c *ClientMap) Send2All(content string) {
	c.data.Range(func(key, value interface{}) bool {
		err := value.(*Client).conn.WriteMessage(websocket.TextMessage, []byte(content))
		if err != nil {
			logrus.Info(err)
		}

		return err == nil
	})
}

// Destroy 销毁客户端
func (c *ClientMap) Destroy(conn *websocket.Conn) {
	c.data.Delete(conn.RemoteAddr().String())
}

// Client 客户端
type Client struct {
	conn *websocket.Conn
	last int64 // 最后发送消息的时间
}

// heartbeat 心跳检测
func (c *Client) heartbeat(duration time.Duration) {
	defer c.conn.Close()

	// 定时检测客户端连接状态
	for {
		time.Sleep(duration)
		// 100s没有提交信息就销毁客户端
		if time.Now().Unix()-c.last > 100 {
			Clients.Destroy(c.conn)
			break
		}
	}
}
