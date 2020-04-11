package commandFunction

import (
	"github.com/lizongshen/gocommand"
)

func Shell(command string) (int, string, error){

	pid, out, err := gocommand.NewCommand().Exec(command)

	return pid, out, err
}