package main

import (
	"os"

	"github.com/mrbns/assetloader/helper"
	"github.com/mrbns/assetloader/internal/generator"
)

func main() {

	args := os.Args

	for _, arg := range args {
		helper.ProcessArg(arg)
	}

	generator.GenerateAsset("")

}
