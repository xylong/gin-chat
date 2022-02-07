package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	upgrader websocket.Upgrader
)

func init() {
	upgrader = websocket.Upgrader{
		// 跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

// Http2Ws http->websocket
func Http2Ws(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	for {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			logrus.Error("", err)
		}

		time.Sleep(time.Second * 5)
	}
}
