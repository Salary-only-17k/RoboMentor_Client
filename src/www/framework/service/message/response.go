package MessageService

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
	"www/framework/config"
	"www/framework/function/command"
	"www/framework/function/serial"
	"www/framework/robot"
	"www/framework/service/common"
)

type ResponseMessage struct {
	messageData
}

type messageData struct {
	MessageType 	string					`json:"message_type"`
	SystemMessage 	systemMessage 			`json:"system_message"`
	RobotConfig 	robotConfig 			`json:"robot_config"`
	RobotRun 		robotRun 				`json:"robot_run"`
	SerialMessage 	serialMessage 			`json:"serial_message"`
	SerialMessageRead 	serialMessageRead 	`json:"serial_message_read"`
	SerialMessageError 	serialMessageError 	`json:"serial_message_error"`
	RemoteMessage	remoteMessage 			`json:"remote_message"`
	ServiceMessage	serviceMessage 			`json:"service_message"`
	TcpMessage		tcpMessage 				`json:"tcp_message"`
	TcpMessageRead	tcpMessageRead 			`json:"tcp_message_read"`
	TcpMessageError	tcpMessageError 		`json:"tcp_message_error"`
	CameraMessage	cameraMessage 			`json:"camera_message"`
	DetectMessage	detectMessage 			`json:"detect_message"`
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

type serialMessageRead struct {
	Content string 		`json:"content"`
}

type serialMessageError struct {
	Content string 		`json:"content"`
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

type tcpMessageRead struct {
	Content string 		`json:"content"`
}

type tcpMessageError struct {
	Content string 		`json:"content"`
}

type cameraMessage struct {
	Content string 		`json:"content"`
}

type detectMessage struct {
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

	if messageData.MessageType == "remote_message" {
		Robot.Init.OnMessages([]byte(messageData.RemoteMessage.Content), messageData.RemoteMessage.Content)
	}

	if messageData.MessageType == "serial_message" {

		sendMessage := ResponseMessage{}

		isExists := CommonService.Exists(messageData.SerialMessage.Port)
		if isExists == false {
			sendMessage.MessageType = "serial_message_error"
			sendMessage.SerialMessageError.Content = "没有找到相关串口，请重新尝试"
		}

		if isExists {
			RateInt, _ := strconv.Atoi(messageData.SerialMessage.Rate)

			sendStatus := serialFunction.SerialWrite(messageData.SerialMessage.Port, RateInt, messageData.SerialMessage.Content)
			if sendStatus == false {
				sendMessage.MessageType = "serial_message_error"
				sendMessage.SerialMessageError.Content = "串口数据发送失败，请重新尝试"
			}

			if messageData.SerialMessage.Switch {

				bits, _ := strconv.Atoi(messageData.SerialMessage.Bits)

				readContent := serialFunction.SerialRead(messageData.SerialMessage.Port, RateInt, bits)

				sendMessage.MessageType = "serial_message_read"
				sendMessage.SerialMessageRead.Content = readContent
			}
		}

		sendString, _ := json.Marshal(sendMessage)
		Send("", string(sendString))
	}

	if messageData.MessageType == "tcp_message" {

		sendMessage := ResponseMessage{}

		conn, err := net.DialTimeout("tcp", messageData.TcpMessage.Ip + ":" + messageData.TcpMessage.Port, 2*time.Second)
		if err != nil {
			sendMessage.MessageType = "tcp_message_error"
			sendMessage.TcpMessageError.Content = "没有找到相关通讯端口，请重新尝试"
		}

		if err == nil {
			writeInt, err := conn.Write([]byte(messageData.TcpMessage.Content))
			if err != nil {
				sendMessage.MessageType = "tcp_message_error"
				sendMessage.TcpMessageError.Content = "数据发送失败，" + strconv.Itoa(writeInt) + "，请求重新尝试"
			}

			if messageData.TcpMessage.Switch {
				sendMessage.MessageType = "tcp_message_read"
				sendMessage.TcpMessageError.Content = "读取数据还在优化中，暂时不可用"
			}
		}

		sendString, _ := json.Marshal(sendMessage)
		Send("", string(sendString))
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
}