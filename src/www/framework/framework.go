package RoboMentor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www/framework/api"
	"www/framework/service/common"
	"www/framework/service/socket"
	"www/framework/service/template"
)

func Init() {}

func InitRouter() *gin.Engine {

	r := gin.New()

	gin.SetMode("debug")

	r.GET("/", TemplateService.Template)

	r.StaticFS("/static", http.Dir("template/dist/static"))

	r.GET("/message", SocketService.WebSocket)

	r.GET("/common/check", CommonService.Ping)

	r.GET("/common/home/index", api.GetHomeIndex)

	r.GET("/common/home/robot", api.GetHomeRobot)

	r.GET("/common/home/robot/submit", api.SetHomeRobotSubmit)

	r.GET("/common/home/tools", api.GetHomeTools)

	r.POST("/common/home/tools/serial/submit", api.SetHomeToolsSerialSubmit)

	return r
}