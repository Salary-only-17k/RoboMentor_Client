package servoPlatform

import (
	"encoding/json"
	"github.com/tarm/goserial"
	"io"
	"log"
	"strconv"
	"time"
	"www/framework/config"
)

var Servo = &Platform{}

type Platform struct {
	conn		io.ReadWriteCloser
	Port 		string
	Baud 		int
	Buf 		int
	Speed		int
	MotionMode 	int
	Action		[]ActionItem
}

type ActionItem struct {
	Channel int
	Id 		int
	Speed	int
	Angle 	int
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

func StartPlatform(MotionMode int, Speed int) {

	Servo.Baud, _ = strconv.Atoi(Config.MentorConfig.RobotBoard.Rate)
	Servo.Buf, _ = strconv.Atoi(Config.MentorConfig.RobotBoard.Bits)
	Servo.Port = Config.MentorConfig.RobotBoard.Port

	serialConfig := &serial.Config{ Name: Servo.Port, Baud: Servo.Baud, ReadTimeout: time.Millisecond * 10}

	Servo.conn , _ = serial.OpenPort(serialConfig)

	Servo.MotionMode = MotionMode

	Servo.Speed = Speed
}

func SetMotion() bool {

	setStatus := true

	if len(Servo.Action) > 0 {
		for _, v := range Servo.Action {

			sendData := ServoMotionWrite{}
			sendData.Type = "SERVO-MOTION-WRITE"
			sendData.Channel = v.Channel
			sendData.Id = v.Id
			sendData.Angle = v.Angle
			sendData.Time = Servo.Speed
			if	v.Speed > 0  {
				sendData.Time = v.Speed
			}

			if sendData.Angle != -1 {

				sendString, err := json.Marshal(sendData)
				if err != nil {
					log.Println("\033[31m[Error]\033[0m", "SetMotion Error", sendString)
				}

				serialWrite, err := Servo.conn.Write(sendString)
				if err != nil {
					log.Println("\033[31m[Error]\033[0m", "SetMotion Error", serialWrite)
				}

				time.Sleep(30 * time.Millisecond)
			}
		}

		Servo.Action = make([]ActionItem, 0)
	}

	return setStatus
}

func SitDownAction() bool {

	var action = []ActionItem{
		{Channel:1, Id:1, Speed:0, Angle:500},
		{Channel:2, Id:1, Speed:0, Angle:500},
		{Channel:1, Id:2, Speed:0, Angle:150},
		{Channel:2, Id:2, Speed:0, Angle:150},
		{Channel:1, Id:3, Speed:0, Angle:890},
		{Channel:2, Id:3, Speed:0, Angle:890},
		{Channel:1, Id:4, Speed:0, Angle:500},
		{Channel:2, Id:4, Speed:0, Angle:500},
		{Channel:1, Id:5, Speed:0, Angle:150},
		{Channel:2, Id:5, Speed:0, Angle:150},
		{Channel:1, Id:6, Speed:0, Angle:890},
		{Channel:2, Id:6, Speed:0, Angle:890},
		{Channel:1, Id:7, Speed:0, Angle:500},
		{Channel:2, Id:7, Speed:0, Angle:500},
		{Channel:1, Id:8, Speed:0, Angle:150},
		{Channel:2, Id:8, Speed:0, Angle:150},
		{Channel:1, Id:9, Speed:0, Angle:890},
		{Channel:2, Id:9, Speed:0, Angle:890},
	}

	Servo.Action = action

	return SetMotion()
}

func StandUpAction() bool {

	var action = []ActionItem{
		{Channel:1, Id:1, Speed:0, Angle:500},
		{Channel:1, Id:2, Speed:0, Angle:320},
		{Channel:1, Id:3, Speed:0, Angle:900},
		{Channel:1, Id:4, Speed:0, Angle:500},
		{Channel:1, Id:5, Speed:0, Angle:320},
		{Channel:1, Id:6, Speed:0, Angle:900},
		{Channel:1, Id:7, Speed:0, Angle:500},
		{Channel:1, Id:8, Speed:0, Angle:320},
		{Channel:1, Id:9, Speed:0, Angle:900},
		{Channel:2, Id:1, Speed:0, Angle:500},
		{Channel:2, Id:2, Speed:0, Angle:320},
		{Channel:2, Id:3, Speed:0, Angle:900},
		{Channel:2, Id:4, Speed:0, Angle:500},
		{Channel:2, Id:5, Speed:0, Angle:320},
		{Channel:2, Id:6, Speed:0, Angle:900},
		{Channel:2, Id:7, Speed:0, Angle:500},
		{Channel:2, Id:8, Speed:0, Angle:320},
		{Channel:2, Id:9, Speed:0, Angle:900},
	}

	Servo.Action = action

	return SetMotion()
}
