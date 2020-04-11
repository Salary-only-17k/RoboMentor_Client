package commandFunction

import (
	"github.com/lizongshen/gocommand"
)

func Shell(command string) (int, string, error){

	pid, out, err := gocommand.NewCommand().Exec("cd /robot/RoboMentor_SDK", "export GOPATH=$PWD", "cd /robot/RoboMentor_SDK/src/www", "go build robot.go")

	return pid, string(out), err
}