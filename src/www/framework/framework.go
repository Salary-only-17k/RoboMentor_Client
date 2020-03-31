package RoboMentor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"www/framework/service/common"
	"www/framework/service/request"
)

type Config struct {
	Comment 		string 				`json:"_comment"`
	RobotName 		string 				`json:"robot_name"`
	RobotVersion 	string 				`json:"robot_version"`
	RobotApi     	string 				`json:"robot_api"`
	RobotService 	configRobotService 	`json:"robot_service"`
	RobotTemplate 	configRobotTemplate `json:"robot_template"`
	RobotRuntime	configRobotTemplate `json:"robot_runtime"`
	RobotAuth 		configRobotAuth 	`json:"robot_auth"`
	RobotMac 		string 				`json:"robot_mac"`
	RobotIp 		string 				`json:"robot_ip"`
	RobotPath 		string 				`json:"robot_path"`
	RobotCameraPort string 				`json:"robot_camera_port"`
	RobotBoardPort 	string 				`json:"robot_board_port"`
	RobotMicPort 	string 				`json:"robot_mic_port"`
	RobotRadarPort 	string 				`json:"robot_radar_port"`
}

type configRobotService struct {
	Comment 		string 	`json:"_comment"`
	Mode 			string 	`json:"mode"`
	HttpPort 		int 	`json:"http_port"`
}

type configRobotTemplate struct {
	Comment	string 	`json:"_comment"`
	Path 	string 	`json:"path"`
}

type configRobotAuth struct {
	Comment 	string 	`json:"_comment"`
	Status		bool 	`json:"status"`
	AppID 		string 	`json:"app_id"`
	AppSecret 	string 	`json:"app_secret"`
}

type responseData struct {
	Code 	int 			`json:"code"`
	Data 	responseRobot 	`json:"data"`
	Msg 	string 			`json:"msg"`
}

type responseRobot struct {
	Token       	string  `json:"token"`
	RobotType		int  	`json:"robot_type"`
	RobotTitle		string  `json:"robot_title"`
	RobotIp			string  `json:"robot_ip"`
	RobotMac		string  `json:"robot_mac"`
	RobotCameraPort	string  `json:"robot_camera_port"`
	RobotBoardPort	string  `json:"robot_board_port"`
	RobotMicPort	string  `json:"robot_mic_port"`
	RobotRadarPort	string  `json:"robot_radar_port"`
	RobotStatus		string  `json:"robot_status"`
}

var MentorConfig = &Config{}

func Init(AppID string, AppSecret string) {

	MentorConfig.RobotAuth.AppID = AppID

	MentorConfig.RobotAuth.AppSecret = AppSecret

	MentorConfig.RobotApi = "https://api.wileho.com"

	MentorConfig.RobotPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	MentorConfig.RobotIp = CommonService.GetIntranetIp()

	interfaces, _ := net.Interfaces()

	for _, inter := range interfaces {
		if inter.HardwareAddr.String() != "" {
			MentorConfig.RobotMac = inter.HardwareAddr.String()
		}
	}

	MentorConfig.RobotAuth.Status = false

	paramsData := map[string]string{"app_id":MentorConfig.RobotAuth.AppID, "app_secret": MentorConfig.RobotAuth.AppSecret, "robot_mac": MentorConfig.RobotMac,"robot_ip": MentorConfig.RobotIp}

	responseData := responseData{}

	Response, _ := RequestService.DefaultGet(MentorConfig.RobotApi + "/oauth/robot/register", paramsData)

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	if Response.StatusCode != 200 {
		log.Println("\033[31m[Error]\033[0m", "MentorClient Robot Register Error")
		return
	}

	err := json.Unmarshal(ResponseBody, &responseData)
	if err != nil {
		log.Println("\033[31m[Error]\033[0m", "MentorClient Robot Register Error")
		return
	}

	if responseData.Code != 0 {
		log.Println("\033[31m[Error]\033[0m", "MentorClient Robot Register Error")
		return
	}

	MentorConfig.RobotAuth.Status = true
	MentorConfig.RobotName = responseData.Data.RobotTitle
	MentorConfig.RobotCameraPort = responseData.Data.RobotCameraPort
	MentorConfig.RobotBoardPort = responseData.Data.RobotBoardPort
	MentorConfig.RobotMicPort = responseData.Data.RobotMicPort
	MentorConfig.RobotRadarPort = responseData.Data.RobotRadarPort
}