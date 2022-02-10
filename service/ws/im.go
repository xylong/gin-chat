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

	Clients.Store(conn.RemoteAddr().String(), conn)
}

func All(ctx *gin.Context) {
	content := ctx.PostForm("content")
	Clients.Send2All(content)
}
