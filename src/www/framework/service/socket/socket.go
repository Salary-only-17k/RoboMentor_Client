package SocketService

import (
	"github.com/gin-gonic/gin"
	"github.com/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Socket(c *gin.Context) {

	GetSocket, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

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