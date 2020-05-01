package cameraDriver

import (
	"encoding/base64"
	"github.com/webcam"
)

var Camera = &Driver{}

type Driver struct {
	Camera 			*webcam.Webcam
	Status 			chan bool
	ReadFrame 		[]byte
	ReadImage 		string
	AccessToken 	string
}

func StartDevice(Port string, visionType string) (*Driver, error) {

	camera, err := webcam.Open(Port)

	for code, formatName := range camera.GetSupportedFormats() {
		if formatName == "Motion-JPEG" {
			camera.SetImageFormat(code, 1000, 562)
		}
	}

	err = camera.StartStreaming()

	Camera := &Driver{}

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

					if visionType != "" {

					}
				}
			}
		}
	}()

	return Camera, err
}