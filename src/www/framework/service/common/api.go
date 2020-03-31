package CommonService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
	"www/framework/function/serial"
	"www/framework/service/socket"
)

func GetHomeIndex(c *gin.Context){

	Success(c, 0, "ok", EmptyData{})
	return
}

func GetHomeTools(c *gin.Context){

	Success(c, 0, "ok", EmptyData{})
	return
}

type requestSetHomeToolsSerial struct {
	Port    string 		`json:"port"`
	Rate 	string 		`json:"rate"`
	Bits 	string 		`json:"bits"`
	Content string 		`json:"content"`
	Switch 	bool 		`json:"Switch"`
}

func SetHomeToolsSerial(c *gin.Context){

	postJson, _ := ioutil.ReadAll(c.Request.Body)

	jsonData := requestSetHomeToolsSerial{}

	err := json.Unmarshal(postJson, &jsonData)
	if err != nil {
		Error(c, 10000, "请求失败，请求参数错误", EmptyData{})
		return
	}

	isExists := Exists(jsonData.Port)
	if isExists == false {
		Error(c, 10000, "没有找到相关串口，请求重新尝试", EmptyData{})
		return
	}

	RateInt, _ := strconv.Atoi(jsonData.Rate)

	sendStatus := serialFunction.SerialWrite(jsonData.Port, RateInt, jsonData.Content)
	if sendStatus == false {
		Error(c, 10000, "串口数据发送失败，请求重新尝试", EmptyData{})
		return
	}

	if jsonData.Switch {

		readContent := serialFunction.SerialRead(jsonData.Port, RateInt, 128, time.Millisecond * 500)

		readData := SocketService.SocketMessage{}

		readData.MessageType = "serial_log"
		readData.SerialMessage.Port = jsonData.Port
		readData.SerialMessage.Content = readContent

		SocketService.Channel.Channel = readData
	}

	Success(c, 0, "ok", EmptyData{})
	return
}