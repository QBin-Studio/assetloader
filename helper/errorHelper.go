package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Error struct {
}

func ErrorFatal(err error, point string) {
	if point != "" {
		println(point)
	}
	if err != nil {
		log.Fatal(err)
	}
}
func WarnErrorPanic(err error) {
	if err != nil {
		fmt.Println(color.YellowString(
			fmt.Sprint(err),
		))
	}

}

func ErrorColorizedExit(err error) {
	if err != nil {
		fmt.Println(color.RedString(
			fmt.Sprintf("Error: %v", err),
		))
		os.Exit(0)
	}
}
