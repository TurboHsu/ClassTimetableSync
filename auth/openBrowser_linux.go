//go:build linux

package auth

func openURL(url string) {
	exec.Command(`xdg-open`, url).Start()
}
