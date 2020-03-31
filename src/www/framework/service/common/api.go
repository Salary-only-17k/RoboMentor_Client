package CommonService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"strconv"
	"time"
	"www/framework/function/serial"
)

type responseGetSystemInfo struct {
	Time    string 		`json:"time"`
	Memory 	float64 	`json:"memory"`
	Cpu 	[]float64 	`json:"cpu"`
}

func GetHomeIndex(c *gin.Context){

	Success(c, 0, "ok", EmptyData{})
	return
}

func GetSystemInfo(c *gin.Context){

	returnData := responseGetSystemInfo{}

	returnData.Time = time.Now().Format("15:04:05")

	memory, _ := mem.VirtualMemory()

	returnData.Memory = memory.UsedPercent

	cpuInfo, _ := cpu.Percent(time.Second, false)

	returnData.Cpu = cpuInfo

	Success(c, 0, "ok", returnData)
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

	sendStatus := serialFunction.SerialWrite(jsonData.Port, RateInt, time.Millisecond * 3, jsonData.Content)
	if sendStatus == false {
		Error(c, 10000, "串口数据发送失败，请求重新尝试", EmptyData{})
		return
	}

	Success(c, 0, "ok", EmptyData{})
	return
}