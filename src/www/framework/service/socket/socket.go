package SocketService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/websocket"
	"log"
	"net/http"
	"time"
	"www/framework/robot"
)

type WebSocketConn struct {
	User map[*websocket.Conn]bool
	Status bool
}

var GetWebSocket = &WebSocketConn{}

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {

	Conn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	if GetWebSocket.Status == false {
		go WebSocketSend()
		GetWebSocket.Status = true
	}

	GetWebSocket.User[Conn] = true

	for {
		_, _, err := Conn.ReadMessage()
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

			for user := range GetWebSocket.User {
				err := user.WriteMessage(1, sendJson)
				if err != nil {
					user.Close()
					delete(GetWebSocket.User, user)
				}
			}
		}

		time.Sleep(60 * time.Millisecond)
	}
}

var RobotSocket = &websocket.Conn{}

func RobotSocketClient() {

	client, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8888/message", nil)
	if err != nil {
		log.Println("\033[31m[Robot]\033[0m", "Socket Error")
	}

	RobotSocket = client

	go func() {
		for {
			_, message, err := client.ReadMessage()
			if err != nil {
				continue
			}

			Robot.Init.OnMessages(message, string(message))
		}
	}()

}

func RobotSocketClientSend(Content string) error {

	message := SocketMessage{}

	message.MessageType = "camera_message"
	message.CameraMessage.Content = Content

	sendJson, _ :=json.Marshal(Channel.Channel)

	err := RobotSocket.WriteMessage(1, sendJson)

	return err
}