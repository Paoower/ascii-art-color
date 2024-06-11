package main

import (
	asciiart "ascii-art-color/ascii-art"
	"fmt"
	"log"
	"os"
	"strings"
)

const NONE = "\033[0m"

func main() {
	colors := make(map[string]string)
	colors["RED"] = "\033[31;1m"
	colors["GREEN"] = "\033[32;1m"
	colors["YELLOW"] = "\033[33;1m"
	colors["ORANGE"] = "\033[33;1m"
	colors["BLUE"] = "\033[34;1m"
	colors["PURPLE"] = "\033[35;1m"
	colors["CYAN"] = "\033[36;1m"
	colors["GREY"] = "\033[37;1m"

	args := os.Args
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-color:")
	var lines []string

	if len(args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	}

	if len(args) == 1 {
		return
	}

	if len(args) == 2 {
		lines = asciiart.GetColoredAscii(args[1], "standard", "", "NONE")
	}

	if len(args) == 3 {
		if strings.HasPrefix(args[1], "--color") {
			lines = asciiart.GetColoredAscii(args[2], "standard", "", colors[strings.ToUpper(strings.TrimPrefix(args[1], "--color="))])
		} else {
			lines = asciiart.GetColoredAscii(args[1], args[2], "", "NONE")
		}
	}

	if len(args) == 4 {
		lines = asciiart.GetColoredAscii(args[3], "standard", args[2], colors[strings.ToUpper(strings.TrimPrefix(args[1], "--color="))])
	}

	if len(args) == 5 {
		lines = asciiart.GetColoredAscii(args[3], args[4], args[2], colors[strings.ToUpper(strings.TrimPrefix(args[1], "--color="))])
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
