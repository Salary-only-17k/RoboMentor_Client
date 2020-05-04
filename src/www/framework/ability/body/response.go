package bodyAbility

const BodyAttrUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_attr"

const BodyAnalysisUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_analysis"

const BodyNumUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_num"

var Body = &BodyDetectData{}

type BodyResult struct {
	ErrorCode 	int 			`json:"error_code"`
	Result 		BodyDetectData 	`json:"result"`
}

type BodyDetectData struct {
	AccessToken  	string
	Type 			string 			`json:"type"`
	PersonNum 		int 			`json:"person_num"`
	PersonInfo		[]PersonInfo 	`json:"person_info"`
}

type PersonInfo struct {
	Location  		BodyLocation 		`json:"location"`
	BodyParts		BodyParts 			`json:"body_parts"`
	Attributes      BodyAttributes 		`json:"attributes"`
}

type BodyLocation struct {
	Left 	float32 	`json:"left"`
	Top  	float32 	`json:"top"`
	Width 	float32 	`json:"width"`
	Height  float32 	`json:"height"`
	Rotation float32 	`json:"rotation"`
}

type BodyParts struct {
	LeftHip 		BodyPartsData `json:"left_hip"`
	TopHead 		BodyPartsData `json:"top_head"`
	RightMouthCorner BodyPartsData `json:"right_mouth_corner"`
	Neck 			BodyPartsData `json:"neck"`
	LeftShoulder 	BodyPartsData `json:"left_shoulder"`
	LeftKnee 		BodyPartsData `json:"left_knee"`
	LeftAnkle 		BodyPartsData `json:"left_ankle"`
	LeftMouthCorner BodyPartsData `json:"left_mouth_corner"`
	RightElbow 		BodyPartsData `json:"right_elbow"`
	RightEar 		BodyPartsData `json:"right_ear"`
	Nose 			BodyPartsData `json:"nose"`
	LeftEye 		BodyPartsData `json:"left_eye"`
	RightEye 		BodyPartsData `json:"right_eye"`
	RightHip 		BodyPartsData `json:"right_hip"`
	LeftWrist 		BodyPartsData `json:"left_wrist"`
	LeftEar 		BodyPartsData `json:"left_ear"`
	LeftElbow 		BodyPartsData `json:"left_elbow"`
	RightShoulder 	BodyPartsData `json:"right_shoulder"`
	RightAnkle 		BodyPartsData `json:"right_ankle"`
	RightKnee 		BodyPartsData `json:"right_knee"`
	RightWrist 		BodyPartsData `json:"right_wrist"`
}

type BodyAttributes struct {
	Orientation 		OrientationData `json:"orientation"`
	IsHuman 			OrientationData `json:"is_human"`
	HeadWear 			OrientationData `json:"headwear"`
	UpperWearTexture 	OrientationData `json:"upper_wear_texture"`
	CarryingItem 		OrientationData `json:"carrying_item"`
	FaceMask 			OrientationData `json:"face_mask"`
	LowerWear 			OrientationData `json:"lower_wear"`
	Vehicle 			OrientationData `json:"vehicle"`
	UpperWearFg 		OrientationData `json:"upper_wear_fg"`
	LowerColor 			OrientationData `json:"lower_color"`
	Umbrella 			OrientationData `json:"umbrella"`
	UpperCut 			OrientationData `json:"upper_cut"`
	LowerCut 			OrientationData `json:"lower_cut"`
	Cellphone 			OrientationData `json:"cellphone"`
	Gender 				OrientationData `json:"gender"`
	Age 				OrientationData `json:"age"`
	Bag 				OrientationData `json:"bag"`
	Smoke 				OrientationData `json:"smoke"`
	UpperColor 			OrientationData `json:"upper_color"`
	Occlusion 			OrientationData `json:"occlusion"`
	Glasses 			OrientationData `json:"glasses"`
}

type BodyPartsData struct {
	Y float32 `json:"y"`
	X float32 `json:"x"`
	Score float32 `json:"score"`
}

type OrientationData struct {
	Name 	string 	`json:"name"`
	Score 	float32 `json:"score"`
}