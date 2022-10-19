//go:build windows

package auth

import (
	"os/exec"
	"syscall"
)

func openURL(url string) {
	cmd := exec.Command("cmd", "/C", "start", url)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_ = cmd.Start()
}
