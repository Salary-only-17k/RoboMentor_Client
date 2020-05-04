package gestureAbility

const GestureUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/gesture"

var Gesture = &GestureDetectData{}

type GestureResult struct {
	ErrorCode 	int 			`json:"error_code"`
	Result 		GestureDetectData 	`json:"result"`
}

type GestureDetectData struct {
	AccessToken  	string
	Type 			string 			`json:"type"`
	ResultNum 		int 			`json:"result_num"`
	Result			[]Result 		`json:"result"`
}

type Result struct {
	Probability float32 `json:"probability"`
	Top 		int 	`json:"top"`
	Height 		int 	`json:"height"`
	Classname   string 	`json:"classname"`
	Width       int 	`json:"width"`
	Left 		int 	`json:"left"`
}