//go:build !windows && !darwin

package open

import (
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
func openWithArgs(input string, args []string) *exec.Cmd {
	args = append(args, input)
	return exec.Command("xdb-open", args...)
}

func revealItem(input string) *exec.Cmd {
	return open(input)
}
