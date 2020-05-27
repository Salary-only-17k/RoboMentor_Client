package robot

import (
	"www/framework/platform/servo"
	"www/framework/robot"
)

type Base struct {
	Robot.Base
	RobotStatus chan bool
}

func InitRobot() Robot.Interface {
	return &Base {
		RobotStatus:make(chan bool),
	}
}

func (robot *Base) OnStart() {

	servoPlatform.StartPlatform(1,50)

	servoPlatform.SitDownAction()
}

func (robot *Base) OnClose() {

}

func (robot *Base) OnMessages(byte []byte, string string) {
	
}