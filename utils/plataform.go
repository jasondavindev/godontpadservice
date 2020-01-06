package utils

import (
	"os/exec"
	"runtime"
)

func GetPlataformCommand(cmd string) *exec.Cmd {
	plat := runtime.GOOS

	if plat == "linux" {
		return exec.Command("sh", "-c", cmd)
	}

	if plat == "windows" {
		return exec.Command("cmd", "/C", cmd)
	}

	return nil
}
