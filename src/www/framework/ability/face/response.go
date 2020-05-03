package faceAbility

const FaceDetectUrl  = "https://aip.baidubce.com/rest/2.0/face/v3/detect"

var Face = &FaceDetectData{}

type FaceResult struct {
	ErrorCode 	int 			`json:"error_code"`
	Result 		FaceDetectData 	`json:"result"`
}

type FaceDetectData struct {
	AccessToken  	string
	FaceNum 		int 		`json:"face_num"`
	FaceList		[]FaceList 	`json:"face_list"`
}

type FaceList struct {
	FaceToken string `json:"face_token"`
	Location  FaceLocation `json:"location"`
}

type FaceLocation struct {
	Left 	float32 `json:"left"`
	Top  	float32 `json:"top"`
	Width 	float32 `json:"width"`
	Height  float32 `json:"height"`
}
