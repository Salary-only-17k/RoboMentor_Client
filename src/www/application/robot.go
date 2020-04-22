package robot

import (
	"www/framework/robot"
)

// @Title  设置机器人基础配置
// @Description  可以定义一些机器人开发中需要的自定义变量数据
// @Robot.Base  预留，请务必保留
// @RobotStatus  自定义机器人状态变量，用户可自定义其他名称、数据类型
type Base struct {
	Robot.Base
	RobotStatus chan bool
}

// @Title  初始化机器人基础配置信息
// @Description	当机器人启动时，可以初始化一些数据，例如：当机器人启动时，将标记机器人状态RobotStatus的初始值设置为false
func InitRobot() Robot.Interface {
	return &Base {
		RobotStatus:make(chan bool),
	}
}

// @Title  机器人启动主程序
// @Description	用户可以在这里设计机器人的各种程序
func (robot *Base) OnStart() {

}

// @Title  机器人停止、意外停止时触发
// @Description	用户可以在这里设计机器人停止、意外停止时的程序
func (robot *Base) OnClose() {
	
}

// @Title  获取机器人本地、远程消息
// @Description	用户可以根据本地、远程发来消息设计一些机器人的程序，实现机器人的远程控制
// @byte  []byte格式数据
// @string  string字符串格式数据
func (robot *Base) OnMessages(byte []byte, string string) {

}