package asciiart

import (
	"fmt"
	"os"
	"strings"
)

const LETTER_HEIGHT = 8

func GetAscii(input, style string) string {
	bannerFile, err := getBannerFile(style)
	if err != nil {
		return ""
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		lines = append(lines, getWord(word, bannerFile)...)
	}

	str := ""
	for _, line := range lines {
		str += line + "\n"
	}

	return str
}

func fileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	// Normalize line endings to Unix-style
	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	return content
}

func getBannerFile(style string) (string, error) {
	switch style {
	case "standard":
		return "banners/standard.txt", nil
	case "shadow":
		return "banners/shadow.txt", nil
	case "thinkertoy":
		return "banners/thinkertoy.txt", nil
	default:
		return "", fmt.Errorf("unknown style: %s. Available styles: standard, shadow, thinkertoy", style)
	}
}

func getWord(input string, bannerFile string) []string {
	lines := make([]string, 8)
	content := fileOpen(bannerFile)

	for _, char := range input {
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}

	return lines
}

func GetLetter(content string, ascii int) string {
	if ascii == 32 {
		s := ""
		for i := 0; i < 8; i++ {
			if i != 7 {
				s += "    " + "\n"
				continue
			}
			s += "    "
		}
		return s
	}

	str := ""
	lines := strings.Split(content, "\n")

	place := ascii - 31
	times := (place - 1) * LETTER_HEIGHT
	beginning := (ascii - 30) + times

	for i := beginning; i < beginning+LETTER_HEIGHT; i++ {
		if i != (beginning+LETTER_HEIGHT)-1 {
			str += lines[i-1] + "\n"
		} else {
			str += lines[i-1]
		}
	}

	return str
}
