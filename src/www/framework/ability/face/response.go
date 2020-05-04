package faceAbility

const FaceDetectUrl = "https://aip.baidubce.com/rest/2.0/face/v3/detect"

const FaceSearchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/multi-search"

var Face = &FaceDetectData{}

type FaceResult struct {
	ErrorCode 	int 			`json:"error_code"`
	Result 		FaceDetectData 	`json:"result"`
}

type FaceDetectData struct {
	AccessToken  	string
	Type 			string 		`json:"type"`
	FaceNum 		int 		`json:"face_num"`
	FaceList		[]FaceList 	`json:"face_list"`
}

type FaceList struct {
	FaceToken 		string 				`json:"face_token"`
	Location  		FaceLocation 		`json:"location"`
	FaceProbability float32 			`json:"face_probability"`
	UserList  		[]FaceUserList 		`json:"user_list"`
}

type FaceLocation struct {
	Left 	float32 `json:"left"`
	Top  	float32 `json:"top"`
	Width 	float32 `json:"width"`
	Height  float32 `json:"height"`
	Rotation float32 `json:"rotation"`
}

type FaceUserList struct {
	GroupId 	string `json:"group_id"`
	UserId  	string `json:"user_id"`
	UserInfo 	string `json:"user_info"`
	Score 		string `json:"score"`
}
