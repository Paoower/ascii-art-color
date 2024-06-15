package main

import (
	asciiart "ascii-art-color/ascii-art"
	utils "ascii-art-color/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

const NONE = "\033[0m"

func main() {
	args := os.Args
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-color:")
	var lines []string

	if len(args) > 5 || len(args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	}

	if len(args) == 2 {
		lines = asciiart.GetColoredAscii(args[1], "standard", "", "NONE")
	}

	if len(args) == 3 {
		if strings.HasPrefix(args[1], "--color=") {
			lines = asciiart.GetColoredAscii(args[2], "standard", "", utils.GetColor(args[1]))
		} else {
			lines = asciiart.GetColoredAscii(args[1], args[2], "", "NONE")
		}
	}

	color := utils.GetColor(args[1])

	if len(args) == 4 {
		lines = asciiart.GetColoredAscii(args[3], "standard", args[2], color)
	}

	if len(args) == 5 {
		lines = asciiart.GetColoredAscii(args[3], args[4], args[2], color)
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
