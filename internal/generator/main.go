package generator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mrbns/assetloader/helper"
	conf "github.com/mrbns/assetloader/internal/config"
)

var (
	config         = conf.GetConfig()
	assets         = []string{} // Contain Only File names;
	projectRoot, _ = os.Getwd()
)

func isValidFiles(fileName string) bool {
	ext := filepath.Ext(fileName)

	switch ext {
	case ".png", ".jpg", ".jpeg", ".webp", ".avif", ".gif", ".mp4", ".svg":
		return true
	default:
		return false
	}
}

func getFiles() error {

	if config.AssetDir == "" {
		return errors.New("please provide asset directory eg. --dir=/src/assets")
	}

	fileDir := path.Join(projectRoot, config.AssetDir)
	entries, err := os.ReadDir(fileDir)
	helper.ErrorColorizedExit(err)
	if len(entries) <= 0 {
		return errors.New("no files found")
	}

	for _, ent := range entries {

		if isValidFiles(ent.Name()) {
			assets = append(assets, ent.Name())
		}
	}

	return nil
}

func GenerateAsset(directoryPath string) {

	//Checking Asset Directory
	if config.AssetDir == "" {
		fmt.Println(color.RedString("Error: please provide asset directory eg. --dir=/src/assets"))
		os.Exit(1)
		return
	}

	err := getFiles()
	helper.ErrorColorizedExit(err)

	projectRoot, _ := os.Getwd()
	outputScriptFile := filepath.Join(projectRoot, fmt.Sprintf("%v/%v", config.AssetDir, config.OutputFile))

	// Writing to file
	file, err := os.Create(outputScriptFile)
	helper.ErrorFatal(err, "")
	defer file.Close()

	fmt.Println("---- ------ BUILDING ------ ----")
	for index, name := range assets {
		ext := filepath.Ext(name)
		fileNameWithoutExtension := name[:len(name)-len(ext)]

		//Normalizing Name for JS variable
		normalizeNameRegex := regexp.MustCompile(`[-. !)(]`)
		normalizedFileName := normalizeNameRegex.ReplaceAllString(fileNameWithoutExtension, "_")

		if isValidFiles(name) {
			_, err := file.WriteString(fmt.Sprintf("export { default as %v_%v } from \"./%v\";\n", config.AssetPrefix, strings.ToUpper(normalizedFileName), name))
			fmt.Println(color.CyanString("%d - %s", index+1, name))
			helper.ErrorFatal(err, "")
		}
	}

	fmt.Println("---- ------- END ------ ----")
	fmt.Println(color.GreenString("Added %s file at %s", config.OutputFile, config.AssetDir))
	fmt.Println(color.HiBlueString("👋 Bye. see you again"))
}
