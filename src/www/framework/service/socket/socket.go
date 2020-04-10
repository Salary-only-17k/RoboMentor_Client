package SocketService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/websocket"
	"net/http"
	"time"
)

var WebSocketUser = make(map[*websocket.Conn]bool)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {

	GetWebSocket, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	go WebSocketSend()

	WebSocketUser[GetWebSocket] = true

	for {
		_, message, err := GetWebSocket.ReadMessage()
		if err != nil {
			break
		}

		err = GetWebSocket.WriteMessage(1, message)
		if err != nil {
			break
		}
	}
}

func WebSocketSend() {

	for {

		if Channel.Channel.MessageType != "" {

			sendJson, _ :=json.Marshal(Channel.Channel)

			Channel.Channel = SocketMessage{}

			for user := range WebSocketUser {
				err := user.WriteMessage(1, sendJson)
				if err != nil {
					user.Close()
					delete(WebSocketUser, user)
				}
			}
		}

		time.Sleep(80 * time.Millisecond)
	}
}