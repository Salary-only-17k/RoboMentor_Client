package MessageService

import (
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
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK MessageClient Connect Error")
	}

	topic := "robot/" + Config.MentorConfig.RobotMac

	if token := client.Client.Subscribe(topic, 0, responseMessage); token.Wait() && token.Error() != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK MessageClient Subscribe Error")
	}
}