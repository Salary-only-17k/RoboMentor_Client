package MessageService

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"www/framework/config"
)

type Message struct {
	Client mqtt.Client
}

var MessageClient = &Message{}

func Messages(AppID string, AppSecret string) {

	clientOptions := mqtt.NewClientOptions().AddBroker(Config.MentorConfig.RobotMessage).SetClientID("robot-" + Config.MentorConfig.RobotMac)

	clientOptions.SetUsername(AppID)

	clientOptions.SetPassword(AppSecret)

	clientOptions.SetAutoReconnect(true)

	clientOptions.SetCleanSession(true)

	client := MessageClient

	client.Client = mqtt.NewClient(clientOptions)
	if token := client.Client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient MessageClient Connect Error")
	}

	topic := "robot/" + Config.MentorConfig.RobotMac

	if token := client.Client.Subscribe(topic, 0, responseMessage); token.Wait() && token.Error() != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient MessageClient Subscribe Error")
	}

	sendMessage := ResponseMessage{}
	sendMessage.MessageType = "robot_run"
	sendMessage.RobotRun.Type = "start"
	sendMessage.RobotRun.Content = ""
	sendString, _ := json.Marshal(sendMessage)
	Send("", string(sendString))
}

func Send(topicData string, messageContent string) error {

	if topicData == "" {
		topicData = "robot/" + Config.MentorConfig.RobotMac
	}

	messageClient := MessageClient

	sendToken := messageClient.Client.Publish(topicData, 0, false, messageContent)

	if sendToken.Error() != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient Send Message Error")
	}

	return sendToken.Error()
}