package deno

import (
	"runtime"

	"github.com/pkg/errors"
)

// get Deno arch for current platform
func GetDenoArch() (*string, error) {
	var denoArch string

	switch runtime.GOARCH {
	case "amd64":
		fallthrough
	case "arm64":
		denoArch = "x64"
		break
	default:
		return nil, errors.New("not support your platform")
	}

	return &denoArch, nil
}
