// cat.go
// author: sergio@laszal.xyz
// date: 14/8/2023

package main

import (
	"bufio"
	"strings"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		numberLines bool
		showEnds bool
		showTabs bool
		showAll bool
		noBuffer bool
	)
		
	flag.BoolVar(&numberLines, "n", false, "Number Lines")
	flag.BoolVar(&showEnds, "E", false, "Show ends of lines")
	flag.BoolVar(&showTabs, "T", false, "Show tabs")
	flag.BoolVar(&showAll, "A", false, "Show all non-printing characters")
	flag.BoolVar(&noBuffer, "u", false, "Disable output buffering")

	flag.Parse()

	// Disable output buffering if -u option is used
	if noBuffer {
		setNoBuffering()
	}

	// Retrieve non-option arguments
	nonOptionsArgs := flag.Args()

	if len(nonOptionsArgs) == 0 {
		// If no files are provided, process STDIN
		processInput(os.Stdin, numberLines, showEnds, showTabs, showAll)
	} else {	
		// Process each file specified as argument
		for _, fileName := range nonOptionsArgs {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}

			defer file.Close()
			
			processInput(file, numberLines, showEnds, showTabs, showAll)

		}
	}
	
}


// Disables output buffering
func setNoBuffering() {
	os.Stdout.Sync()
}


// Process input from a file or STDIN
func processInput(input *os.File, numberLines bool, showEnds bool, showTabs bool, showAll bool) {
	fileScanner := bufio.NewScanner(input)
	lineNumber := 1

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Show all non-printing characters if -A option is used
		if showAll {
			line = showNonPrintingCharacters(line)
		}
		
		// Number lines if -n option is used
		if numberLines {
			fmt.Printf("%6d ", lineNumber)
			lineNumber++ 
		}

		// Show TAB characters as ^I if -T option is used
		if showTabs {
			showTabCharacters(line)
		}

		// Show $ at the end of each line if -E option is used
		if showEnds {
			line = line + "$"
		}

		fmt.Println(line)
	}
	
}


// Replace TAB characters with ^I
func showTabCharacters(line string) string {
	line = strings.Replace(line, "\t", "^I", -1)
	return line
}


// Show all non-printing characters using ^x notation
func showNonPrintingCharacters(input string) string {
	var output strings.Builder
	for _, char := range input {
		if char < 32 || char == 127 {
			output.WriteString(fmt.Sprintf("^%c", char+64))
		} else {
			output.WriteRune(char)
		}
	}

	output.WriteString("$")

	return output.String()
}
