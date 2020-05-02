package cameraDriver

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func WebCamera(c *gin.Context){

	c.Header("Content-Type", "multipart/x-mixed-replace;boundary=frame")

	c.Header("Content-Length", strconv.Itoa(len(Camera.ReadFrame)))

	c.Writer.Write(Camera.ReadFrame)
}