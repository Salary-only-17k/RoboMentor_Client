package Router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www/framework"
	"www/framework/service/common"
	"www/framework/service/socket"
	"www/framework/service/template"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	gin.SetMode(RoboMentor.MentorConfig.RobotService.Mode)

	r.GET("/", TemplateService.Template)

	r.StaticFS("/static", http.Dir("template/dist/static"))

	r.GET("/message", SocketService.Socket)

	r.GET("/common/check", CommonService.Ping)

	r.GET("/common/home/index", CommonService.GetHomeIndex)

	r.GET("/common/home/system/info", CommonService.GetSystemInfo)

	r.GET("/common/home/tools", CommonService.GetHomeTools)

	r.POST("/common/home/tools/serial", CommonService.SetHomeToolsSerial)

	return r
}
