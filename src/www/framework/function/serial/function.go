package serialFunction

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func SerialWrite(Port string, Baud string, Data string) bool {

	status := true

	isExists := exists(Port)
	if isExists == false {
		status = false
		return status
	}

	RateInt, _ := strconv.Atoi(Baud)

	serialConfig := &serial.Config{ Name: Port, Baud: RateInt, ReadTimeout: time.Millisecond * 10}

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

func SerialRead(Port string, Baud string, Buf string) string {

	stringData := ""

	isExists := exists(Port)
	if isExists == false {
		return stringData
	}

	RateInt, _ := strconv.Atoi(Baud)

	bits, _ := strconv.Atoi(Buf)

	serialConfig := &serial.Config{ Name: Port, Baud: RateInt, ReadTimeout: 128}

	serialOpen, err := serial.OpenPort(serialConfig)
	if err != nil {
		stringData = ""
	}

	readBuf := make([]byte, bits)

	for {
		serialRead, err := serialOpen.Read(readBuf)
		if err != nil {
			break
		}else{
			if strings.Index(stringData, "}") > 0 {
				stringData += fmt.Sprintf("%s", string(readBuf[:serialRead]))
				break
			}else{
				stringData += fmt.Sprintf("%s", string(readBuf[:serialRead]))
			}
		}
	}

	return stringData
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}