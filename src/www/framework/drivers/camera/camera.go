package cameraDriver

import (
	"encoding/base64"
	"github.com/webcam"
	"log"
	"www/framework/service/socket"
)

type Driver struct {
	Camera *webcam.Webcam
	Status chan bool
	ReadFrame []byte
	ReadImage string
}

func StartDevice(Port string) (*Driver, error) {

	log.Println("[robot]", 1)

	camera, err := webcam.Open(Port)

	log.Println("[robot]", 2)

	for code, formatName := range camera.GetSupportedFormats() {
		if formatName == "Motion-JPEG" {
			camera.SetImageFormat(code, 1280, 720)
		}
	}

	err = camera.StartStreaming()

	log.Println("[robot]", 3)

	c := &Driver{}

	c.Camera = camera
	c.Status<-false

	log.Println("[robot]", 4)

	go func() {
		log.Println("[robot]", 5)
		for {
			log.Println("[robot]", 6)
			select {
				case <-c.Status:
					log.Println("[robot]", c.ReadFrame)
					return
				default:
					frame, err := c.Camera.ReadFrame()
					if err != nil {
						continue
					}

					if len(frame) != 0 {
						c.ReadFrame = frame
						c.ReadImage = base64.StdEncoding.EncodeToString(frame)

						log.Println("[robot]", c.ReadFrame)

						SocketService.RobotSocketClientSend(c.ReadImage)
					}
			}
		}
	}()

	return c, err
}

func (c *Driver) OnClose() {
	c.Status <- true
}