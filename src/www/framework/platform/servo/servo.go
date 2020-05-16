package servoPlatform

import "strconv"

var Servo = &Platform{}

type Platform struct {
	Status		chan bool
	Port 		string
	Baud 		int
	Buf 		int
}



func StartPlatform(Port string, Baud string, Buf string) {
	Servo.Baud, _ = strconv.Atoi(Baud)
	Servo.Buf, _ = strconv.Atoi(Buf)
	Servo.Port = Port
}