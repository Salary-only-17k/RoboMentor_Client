package objectAbility

var Object = &ObjectDetectData{}

type ObjectResult struct {
	ErrorCode 	int 			`json:"error_code"`
	Result 		ObjectDetectData 	`json:"result"`
}

type ObjectDetectData struct {
	AccessToken  	string
	Type 			string 			`json:"type"`
	ResultNum 		int 			`json:"result_num"`
	Results			[]Results 		`json:"results"`
}

type Results struct {
	Name 		string 		`json:"name"`
	Score 		float32 	`json:"score"`
	Location    Location 		`json:"location"`
}

type Location struct {
	Left 	float32 `json:"left"`
	Top  	float32 `json:"top"`
	Width 	float32 `json:"width"`
	Height  float32 `json:"height"`
}