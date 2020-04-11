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

	router := gin.New()

	gin.SetMode("debug")

	router.GET("/", TemplateService.Template)

	router.StaticFS("/static", http.Dir("template/dist/static"))

	router.GET("/message", SocketService.WebSocket)

	router.GET("/common/check", CommonService.Ping)

	router.GET("/common/home/index", api.GetHomeIndex)

	router.GET("/common/home/robot", api.GetHomeRobot)

	router.GET("/common/home/robot/submit", api.SetHomeRobotSubmit)

	router.GET("/common/home/tools", api.GetHomeTools)

	router.POST("/common/home/tools/serial/submit", api.SetHomeToolsSerialSubmit)

	router.POST("/common/home/tools/remote/submit", api.SetHomeToolsRemoteSubmit)

	router.POST("/common/home/tools/tcp/submit", api.SetHomeToolsTcpSubmit)

	return router
}