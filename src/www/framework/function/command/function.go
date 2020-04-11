package commandFunction

import (
	"os/exec"
)

func Shell(shell string) (string, error){

	run, err := exec.Command("/bin/sh","-c", shell).Output()

	return string(run), err
}