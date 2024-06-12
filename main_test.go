package main_test

import (
	asciiart "ascii-art-color/ascii-art"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestRedHW(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=red", "Hello world")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	input := "Hello world"
	color := "RED"
	lines := asciiart.GetColoredAscii(input, "standard", input, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestGreen12(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=green", "1 + 1 = 2")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	color := "GREEN"
	input := "1 + 1 = 2"
	lines := asciiart.GetColoredAscii(input, "standard", input, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestYellowSpecialChars(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=yellow", "(%&) ??")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	color := "YELLOW"
	input := "(%&) ??"
	lines := asciiart.GetColoredAscii(input, "standard", input, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomStringColor1(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=blue", "abcDEF")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	color := "BLUE"
	input := "abcDEF"
	lines := asciiart.GetColoredAscii(input, "standard", input, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomStringColor2(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "--color=purple", "abc  123")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	color := "PURPLE"
	input := "abc  123"
	lines := asciiart.GetColoredAscii(input, "standard", input, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomStringSubstr1(t *testing.T) {
	color := "CYAN"
	input := "hello12,:?"
	substr := "12"

	cmd := exec.Command("go", "run", ".", "--color=cyan", substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomStringSubstr2(t *testing.T) {
	input := "efUUEF;;,//)hello12,:?"
	substr := "//)"

	cmd := exec.Command("go", "run", ".", "--color=green", substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	color := "GREEN"

	lines := asciiart.GetColoredAscii(input, "standard", substr, getColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func getColor(c string) string {
	colors := make(map[string]string)
	colors["RED"] = "\033[31;1m"
	colors["GREEN"] = "\033[32;1m"
	colors["YELLOW"] = "\033[33;1m"
	colors["ORANGE"] = "\033[33;1m"
	colors["BLUE"] = "\033[34;1m"
	colors["PURPLE"] = "\033[35;1m"
	colors["CYAN"] = "\033[36;1m"
	colors["GREY"] = "\033[37;1m"

	co := ""

	for k, v := range colors {
		if k == c {
			return v
		}
	}

	return co
}
