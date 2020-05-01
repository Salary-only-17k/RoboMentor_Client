package cameraDriver

import "github.com/gin-gonic/gin"

func WebCamera(c *gin.Context){

	for {
		c.Writer.WriteString(string(Camera.ReadFrame))
	}
}