package cameraDriver

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func WebCamera(c *gin.Context){

	c.Header("Content-Type", "image/jpeg")

	c.Header("Content-Length", strconv.Itoa(len(Camera.ReadFrame)))

	c.Writer.WriteString(string(Camera.ReadFrame))
}