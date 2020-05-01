package cameraDriver

import (
	"encoding/base64"
	"gocv.io/x/gocv"
	"image"
)

type Driver struct {
	Camera *gocv.VideoCapture
	Status chan bool
	ReadFrame []byte
	ReadImage string
}

func StartDevice(Port interface{}, Status bool) (*Driver, error) {

	camera, err := gocv.OpenVideoCapture(Port)

	cameraImage := gocv.NewMat()

	cameraImageResize := gocv.NewMat()

	c := &Driver{}

	c.Camera = camera

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

				gocv.Resize(cameraImage, &cameraImageResize, image.Pt(cameraImage.Cols(), cameraImage.Rows()), 0, 0, gocv.InterpolationNearestNeighbor)

				frame, _ := gocv.IMEncode(gocv.JPEGFileExt, cameraImageResize)

				c.ReadFrame = frame

				c.ReadImage = base64.StdEncoding.EncodeToString(frame)

				if Status {

				}
			}
		}
	}()

	return c, err
}

func (c *Driver) OnClose() {
	c.Status <- true
}