package gestureAbility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"www/framework/service/ai"
	"www/framework/service/message"
	"www/framework/service/request"
)

func DetectGesture(image string) (GestureDetectData, error) {

	gestureDetectData := GestureDetectData{}

	paramsData := map[string]interface{}{"image":image}

	if Gesture.AccessToken == "" {
		Gesture.AccessToken = AiService.GetAccessToken("body")
	}

	Response, _ := RequestService.Post(GestureUrl, paramsData, map[string]string{"access_token":Gesture.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := GestureResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		gestureDetectData = responseData.Result
		gestureDetectData.Type = "gesture"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(gestureDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return gestureDetectData, err
}
