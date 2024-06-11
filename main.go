package main

import (
	asciiart "ascii-art-color/ascii-art"
	"fmt"
	"log"
	"os"
	"strings"
)

const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"
const PURPLE = "\033[35m"
const CYAN = "\033[36m"
const GREY = "\033[37m"
const NONE = "\033[0m"

// Basic case : go run . --color=<color> <substring to be colored> "something" - len 4
// Basic case w style : go run . --color=<color> <substring to be colored> "something" "banner" - len 5
// No sub : go run . --color=<color> "something" - len 3 OK
// Basic ascii art : go run . "Hello" - len 2 OK
// Ascii art with banner: go run . "Hello" standard - len 3 OK

func main() {
	colors := make(map[string]string)
	colors["RED"] = "\033[31m"
	colors["GREEN"] = "\033[32m"
	colors["YELLOW"] = "\033[33m"
	colors["ORANGE"] = "\033[33m"
	colors["BLUE"] = "\033[34m"
	colors["PURPLE"] = "\033[35m"
	colors["CYAN"] = "\033[36m"
	colors["GREY"] = "\033[37m"

	args := os.Args
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-color:")
	var lines []string

	if len(args) == 1 {
		return
	}

	if len(args) == 2 {
		lines = asciiart.GetColoredAscii(args[1], "standard", "", "NONE")
	}

	if len(args) == 3 {
		if strings.HasPrefix(args[1], "--color") {
			lines = asciiart.GetColoredAscii(args[2], "standard", "", strings.TrimPrefix(args[1], "--color="))
		} else {
			lines = asciiart.GetColoredAscii(args[1], args[2], "", "NONE")
		}
	}

	if len(args) == 4 {
		lines = asciiart.GetColoredAscii(args[3], "standard", args[2], strings.TrimPrefix(args[1], "--color="))
	}

	if len(args) == 5 {
		lines = asciiart.GetColoredAscii(args[3], args[4], args[2], strings.TrimPrefix(args[1], "--color="))
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
