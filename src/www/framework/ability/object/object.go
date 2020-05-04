package objectAbility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"www/framework/service/ai"
	"www/framework/service/message"
	"www/framework/service/request"
)

func DetectObject(image string, Url string) (ObjectDetectData, error) {

	objectDetectData := ObjectDetectData{}

	paramsData := map[string]interface{}{"image":image}

	if Object.AccessToken == "" {
		Object.AccessToken = AiService.GetAccessToken("object")
	}

	Response, _ := RequestService.Post(Url, paramsData, map[string]string{"access_token":Object.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := ObjectResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		objectDetectData = responseData.Result
		objectDetectData.Type = "object"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(objectDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return objectDetectData, err
}