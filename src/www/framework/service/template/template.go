package TemplateService

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func Template(c *gin.Context) {

	Action, _ := template.ParseFiles("template/dist/index.html")

	err := Action.Execute(c.Writer, nil)
	if err != nil {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient Template Error")
	}
}