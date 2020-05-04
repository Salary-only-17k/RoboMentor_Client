package speechAbility

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"www/framework/service/ai"
	"www/framework/service/message"
	"www/framework/service/request"
)

func DetectSpeech(format string, rate int, speech string, len int) (SpeechResult, error) {

	speechDetectData := SpeechResult{}

	paramsData := map[string]interface{}{"speech":speech,"format":format,"rate":rate,"channel":1,"cuid":"robomentor","len":len}

	if Speech.AccessToken == "" {
		Speech.AccessToken = AiService.GetAccessToken("speech")
	}

	Response, _ := RequestService.Post(SpeechDetectUrl, paramsData, map[string]string{"token":Speech.AccessToken}, map[string]string{}, "", "")

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	log.Println("[info]", string(ResponseBody))

	responseData := SpeechResult{}

	err := json.Unmarshal(ResponseBody, &responseData)

	if responseData.ErrNo == 0 {

		speechDetectData.Result = responseData.Result
		speechDetectData.Type = "speech"

		sendMessage := MessageService.ResponseMessage{}
		sendMessage.MessageType = "detect_message"
		contentString, _ := json.Marshal(speechDetectData)
		sendMessage.DetectMessage.Content = string(contentString)
		sendString, _ := json.Marshal(sendMessage)
		MessageService.Send("", string(sendString))
	}

	return speechDetectData, err
}

func Text2Audio(text string, aue string, per string) string {

	audioUrl := ""

	if Speech.AccessToken == "" {
		Speech.AccessToken = AiService.GetAccessToken("speech")
	}

	text = url.QueryEscape(url.QueryEscape(text))

	audioUrl = Text2AudioUrl + "?lan=zh&tok=" + Speech.AccessToken + "&tex=" + text + "&cuid=robomentor&ctp=1&per=" + per + "&aue=" + aue

	return audioUrl
}