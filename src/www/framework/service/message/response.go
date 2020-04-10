package MessageService

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"www/framework/robot"
)

type ResponseMessage struct {
	messageData
}

type messageData struct {
	MessageType string	`json:"message_type"`
	MessageData string 	`json:"message_data"`
}

var responseMessage mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {

	messageData := messageData{}

	err := json.Unmarshal(message.Payload(), &messageData)
	if err != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK ResponseMessage Error")
		return
	}

	Robot.Init.OnMessages(message.Payload(), string(message.Payload()))
}