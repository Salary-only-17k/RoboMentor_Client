package CommonService

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
)

type EmptyData struct {

}

func Ping(c *gin.Context){
	Success(c, 0, "ok", EmptyData{})
	return
}

func Success(c *gin.Context, code int, msg string, data interface{}) {

	if msg == "" {
		msg = "请求成功"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	c.Set("code", code)
}

func Error(c *gin.Context, code int, msg string, data interface{}) {

	if msg == "" {
		msg = "请求失败"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func GetIntranetIp() string {

	Ip := "0.0.0.0"

	addRs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}

	for _, address := range addRs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				Ip = ipNet.IP.String()
			}
		}
	}

	return Ip
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsFile(path string) bool {
	return !IsDir(path)
}