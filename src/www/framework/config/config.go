package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"www/framework/service/common"
	"www/framework/service/request"
)

type Config struct {
	Token       	string  			`json:"token"`
	RobotName 		string 				`json:"robot_name"`
	RobotVersion 	string 				`json:"robot_version"`
	RobotApi     	string 				`json:"robot_api"`
	RobotMessage    string 				`json:"robot_message"`
	RobotRuntime	string 				`json:"robot_runtime"`
	RobotAuth 		configRobotAuth 	`json:"robot_auth"`
	RobotMac 		string 				`json:"robot_mac"`
	RobotIp 		string 				`json:"robot_ip"`
	RobotNetIp		string  			`json:"robot_net_ip"`
	RobotPath 		string 				`json:"robot_path"`
	RobotBoard 		configRobotBoard 	`json:"robot_board"`
}

type configRobotAuth struct {
	Status		bool 	`json:"status"`
	AppID 		string 	`json:"app_id"`
	AppSecret 	string 	`json:"app_secret"`
}

type configRobotBoard struct {
	Port		string 	`json:"port"`
	Rate 		string 	`json:"rate"`
	Bits 		string 	`json:"bits"`
}

type responseData struct {
	Code 	int 			`json:"code"`
	Data 	responseRobot 	`json:"data"`
	Msg 	string 			`json:"msg"`
}

type responseRobot struct {
	Token       	string  			`json:"token"`
	RobotType		int  				`json:"robot_type"`
	RobotTitle		string  			`json:"robot_title"`
	RobotIp			string  			`json:"robot_ip"`
	RobotNetIp		string  			`json:"robot_net_ip"`
	RobotMac		string  			`json:"robot_mac"`
	RobotBoard 		configRobotBoard 	`json:"robot_board"`
}

var MentorConfig = &Config{}

func Init(AppID string, AppSecret string) {

	MentorConfig.RobotAuth.AppID = AppID

	MentorConfig.RobotAuth.AppSecret = AppSecret

	MentorConfig.RobotApi = "https://api.robomentor.cn"

	MentorConfig.RobotMessage = "tcp://47.111.96.253:1883"

	MentorConfig.RobotVersion = "1.0.0"

	MentorConfig.RobotPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	MentorConfig.RobotIp = CommonService.GetIntranetIp()

	MentorConfig.RobotRuntime = "runtime"

	interfaces, _ := net.Interfaces()

	for _, inter := range interfaces {
		if inter.HardwareAddr.String() != "" {
			MentorConfig.RobotMac = inter.HardwareAddr.String()
		}
	}

	MentorConfig.RobotAuth.Status = true

	paramsData := map[string]string{"app_id":MentorConfig.RobotAuth.AppID, "app_secret": MentorConfig.RobotAuth.AppSecret, "robot_mac": MentorConfig.RobotMac,"robot_ip": MentorConfig.RobotIp}

	responseData := responseData{}

	Response, _ := RequestService.DefaultGet(MentorConfig.RobotApi + "/oauth/robot/register", paramsData, map[string]string{"Robot-Token":MentorConfig.RobotAuth.AppID + "@" + MentorConfig.RobotAuth.AppSecret + "@" + strconv.Itoa(int(time.Now().Unix()))})

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	if Response.StatusCode != 200 {
		MentorConfig.RobotAuth.Status = false
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK Register Error")
	}

	err := json.Unmarshal(ResponseBody, &responseData)
	if err != nil {
		MentorConfig.RobotAuth.Status = false
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK Register Error")
	}

	if responseData.Code != 0 {
		MentorConfig.RobotAuth.Status = false
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK Register Error")
	}

	MentorConfig.Token = responseData.Data.Token
	MentorConfig.RobotName = responseData.Data.RobotTitle
	MentorConfig.RobotNetIp = responseData.Data.RobotNetIp
	MentorConfig.RobotBoard.Port = responseData.Data.RobotBoard.Port
	MentorConfig.RobotBoard.Bits = responseData.Data.RobotBoard.Bits
	MentorConfig.RobotBoard.Rate = responseData.Data.RobotBoard.Rate
}