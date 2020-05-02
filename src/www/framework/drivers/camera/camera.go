package cameraDriver

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/webcam"
	"strconv"
	"time"
)

var Camera = &Driver{}

type Driver struct {
	Camera 			*webcam.Webcam
	Status 			chan bool
	ReadFrame 		[]byte
	ReadImage 		string
}

func StartDevice(Port string) (*Driver, error) {

	camera, err := webcam.Open(Port)

	formatDesc := camera.GetSupportedFormats()
	for code, formatName := range formatDesc {
		if formatName == "Motion-JPEG" {
			camera.SetImageFormat(code, 1280, 720)
		}
	}

	err = camera.StartStreaming()

	Camera.Camera = camera

	go func() {
		for {
			select {
			case <-Camera.Status:
				return
			default:
				frame, err := Camera.Camera.ReadFrame()
				if err != nil {
					continue
				}

				if len(frame) != 0 {
					Camera.ReadFrame = frame
					Camera.ReadImage = base64.StdEncoding.EncodeToString(frame)

					time.Sleep(20 * time.Millisecond)
				}
			}
		}
	}()

	return Camera, err
}

func WebCamera(c *gin.Context){

	c.Header("Content-Type", "image/jpeg")

	c.Header("Content-Length", strconv.Itoa(len(Camera.ReadFrame)))

	c.Writer.WriteString(string(Camera.ReadFrame))
}