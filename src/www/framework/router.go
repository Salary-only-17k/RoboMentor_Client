package RoboMentor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www/application"
	"www/framework/service/common"
	"www/framework/service/socket"
	"www/framework/service/template"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	gin.SetMode("debug")

	r.GET("/", TemplateService.Template)

	r.StaticFS("/static", http.Dir("template/dist/static"))

	r.GET("/message", SocketService.Socket)

	r.GET("/common/check", CommonService.Ping)

	r.GET("/common/home/index", Api.GetHomeIndex)

	r.GET("/common/home/robot", Api.GetHomeRobot)

	r.GET("/common/home/robot/submit", Api.SetHomeRobotSubmit)

	r.GET("/common/home/tools", Api.GetHomeTools)

	r.POST("/common/home/tools/serial/submit", Api.SetHomeToolsSerialSubmit)

	return r
}
