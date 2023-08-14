// echo.go
// author: sergio@laszal.xyz
// date: 14/8/2023

package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
)

func main() {
	var (
		noNewline bool
		enableEscape bool
		disableBuffer bool
	)

	flag.BoolVar(&noNewline, "n", false, "Do not output the trailing newline")
	flag.BoolVar(&enableEscape, "e", false, "Enable interpretation of backslash escapes")
	flag.BoolVar(&disableBuffer, "u", false, "Disable output buffering")

	flag.Parse()
	
	if disableBuffer {
		setNoBuffering()
	}

	echoText := strings.Join(flag.Args(), " ")

	
	if enableEscape {
		echoText = interpretEscapes(echoText)
	}

	if !noNewline {
		echoText = echoText + "\n"
	}

	fmt.Printf(echoText)
}


// Disables output buffering
func setNoBuffering() {
	os.Stdout.Sync()
}

// Interprets escapes sequences
func interpretEscapes(input string) string {
	input = strings.Replace(input, "\\n", "\n", -1)
	input = strings.Replace(input, "\\t", "\t", -1)
	input = strings.Replace(input, "\\\"", "\"", -1)

	return input

}
