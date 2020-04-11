package commandFunction

import (
	"os/exec"
)

func Shell(command string) (string, error){

	run := exec.Command("/bin/sh", "-c", command)
	bytes, err := run.Output()

	return string(bytes), err
}