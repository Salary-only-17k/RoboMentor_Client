package cameraDriver

import (
	"encoding/base64"
	"fmt"
	"github.com/webcam"
	"gocv.io/x/gocv"
	"image"
	"time"
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
			camera.SetImageFormat(code, 1000, 562)
		}else{
			camera.SetImageFormat(code, 640, 480)
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
							time.Sleep(20 * time.Millisecond)
						}
					}
			}
		}
	}()

	return c, err
}


func StartDeviceOpenCV(Port interface{}, Status bool) (*Driver, error) {

	camera, _ := gocv.OpenVideoCapture(Port)

	cameraImage := gocv.NewMat()

	cameraImageResize := gocv.NewMat()

	c := &Driver{}

	go func() {
		for {
			select {
			case <-c.Status:
				return
			default:

				if ok := camera.Read(&cameraImage); !ok {
					continue
				}

				if cameraImage.Empty() {
					continue
				}

				gocv.Resize(cameraImage, &cameraImageResize, image.Pt(1000, 562), 0, 0, gocv.InterpolationLinear)

				frame, _ := gocv.IMEncode(gocv.JPEGFileExt, cameraImageResize)

				c.ReadFrame = frame

				c.ReadImage = base64.StdEncoding.EncodeToString(frame)

				if Status {
					SocketService.RobotSocketClientSend("camera_message", c.ReadImage)
					time.Sleep(20 * time.Millisecond)
				}

			}
		}
	}()
}

func (c *Driver) OnClose() {
	c.Status <- true
}