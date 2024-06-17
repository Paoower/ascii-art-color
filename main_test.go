package main_test

import (
	asciiart "ascii-art-color/ascii-art"
	utils "ascii-art-color/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRedHW(t *testing.T) {
	input := "Hello world"
	color := "--color=red"

	fmt.Println("Input: ", input)
	fmt.Println("Color: ", strings.TrimPrefix(color, "--color="))

	cmd := exec.Command("go", "run", ".", color, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", input, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestGreen11(t *testing.T) {
	input := "1 + 1 = 2"
	color := "--color=green"

	fmt.Println("Input: ", input)
	fmt.Println("Color: ", strings.TrimPrefix(color, "--color="))

	cmd := exec.Command("go", "run", ".", color, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", input, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestYellowChars(t *testing.T) {
	input := "1 + 1 = 2"
	color := "--color=yellow"

	fmt.Println("Input: ", input)
	fmt.Println("Color: ", strings.TrimPrefix(color, "--color="))

	cmd := exec.Command("go", "run", ".", color, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", input, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestSecondToLastLetter(t *testing.T) {
	input := "welcome"
	substr := "elcome"
	color := "--color=green"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestLastLetter(t *testing.T) {
	input := "wlcome"
	substr := "e"
	color := "--color=purple"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestTwoLetters(t *testing.T) {
	input := "welcome"
	substr := "lc"
	color := "--color=orange"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestHeyGuys(t *testing.T) {
	input := "HeY GuYs"
	substr := "GuYs"
	color := "--color=orange"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRBG(t *testing.T) {
	input := "RGB()"
	substr := "B"
	color := "--color=blue"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomString1(t *testing.T) {
	input := "rdmDZoez"
	color := "--color=rgb(100, 14, 242)"

	cmd := exec.Command("go", "run", ".", color, input, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", input, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomString2(t *testing.T) {
	input := "hey35 there"
	color := "--color=rgb(52, 140, 167)"

	cmd := exec.Command("go", "run", ".", color, input, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", input, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomString3(t *testing.T) {
	input := "hey35 ther:;e"
	color := "--color=hsl(58, 79%, 50%)"
	substr := "3"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}

func TestRandomString4(t *testing.T) {
	input := "hey35"
	color := "--color=hsl(314, 80%, 24%)"
	substr := "35"

	cmd := exec.Command("go", "run", ".", color, substr, input)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)

	lines := asciiart.GetColoredAscii(input, "standard", substr, utils.GetColor(color))
	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	if str != s {
		t.Error("\033[31;1mOutput not valid\033[0m")
	} else {
		fmt.Println(str)
		fmt.Println("\033[32;1mValid output\033[0m")
	}
}
