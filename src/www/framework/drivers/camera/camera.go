package cameraDriver

import (
	"encoding/base64"
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
			camera.SetImageFormat(code, 800, 450)
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
							time.Sleep(15 * time.Millisecond)
						}
					}
			}
		}
	}()

	return c, err
}


func StartVisionDevice(Port interface{}, Status bool) (*Driver, error) {

	camera, err := gocv.OpenVideoCapture(Port)

	cameraImage := gocv.NewMat()

	cameraImageResize := gocv.NewMat()

	c := &Driver{}

	prev := time.Now()

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

				timeElapsed := time.Now().Sub(prev)

				if float64(timeElapsed)/1000/1000/1000 >= 1.0/10.0 {

					gocv.Resize(cameraImage, &cameraImageResize, image.Pt(0, 0), 0, 0, gocv.InterpolationLinear)

					frame, _ := gocv.IMEncode(gocv.JPEGFileExt, cameraImageResize)

					c.ReadFrame = frame

					c.ReadImage = base64.StdEncoding.EncodeToString(frame)

					if Status {
						SocketService.RobotSocketClientSend("camera_message", c.ReadImage)
					}

					prev = time.Now()
				}
			}
		}
	}()

	return c, err
}

func (c *Driver) OnClose() {
	c.Status <- true
}