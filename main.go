package main

import (
	"os"

	"github.com/qbin-studio/assetloader/helper"
	"github.com/qbin-studio/assetloader/internal/generator"
)

func main() {

	args := os.Args

	for _, arg := range args {
		helper.ProcessArg(arg)
	}

	generator.GenerateAsset("")

}
