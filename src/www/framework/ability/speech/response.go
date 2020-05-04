package speechAbility

const SpeechDetectUrl = "https://vop.baidu.com/server_api"

const Text2AudioUrl = "https://tsn.baidu.com/text2audio"

var Speech = &SpeechResult{}

type SpeechResult struct {
	ErrNo 		int 			`json:"err_no"`
	Result 		[]string 		`json:"result"`
	AccessToken  	string
	Type 			string 		`json:"type"`
}