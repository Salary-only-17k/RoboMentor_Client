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
	"www/framework/platform/servo"
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
	ServoMessage	servoMessage 			`json:"servo_message"`
	ServoControlMessage servoControlMessage `json:"servo_control_message"`
	ServoMessageRead servoMessageRead 		`json:"servo_message_read"`
	ServoMessageError	servoMessageError 	`json:"servo_message_error"`
	CameraMessage	cameraMessage 			`json:"camera_message"`
	DetectMessage	detectMessage 			`json:"detect_message"`
}

type systemMessage struct {
	Time    string 		`json:"time"`
	Memory 	float64 	`json:"memory"`
	Cpu 	[]float64 	`json:"cpu"`
}

type serialMessage struct {
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
	Type 	string 		`json:"type"`
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

type ServoMessage struct {
	servoMessage
}

type servoMessage struct {
	Type        string `json:"type"`
	Id 			string `json:"id"`
	Channel 	string `json:"channel"`
	Mode 		string `json:"mode"`
	Speed 		string `json:"speed"`
	Value 		string `json:"value"`
	MinAngle 	string `json:"minAngle"`
	MaxAngle 	string `json:"maxAngle"`
	MinVin 		string `json:"minVin"`
	MaxVin 		string `json:"maxVin"`
	Status 		string `json:"status"`
}

type ServoControlMessage struct {
	servoControlMessage
}

type servoControlMessage struct {
	Id 			string `json:"id"`
	Channel 	string `json:"channel"`
	Angle 		string `json:"angle"`
	Time 		string `json:"time"`
}

type servoMessageRead struct {
	servoMessage
}

type servoMessageError struct {
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

		isExists := CommonService.Exists(Config.MentorConfig.RobotBoard.Port)
		if isExists == false {
			sendMessage.MessageType = "serial_message_error"
			sendMessage.SerialMessageError.Content = "没有找到相关串口，请重新尝试"
		}

		if isExists {

			sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, messageData.SerialMessage.Content)
			if sendStatus == false {
				sendMessage.MessageType = "serial_message_error"
				sendMessage.SerialMessageError.Content = "串口数据发送失败，请重新尝试"
			}

			if messageData.SerialMessage.Switch {

				readContent := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)

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

	if messageData.MessageType == "servo_message" {

		sendMessage := ResponseMessage{}

		isExists := CommonService.Exists(Config.MentorConfig.RobotBoard.Port)
		if isExists == false {
			sendMessage.MessageType = "servo_message_error"
			sendMessage.ServoMessageError.Content = "舵机通讯失败，请重新尝试"
		}

		if isExists {

			if messageData.ServoMessage.Id != "" && messageData.ServoMessage.Type == "write" {
				serialData := servoPlatform.ServoWriteId{}
				serialData.Type = "SERVO-WRITE-ID"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.OldId = 254
				serialData.NewId, _ = strconv.Atoi(messageData.ServoMessage.Id)

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}

				time.Sleep(40 * time.Millisecond)
			}

			if messageData.ServoMessage.Mode != "" && messageData.ServoMessage.Type == "write" {

				serialData := servoPlatform.ServoWriteMode{}
				serialData.Type = "SERVO-WRITE-MODE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialData.Mode, _ = strconv.Atoi(messageData.ServoMessage.Mode)
				serialData.Speed = 0
				if serialData.Mode == 1 {
					serialData.Speed, _ = strconv.Atoi(messageData.ServoMessage.Speed)
				}

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}

				time.Sleep(40 * time.Millisecond)
			}

			if messageData.ServoMessage.Value != "" && messageData.ServoMessage.Type == "write" {

				serialData := servoPlatform.ServoAngleOffsetWrite{}
				serialData.Type = "SERVO-ANGLE-OFFSET-WRITE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialData.Value, _ = strconv.Atoi(messageData.ServoMessage.Value)

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}

				time.Sleep(40 * time.Millisecond)
			}

			if messageData.ServoMessage.MinAngle != "" && messageData.ServoMessage.MaxAngle != "" && messageData.ServoMessage.Type == "write" {

				serialData := servoPlatform.ServoAngleLimitWrite{}
				serialData.Type = "SERVO-ANGLE-LIMIT-WRITE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialData.Min, _ = strconv.Atoi(messageData.ServoMessage.MinAngle)
				serialData.Max, _ = strconv.Atoi(messageData.ServoMessage.MaxAngle)

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}

				time.Sleep(40 * time.Millisecond)
			}

			if messageData.ServoMessage.MinVin != "" && messageData.ServoMessage.MaxVin != "" && messageData.ServoMessage.Type == "write" {

				serialData := servoPlatform.ServoVinLimitWrite{}
				serialData.Type = "SERVO-VIN-LIMIT-WRITE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialData.Min, _ = strconv.Atoi(messageData.ServoMessage.MinVin)
				serialData.Max, _ = strconv.Atoi(messageData.ServoMessage.MaxVin)

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}

				time.Sleep(60 * time.Millisecond)
			}

			if messageData.ServoMessage.Status != "" && messageData.ServoMessage.Type == "write" {

				serialData := servoPlatform.ServoStatusWrite{}
				serialData.Type = "SERVO-STATUS-WRITE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialData.Status, _ = strconv.Atoi(messageData.ServoMessage.Status)

				serialDataString, _ := json.Marshal(serialData)

				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
			}

			if messageData.ServoMessage.Type == "read" {

				sendMessage.MessageType = "servo_message_read"
				sendMessage.ServoMessageRead.Id = messageData.ServoMessage.Id
				sendMessage.ServoMessageRead.Channel = messageData.ServoMessage.Channel
				sendMessage.ServoMessageRead.Type = "read"

				serialData := servoPlatform.ServoReadData{}

				serialData.Type = "SERVO-READ-MODE"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialDataString, _ := json.Marshal(serialData)
				sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
				if sendStatus == true {
					readString := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)
					if readString != "" {

						log.Println("[info]", readString)

						readStringJson := servoPlatform.ServoWriteMode{}

						json.Unmarshal([]byte(readString), &readStringJson)

						sendMessage.ServoMessageRead.Mode = strconv.Itoa(readStringJson.Mode)
						sendMessage.ServoMessageRead.Speed = strconv.Itoa(readStringJson.Speed)
					}
				}

				time.Sleep(200 * time.Millisecond)

				serialData.Type = "SERVO-ANGLE-OFFSET-READ"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialDataString, _ = json.Marshal(serialData)
				sendStatus = serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
				if sendStatus == true {
					readString := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)
					if readString != "" {

						log.Println("[info]", readString)

						readStringJson := servoPlatform.ServoAngleOffsetWrite{}

						json.Unmarshal([]byte(readString), &readStringJson)

						sendMessage.ServoMessageRead.Value = strconv.Itoa(readStringJson.Value)
					}
				}

				time.Sleep(200 * time.Millisecond)

				serialData.Type = "SERVO-ANGLE-LIMIT-READ"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialDataString, _ = json.Marshal(serialData)
				sendStatus = serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
				if sendStatus == true {
					readString := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)
					if readString != "" {

						log.Println("[info]", readString)

						readStringJson := servoPlatform.ServoAngleLimitWrite{}

						json.Unmarshal([]byte(readString), &readStringJson)

						sendMessage.ServoMessageRead.MinAngle = strconv.Itoa(readStringJson.Min)
						sendMessage.ServoMessageRead.MaxAngle = strconv.Itoa(readStringJson.Max)
					}
				}

				time.Sleep(200 * time.Millisecond)

				serialData.Type = "SERVO-VIN-LIMIT-READ"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialDataString, _ = json.Marshal(serialData)
				sendStatus = serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
				if sendStatus == true {
					readString := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)
					if readString != "" {

						log.Println("[info]", readString)

						readStringJson := servoPlatform.ServoAngleLimitWrite{}

						json.Unmarshal([]byte(readString), &readStringJson)

						sendMessage.ServoMessageRead.MinVin = strconv.Itoa(readStringJson.Min)
						sendMessage.ServoMessageRead.MaxVin = strconv.Itoa(readStringJson.Max)
					}
				}

				time.Sleep(200 * time.Millisecond)

				serialData.Type = "SERVO-STATUS-READ"
				serialData.Channel, _ = strconv.Atoi(messageData.ServoMessage.Channel)
				serialData.Id, _ = strconv.Atoi(messageData.ServoMessage.Id)
				serialDataString, _ = json.Marshal(serialData)
				sendStatus = serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
				if sendStatus == false {
					sendMessage.MessageType = "servo_message_error"
					sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
				}
				if sendStatus == true {
					readString := serialFunction.SerialRead(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, Config.MentorConfig.RobotBoard.Bits)
					if readString != "" {

						log.Println("[info]", readString)

						readStringJson := servoPlatform.ServoStatusWrite{}

						json.Unmarshal([]byte(readString), &readStringJson)

						sendMessage.ServoMessageRead.Status = strconv.Itoa(readStringJson.Status)
					}
				}
			}
		}

		sendString, _ := json.Marshal(sendMessage)
		Send("", string(sendString))
	}

	if messageData.MessageType == "servo_message" {

		sendMessage := ResponseMessage{}

		isExists := CommonService.Exists(Config.MentorConfig.RobotBoard.Port)
		if isExists == false {
			sendMessage.MessageType = "servo_message_error"
			sendMessage.ServoMessageError.Content = "舵机通讯失败，请重新尝试"
		}

		if isExists == true {

			serialData := servoPlatform.ServoMotionWrite{}
			serialData.Type = "SERVO-MOTION-WRITE"
			serialData.Channel, _ = strconv.Atoi(messageData.ServoControlMessage.Channel)
			serialData.Id, _ = strconv.Atoi(messageData.ServoControlMessage.Id)
			serialData.Angle, _ = strconv.Atoi(messageData.ServoControlMessage.Angle)
			serialData.Time, _ = strconv.Atoi(messageData.ServoControlMessage.Time)

			serialDataString, _ := json.Marshal(serialData)

			sendStatus := serialFunction.SerialWrite(Config.MentorConfig.RobotBoard.Port, Config.MentorConfig.RobotBoard.Rate, string(serialDataString))
			if sendStatus == false {
				sendMessage.MessageType = "servo_message_error"
				sendMessage.ServoMessageError.Content = "舵机通讯数据发送失败，请重新尝试"
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