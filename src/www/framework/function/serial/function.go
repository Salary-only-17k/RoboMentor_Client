package serialFunction

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"strings"
	"time"
)

/*
Example usage:

  package main

  import (
        "github.com/tarm/goserial"
        "log"
  )

  func main() {
        c := &serial.Config{Name: "COM5", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }

        n, err := s.Write([]byte("test"))
        if err != nil {
                log.Fatal(err)
        }

        buf := make([]byte, 128)
        n, err = s.Read(buf)
        if err != nil {
                log.Fatal(err)
        }
        log.Print("%q", buf[:n])
  }
*/

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
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient Serial Write Error", serialWrite)
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
			if strings.Index(stringData, "\n") > 0 {
				break
			}else{
				stringData += fmt.Sprintf("%s", string(readBuf[:serialRead]))
			}
		}
	}

	return stringData
}

