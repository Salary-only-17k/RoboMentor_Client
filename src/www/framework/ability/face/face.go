package faceAbility

import (
	"encoding/json"
	"io/ioutil"
	"www/framework/service/ai"
	"www/framework/service/message"
	"www/framework/service/request"
)

func DetectFace(image string) (FaceDetectData, error) {

	faceDetectData := FaceDetectData{}

	paramsData := map[string]interface{}{"image":image,"image_type":"BASE64","max_face_num":5}

	if Face.AccessToken == "" {
		Face.AccessToken = AiService.GetAccessToken("face")
	}

	Response, _ := RequestService.Post(FaceDetectUrl, paramsData, map[string]string{"access_token":Face.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	responseData := FaceResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrorCode == 0 {

		faceDetectData = responseData.Result

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(faceDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return faceDetectData, err
}

