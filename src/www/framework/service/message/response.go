package MessageService

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"strings"
	"www/framework/config"
	"www/framework/function/command"
	"www/framework/robot"
	"www/framework/service/common"
)

type ResponseMessage struct {
	messageData
}

type messageData struct {
	MessageType 	string			`json:"message_type"`
	SystemMessage 	systemMessage 	`json:"system_message"`
	RobotConfig 	robotConfig 	`json:"robot_config"`
	RobotRun 		robotRun 		`json:"robot_run"`
	SerialMessage 	serialMessage 	`json:"serial_message"`
	RemoteMessage	remoteMessage 	`json:"remote_message"`
	ServiceMessage	serviceMessage 	`json:"service_message"`
	TcpMessage		tcpMessage 		`json:"tcp_message"`
	CameraMessage	cameraMessage 	`json:"camera_message"`
}

type systemMessage struct {
	Time    string 		`json:"time"`
	Memory 	float64 	`json:"memory"`
	Cpu 	[]float64 	`json:"cpu"`
}

type serialMessage struct {
	Port    string 		`json:"port"`
	Rate 	string 		`json:"rate"`
	Bits 	string 		`json:"bits"`
	Content string 		`json:"content"`
	Switch 	bool 		`json:"Switch"`
}

type robotConfig struct {
	Content string 		`json:"content"`
}

type robotRun struct {
	Type 	string `json:"type"`
	Content string 		`json:"content"`
}

type remoteMessage struct {
	Content string 		`json:"content"`
}

type serviceMessage struct {
	Content string 		`json:"content"`
}

type tcpMessage struct {
	Ip    	string 		`json:"ip"`
	Port 	string 		`json:"port"`
	Content string 		`json:"content"`
	Switch 	bool 		`json:"Switch"`
}

type cameraMessage struct {
	Content string 		`json:"content"`
}

var responseMessage mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {

	messageData := messageData{}

	err := json.Unmarshal(message.Payload(), &messageData)
	if err != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient ResponseMessage Error")
		return
	}

	if messageData.MessageType == "robot_config" {

		configData := Config.ResponseRobot{}

		json.Unmarshal([]byte(messageData.RobotConfig.Content), &configData)

		if configData.RobotTitle != "" {
			Config.MentorConfig.RobotName = configData.RobotTitle
		}

		if configData.RobotBoard.Port != "" {
			isExists := CommonService.Exists(configData.RobotBoard.Port)
			if isExists == true {
				Config.MentorConfig.RobotBoard.Port = configData.RobotBoard.Port
				Config.MentorConfig.RobotBoard.Bits = configData.RobotBoard.Bits
				Config.MentorConfig.RobotBoard.Rate = configData.RobotBoard.Rate
			}
		}
	}

	if messageData.MessageType == "robot_run" {
		if messageData.RobotRun.Type == "code" {
			ioutil.WriteFile("application/robot.go", []byte(messageData.RobotRun.Content), 0777)
		}

		if messageData.RobotRun.Type == "build" {
			shell, err := commandFunction.Shell("sudo framework/function/command/build.sh")
			if err != nil {
				sendMessage := ResponseMessage{}
				sendMessage.MessageType = "robot_run"
				sendMessage.RobotRun.Type = "build_error"
				sendMessage.RobotRun.Content = ""
				sendString, _ := json.Marshal(sendMessage)
				Send("", string(sendString))
			}else{
				if strings.Contains(shell, "Build Success") == false {
					sendMessage := ResponseMessage{}
					sendMessage.MessageType = "robot_run"
					sendMessage.RobotRun.Type = "build_error"
					sendMessage.RobotRun.Content = ""
					sendString, _ := json.Marshal(sendMessage)
					Send("", string(sendString))
				}else{
					sendMessage := ResponseMessage{}
					sendMessage.MessageType = "robot_run"
					sendMessage.RobotRun.Type = "build_success"
					sendMessage.RobotRun.Content = ""
					sendString, _ := json.Marshal(sendMessage)
					Send("", string(sendString))
				}
			}
		}

		if messageData.RobotRun.Type == "restart" {
			commandFunction.Shell("sudo framework/function/command/restart.sh")
		}
	}

	Robot.Init.OnMessages(message.Payload(), string(message.Payload()))
}