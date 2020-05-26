package servoPlatform

import (
	"strconv"
)

var Servo = &Platform{}

type Platform struct {
	Status		chan bool
	Port 		string
	Baud 		int
	Buf 		int
}

type ServoReadData struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
}

type ServoWriteMode struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Mode 	int 	`json:"mode"`
	Speed 	int 	`json:"speed"`
}

type ServoWriteId struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	OldId 	int 	`json:"old_id"`
	NewId 	int 	`json:"new_id"`
}

type ServoAngleOffsetWrite struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Value 	int 	`json:"value"`
}

type ServoAngleLimitWrite struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Min 	int 	`json:"min"`
	Max 	int 	`json:"max"`
}

type ServoVinLimitWrite struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Min 	int 	`json:"min"`
	Max 	int 	`json:"max"`
}

type ServoMotionWrite struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Angle 	int 	`json:"angle"`
	Time 	int 	`json:"time"`
}

type ServoMotionStop struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
}

type ServoStatusWrite struct {
	Type 	string 	`json:"type"`
	Channel int 	`json:"channel"`
	Id 		int 	`json:"id"`
	Status 	int 	`json:"status"`
}

func StartPlatform(Port string, Baud string, Buf string) {
	Servo.Baud, _ = strconv.Atoi(Baud)
	Servo.Buf, _ = strconv.Atoi(Buf)
	Servo.Port = Port
}

