package deno

import (
	"runtime"
)

// get Deno os for current platform
func GetDenoOS() (*string, error) {
	var denoOS string

	switch runtime.GOOS {
	case "darwin":
		denoOS = "osx"
		break
	case "windows":
		denoOS = "win"
		break
	default:
		// default to linux
		denoOS = "linux"
	}

	return &denoOS, nil
}
