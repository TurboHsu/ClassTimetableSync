//go:build darwin

package auth

func openURL(url string) {
	exec.Command(`open`, url).Start()
}
