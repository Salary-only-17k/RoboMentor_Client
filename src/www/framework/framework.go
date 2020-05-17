package RoboMentor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www/framework/drivers/camera"
	"www/framework/service/system"
	"www/framework/service/template"
)

func Init() {
	SystemService.System()
}

func InitRouter() *gin.Engine {

	router := gin.New()

	gin.SetMode("debug")

	router.GET("/", TemplateService.Template)

	router.StaticFS("/static", http.Dir("template/dist/static"))

	router.GET("/camera", cameraDriver.WebCamera)

	router.GET("/video", cameraDriver.WebVideo)

	return router
}