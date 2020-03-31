package serialFunction

import (
	"github.com/tarm/goserial"
	"log"
	"time"
)

func SerialWrite(Port string, Baud int, ReadTimeout time.Duration, Data string) bool {

	status := true

	serialConfig := &serial.Config{ Name: Port, Baud: Baud, ReadTimeout: ReadTimeout}

	serialOpen, err := serial.OpenPort(serialConfig)
	if err != nil {
		status = false
		return status
	}

	serialWrite, err := serialOpen.Write([]byte(Data))
	if err != nil {
		status = false
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient Serial Write Error", serialWrite)
		return status
	}

	return status
}