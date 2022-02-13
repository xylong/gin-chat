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

// Store 保存客户端连接对象
// key 客户端地址 ip:port
// conn 连接对象
func (c *ClientMap) Store(conn *websocket.Conn) {
	// 保存客户端
	client := NewClient(conn)
	c.data.Store(conn.RemoteAddr().String(), client)

	go client.heartbeat(time.Minute * 1)
	go client.handler() // 总控制循环
	go client.read()    // 读消息循环
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
	conn      *websocket.Conn
	readChan  chan *Message // 读队列
	closeChan chan struct{}
	last      int64 // 最后发送消息的时间
}

// NewClient 创建客户端
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:      conn,
		readChan:  make(chan *Message, 10),
		closeChan: make(chan struct{}),
		last:      time.Now().Unix(),
	}
}

// heartbeat 心跳检测
func (c *Client) heartbeat(duration time.Duration) {
	defer destroyClient(c)

	// 定时检测客户端连接状态
	for {
		time.Sleep(duration)
		// 100s没有提交信息就销毁客户端
		if time.Now().Unix()-c.last > 100 {
			break
		}
	}
}

// 循环读消息
func (c *Client) read() {
	defer destroyClient(c)

loop:
	for {
		t, p, err := c.conn.ReadMessage()
		if err != nil {
			break loop
		}

		c.readChan <- NewMessage(t, p)
	}
}

// handler 消息处理
func (c *Client) handler() {
loop:
	for {
		select {
		case msg := <-c.readChan:
			logrus.Info(string(msg.Data))
			if err := msg.parseToCommand(); err != nil {
				logrus.Error(err)
			}
		case <-c.closeChan:
			break loop
		}
	}
}

// 销毁客户端
func destroyClient(client *Client) {
	if err := client.conn.Close(); err == nil {
		Clients.Destroy(client.conn)
	} else {
		logrus.Info(err)
	}
}
