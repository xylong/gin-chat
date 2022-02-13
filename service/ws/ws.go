package ws

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ws http->websocket
func Ws(ctx *gin.Context) {
	// 升级http协议
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	Clients.Store(conn)
}

// All 向所有客户端发消息
func All(ctx *gin.Context) {
	content := ctx.PostForm("content")
	Clients.Send2All(content)
}

func Json(ctx *gin.Context) {
	Clients.data.Range(func(key, value interface{}) bool {
		_ = value.(*Client).conn.WriteJSON([]struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			{1, "apple"},
			{2, "banana"},
		})
		return true
	})
}
