package main

import (
	"log"
	"os"
	"os/signal"
	"www/application"
	"www/framework/robot"
	"www/framework/service/socket"
)

func init() {
	SocketService.RobotSocketClient()
}

func main() {

	log.Println("[robot]", "Robot Start")

	Robot.Init = robot.InitRobot()
	Robot.Init.OnStart()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	Robot.Init.OnClose()

	log.Println("[robot]", "Robot Close")
}