package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader *websocket.Upgrader
)

func init() {
	upgrader = &websocket.Upgrader{
		// 跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
