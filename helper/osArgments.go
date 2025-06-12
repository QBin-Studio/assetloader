package helper

import (
	"strings"

	"github.com/qbin-studio/assetloader/internal/config"
)

func valueWithFallback(value, fallback string) string {
	if value != "" {
		return value
	} else {
		return fallback
	}
}

func ProcessArg(arg string) {
	splitArg := strings.Split(arg, "=")

	config := config.GetConfig()
	var flag, value string = splitArg[0], ""

	if len(splitArg) > 1 {
		value = splitArg[1]
	}

	if !strings.HasPrefix(flag, "--") {
		return
	}

	switch flag {

	case "--dir":

		config.AssetDir = valueWithFallback(value, "")
		return

	case "--prefix":
		config.AssetPrefix = valueWithFallback(value, "IMG_")
		return

	case "--output":
		config.OutputFile = valueWithFallback(value, "index.ts")
		return

	}

}
