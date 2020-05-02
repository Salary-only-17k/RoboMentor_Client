package cameraDriver

import (
	"encoding/base64"
	"fmt"
	"github.com/webcam"
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
	var formats []webcam.PixelFormat
	for f := range formatDesc {
		formats = append(formats, f)
	}

	choice := readChoice(fmt.Sprintf("Choose format [1-%d]: ", len(formats)))

	format := formats[choice-1]

	camera.SetImageFormat(format, 1280, 720)

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
				}
			}
		}
	}()

	return Camera, err
}

func readChoice(s string) int {
	var i int
	for true {
		print(s)
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil || i < 1 {

		} else {
			break
		}
	}
	return i
}