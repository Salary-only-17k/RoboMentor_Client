package SocketService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/websocket"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Socket(c *gin.Context) {

	GetSocket, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	go SocketSend(GetSocket)

	for {
		_, message, err := GetSocket.ReadMessage()
		if err != nil {
			break
		}

		err = GetSocket.WriteMessage(1, message)
		if err != nil {
			break
		}
	}
}

func SocketSend(GetSocket *websocket.Conn) {

	for {

		if Channel.Channel.MessageType != "" {

			sendJson, _ :=json.Marshal(Channel.Channel)

			err := GetSocket.WriteMessage(1, sendJson)
			if err != nil {
				continue
			}

			Channel.Channel = SocketMessage{}
		}

		time.Sleep(80 * time.Millisecond)
	}
}