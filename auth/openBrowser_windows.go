//go:build windows

package auth

import (
	"os/exec"
	"strings"
	"syscall"
)

func openURL(url string) {
	url = strings.ReplaceAll(url, "&", "^&")
	cmd := exec.Command("cmd", "/C", "start", url)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_ = cmd.Start()
}
