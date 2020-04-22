package cameraDriver

import (
	"encoding/base64"
	"github.com/webcam"
	"www/framework/service/socket"
)

type Driver struct {
	Camera *webcam.Webcam
	Status chan bool
	ReadFrame []byte
	ReadImage string
}

func StartDevice(Port string, Status bool) (*Driver, error) {

	camera, err := webcam.Open(Port)

	for code, formatName := range camera.GetSupportedFormats() {
		if formatName == "Motion-JPEG" {
			camera.SetImageFormat(code, 1280, 720)
		}
	}

	err = camera.StartStreaming()

	c := &Driver{}

	c.Camera = camera

	go func() {
		for {
			select {
				case <-c.Status:
					return
				default:
					frame, err := c.Camera.ReadFrame()
					if err != nil {
						continue
					}

					if len(frame) != 0 {
						c.ReadFrame = frame
						c.ReadImage = base64.StdEncoding.EncodeToString(frame)

						if Status {
							SocketService.RobotSocketClientSend("camera_message", c.ReadImage)
						}
					}
			}
		}
	}()

	return c, err
}

func (c *Driver) OnClose() {
	c.Status <- true
}