package AiService

import (
	"encoding/json"
	"io/ioutil"
	"www/framework/config"
	"www/framework/service/request"
)

type responseData struct {
	Code 	int 			`json:"code"`
	Data 	reToken 		`json:"data"`
	Msg 	string 			`json:"msg"`
}

type reToken struct {
	AccessToken  string	`json:"access_token"`
}

func GetAccessToken(Type string) string {

	accessToken := ""

	paramsData := map[string]string{"token":Config.MentorConfig.Token, "type": Type}

	responseData := responseData{}

	Response, _ := RequestService.Get(Config.MentorConfig.RobotApi + "/oauth/robot/ai/access/token", paramsData, map[string]string{}, Config.MentorConfig.RobotAuth.AppID, Config.MentorConfig.RobotAuth.AppSecret)

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	json.Unmarshal(ResponseBody, &responseData)

	if responseData.Code == 0 {
		accessToken = responseData.Data.AccessToken
	}

	return accessToken
}