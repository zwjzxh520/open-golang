//go:build darwin
// +build darwin

package open

import (
	"os/exec"
)

func open(input string) *exec.Cmd {
	return exec.Command("open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command("open", "-a", appName, input)
}

func openWithArgs(input string, args []string) *exec.Cmd {
	args = append(args, input)
	return exec.Command("open", args...)
}

func revealItem(input string) *exec.Cmd {
	return exec.Command("open", "-R", input)
}
