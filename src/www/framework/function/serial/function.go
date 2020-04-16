package serialFunction

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"strings"
	"time"
)

func SerialWrite(Port string, Baud int, Data string) bool {

	status := true

	serialConfig := &serial.Config{ Name: Port, Baud: Baud, ReadTimeout: time.Millisecond * 10}

	serialOpen, err := serial.OpenPort(serialConfig)
	if err != nil {
		status = false
		return status
	}

	serialWrite, err := serialOpen.Write([]byte(Data))
	if err != nil {
		status = false
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK Serial Write Error", serialWrite)
		return status
	}

	return status
}

func SerialRead(Port string, Baud int, Buf int) string {

	stringData := ""

	serialConfig := &serial.Config{ Name: Port, Baud: Baud, ReadTimeout: time.Millisecond * 500}

	serialOpen, err := serial.OpenPort(serialConfig)
	if err != nil {
		stringData = ""
	}

	readBuf := make([]byte, Buf)

	for {
		serialRead, err := serialOpen.Read(readBuf)
		if err != nil {
			stringData = ""
		}else{
			log.Println("[info]", stringData)
			if strings.Index(stringData, "\r\n") > 0 {
				break
			}else{
				stringData += fmt.Sprintf("%s", string(readBuf[:serialRead]))
			}
		}
	}

	return stringData
}

