package bodyAbility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"www/framework/service/ai"
	"www/framework/service/message"
	"www/framework/service/request"
)

func DetectBodyAnalysis(image string) (BodyDetectData, error) {

	bodyAttrDetectData := BodyDetectData{}

	paramsData := map[string]interface{}{"image":image}

	if Body.AccessToken == "" {
		Body.AccessToken = AiService.GetAccessToken("body")
	}

	Response, _ := RequestService.Post(BodyAnalysisUrl, paramsData, map[string]string{"access_token":Body.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := BodyResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		bodyAttrDetectData = responseData.Result
		bodyAttrDetectData.Type = "body"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(bodyAttrDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return bodyAttrDetectData, err
}

func DetectBodyAttr(image string) (BodyDetectData, error) {

	bodyAttrDetectData := BodyDetectData{}

	paramsData := map[string]interface{}{"image":image}

	if Body.AccessToken == "" {
		Body.AccessToken = AiService.GetAccessToken("body")
	}

	Response, _ := RequestService.Post(BodyAttrUrl, paramsData, map[string]string{"access_token":Body.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := BodyResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		bodyAttrDetectData = responseData.Result
		bodyAttrDetectData.Type = "body"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(bodyAttrDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return bodyAttrDetectData, err
}

func DetectBodyNum(image string) (BodyDetectData, error) {

	bodyAttrDetectData := BodyDetectData{}

	paramsData := map[string]interface{}{"image":image}

	if Body.AccessToken == "" {
		Body.AccessToken = AiService.GetAccessToken("body")
	}

	Response, _ := RequestService.Post(BodyNumUrl, paramsData, map[string]string{"access_token":Body.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := BodyResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		bodyAttrDetectData = responseData.Result
		bodyAttrDetectData.Type = "body"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(bodyAttrDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return bodyAttrDetectData, err
}