package service

import (
	"encoding/json"
	"fmt"
	"gin-chat/cache"
	"gin-chat/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	// Manager 用户管理
	Manager = ClientManager{
		// 参与连接的用户，出于性能的考虑，需要设置最大连接数
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan *Broadcast),
		Register:   make(chan *Client),
		Reply:      make(chan *Client),
		UnRegister: make(chan *Client),
	}
)

// SendMsg 发送消息
type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReplyMsg 回复消息
type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// Client 客户端
type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

// Read 读取消息
func (c *Client) Read() {
	// 下线
	defer func() {
		Manager.UnRegister <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(sendMsg)
		if err != nil {
			Manager.UnRegister <- c
			_ = c.Socket.Close()
			break
		}

		if sendMsg.Type == 1 {
			r1, _ := cache.RedisClient.Get(c.ID).Result()     // 1->2
			r2, _ := cache.RedisClient.Get(c.SendID).Result() //2->1

			// 1给2发消息，发了3条，但2没有看到或者回复，1就停止发送
			if r1 > "3" && r2 == "" {
				replyMsg := ReplyMsg{
					Code:    int(e.WsMsgLimit),
					Content: e.WsMsgLimit.String(),
				}

				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				cache.RedisClient.Incr(c.ID)
				cache.RedisClient.Expire(c.ID, time.Hour)
			}

			// 广播
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			replyMsg := &ReplyMsg{
				Code:    int(e.WsParseMsg),
				Content: string(message),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// Broadcast 广播
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// ClientManager 用户管理
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	UnRegister chan *Client
}

// Start 开始监听连接
func (cm *ClientManager) Start() {
	for {
		select {
		case conn := <-Manager.Register:
			fmt.Printf("新连接：%v", conn.ID)
			Manager.Clients[conn.ID] = conn // 新连接加入用户管理

			replyMsg := &ReplyMsg{
				Code:    int(e.WsRecord),
				Content: "已连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// Message 消息
type Message struct {
	// 发送者
	Sender string `json:"sender,omitempty"`
	// 接受者
	Recipient string `json:"recipient,omitempty"`
	// 消息内容
	Content string `json:"content,omitempty"`
}

// CreateID 创建客户端🆔
// 1->2 1发给2
func CreateID(from, to string) string {
	return from + "->" + to
}

func Handler(ctx *gin.Context) {
	from, to := ctx.Query("from"), ctx.Query("to")

	// 升级为ws协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	// 创建客户端实例
	client := &Client{
		ID:     CreateID(from, to),
		SendID: CreateID(to, from),
		Socket: conn,
		Send:   make(chan []byte),
	}

	// 注册到用户管理
	Manager.Register <- client

	go client.Read()
	go client.Write()
}
