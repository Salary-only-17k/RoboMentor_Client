package commandFunction

import (
	"os/exec"
)

func Shell(command string) (string, error){

	run, err := exec.Command(command).Output()

	return string(run), err
}