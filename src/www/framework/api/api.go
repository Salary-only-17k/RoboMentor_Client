package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
	"www/framework/config"
	"www/framework/function/serial"
	"www/framework/service/common"
	"www/framework/service/request"
	"www/framework/service/socket"
)

type requestGetHomeIndex struct {
	Goos 		string `json:"goos"`
	Goarch 		string `json:"goarch"`
	HostName 	string `json:"host_name"`
	Version		string `json:"version"`
	Config 		*Config.Config `json:"config"`
}

func GetHomeIndex(c *gin.Context){

	returnData := requestGetHomeIndex{}

	returnData.Goos = runtime.GOOS
	returnData.Goarch = runtime.GOARCH

	hostName, _ := os.Hostname()

	returnData.HostName = hostName

	returnData.Version = runtime.Version()
	returnData.Config = Config.MentorConfig

	CommonService.Success(c, 0, "ok", returnData)
	return
}

type requestGetHomeRobot struct {
	Config 		*Config.Config `json:"config"`
}

func GetHomeRobot(c *gin.Context){

	returnData := requestGetHomeRobot{}

	returnData.Config = Config.MentorConfig

	CommonService.Success(c, 0, "ok", returnData)
	return
}

func SetHomeRobotSubmit(c *gin.Context){

	name := c.DefaultQuery("name", "")
	port := c.DefaultQuery("port", "")
	rate := c.DefaultQuery("rate", "")
	bits := c.DefaultQuery("bits", "")

	if name == "" {
		CommonService.Error(c, 10000, "请求失败，缺少必要参数", CommonService.EmptyData{})
		return
	}

	if port != "" {
		isExists := CommonService.Exists(port)
		if isExists == false {
			CommonService.Error(c, 10000, "没有找到相关串口，请重新尝试", CommonService.EmptyData{})
			return
		}

		RateInt, _ := strconv.Atoi(rate)

		sendStatus := serialFunction.SerialWrite(port, RateInt, "ping")
		if sendStatus == false {
			CommonService.Error(c, 10000, "串口配置不正确，请重新尝试", CommonService.EmptyData{})
			return
		}
	}

	responseData := RequestService.ResponseData{}

	paramsData := map[string]string{"name":name, "port": port, "bits": bits, "rate": rate, "token": Config.MentorConfig.Token}

	Response, _ := RequestService.DefaultGet(Config.MentorConfig.RobotApi + "/oauth/robot/update", paramsData, map[string]string{"Robot-Token":Config.MentorConfig.RobotAuth.AppID + "@" + Config.MentorConfig.RobotAuth.AppSecret + "@" + strconv.Itoa(int(time.Now().Unix()))})

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	if Response.StatusCode != 200 {
		CommonService.Error(c, 10000, "请求失败，请重新尝试", CommonService.EmptyData{})
		return
	}

	err := json.Unmarshal(ResponseBody, &responseData)
	if err != nil {
		CommonService.Error(c, 10000, "请求失败，请重新尝试", CommonService.EmptyData{})
		return
	}

	if responseData.Code != 0 {
		CommonService.Error(c, 10000, "请求失败，请重新尝试", CommonService.EmptyData{})
		return
	}

	Config.MentorConfig.RobotName = name
	Config.MentorConfig.RobotBoard.Port = port
	Config.MentorConfig.RobotBoard.Bits = bits
	Config.MentorConfig.RobotBoard.Rate = rate

	CommonService.Success(c, 0, "ok", CommonService.EmptyData{})
	return
}

type requestGetHomeTools struct {
	Config 		*Config.Config `json:"config"`
}

func GetHomeTools(c *gin.Context){

	returnData := requestGetHomeTools{}

	returnData.Config = Config.MentorConfig

	CommonService.Success(c, 0, "ok", returnData)
	return
}

type requestSetHomeToolsRemoteSubmit struct {
	Content string 		`json:"content"`
}

func SetHomeToolsRemoteSubmit(c *gin.Context){

	postJson, _ := ioutil.ReadAll(c.Request.Body)

	jsonData := requestSetHomeToolsRemoteSubmit{}

	err := json.Unmarshal(postJson, &jsonData)
	if err != nil {
		CommonService.Error(c, 10000, "请求失败，请求参数错误", CommonService.EmptyData{})
		return
	}

	readData := SocketService.SocketMessage{}

	readData.MessageType = "remote_message"
	readData.RemoteMessage.Content = jsonData.Content

	SocketService.Channel.Channel = readData

	CommonService.Success(c, 0, "ok", CommonService.EmptyData{})
	return
}

type requestSetHomeToolsSerialSubmit struct {
	Port    string 		`json:"port"`
	Rate 	string 		`json:"rate"`
	Bits 	string 		`json:"bits"`
	Content string 		`json:"content"`
	Switch 	bool 		`json:"Switch"`
}

func SetHomeToolsSerialSubmit(c *gin.Context){

	postJson, _ := ioutil.ReadAll(c.Request.Body)

	jsonData := requestSetHomeToolsSerialSubmit{}

	err := json.Unmarshal(postJson, &jsonData)
	if err != nil {
		CommonService.Error(c, 10000, "请求失败，请求参数错误", CommonService.EmptyData{})
		return
	}

	isExists := CommonService.Exists(jsonData.Port)
	if isExists == false {
		CommonService.Error(c, 10000, "没有找到相关串口，请重新尝试", CommonService.EmptyData{})
		return
	}

	RateInt, _ := strconv.Atoi(jsonData.Rate)

	sendStatus := serialFunction.SerialWrite(jsonData.Port, RateInt, jsonData.Content)
	if sendStatus == false {
		CommonService.Error(c, 10000, "串口数据发送失败，请重新尝试", CommonService.EmptyData{})
		return
	}

	if jsonData.Switch {

		readContent := serialFunction.SerialRead(jsonData.Port, RateInt, 5)

		readData := SocketService.SocketMessage{}

		readData.MessageType = "serial_message"
		readData.SerialMessage.Port = jsonData.Port
		readData.SerialMessage.Content = readContent

		SocketService.Channel.Channel = readData
	}

	CommonService.Success(c, 0, "ok", CommonService.EmptyData{})
	return
}

type requestSetHomeToolsTcpSubmit struct {
	Ip    	string 		`json:"ip"`
	Write 	string 		`json:"write"`
	Read 	string 		`json:"read"`
	Content string 		`json:"content"`
	Switch 	bool 		`json:"Switch"`
}

func SetHomeToolsTcpSubmit(c *gin.Context){

	postJson, _ := ioutil.ReadAll(c.Request.Body)

	jsonData := requestSetHomeToolsTcpSubmit{}

	err := json.Unmarshal(postJson, &jsonData)
	if err != nil {
		CommonService.Error(c, 10000, "请求失败，请求参数错误", CommonService.EmptyData{})
		return
	}

	if jsonData.Ip != "" && jsonData.Write != "" && jsonData.Content != "" {

		conn, err := net.DialTimeout(jsonData.Ip, jsonData.Write, 2*time.Second)
		if err != nil {
			CommonService.Error(c, 10000, "请求失败，请求重新尝试", CommonService.EmptyData{})
			return
		}

		defer conn.Close()

		writeInt, err := conn.Write([]byte(jsonData.Content))
		if err != nil {
			CommonService.Error(c, 10000, "数据发送失败，请求重新尝试", writeInt)
			return
		}

		if jsonData.Switch {

			reply := make([]byte, 1024)

			replyInt, err := conn.Read(reply)
			if err != nil {
				CommonService.Error(c, 10000, "数据读取失败，请求重新尝试", replyInt)
				return
			}

			readData := SocketService.SocketMessage{}

			readData.MessageType = "tcp_message"
			readData.TcpMessage.Content = string(reply)

			SocketService.Channel.Channel = readData
		}
	}

	if jsonData.Ip != "" && jsonData.Read != "" && jsonData.Content != "" && jsonData.Switch {

		conn, err := net.DialTimeout(jsonData.Ip, jsonData.Read, 5*time.Second)
		if err != nil {
			CommonService.Error(c, 10000, "请求失败，请求重新尝试", CommonService.EmptyData{})
			return
		}

		defer conn.Close()

		reply := make([]byte, 1024)

		replyInt, err := conn.Read(reply)
		if err != nil {
			CommonService.Error(c, 10000, "数据读取失败，请求重新尝试", replyInt)
			return
		}

		readData := SocketService.SocketMessage{}

		readData.MessageType = "tcp_message"
		readData.TcpMessage.Content = string(reply)

		SocketService.Channel.Channel = readData
	}

	CommonService.Success(c, 0, "ok", CommonService.EmptyData{})
	return
}