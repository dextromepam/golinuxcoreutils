// chmod.go
// author: sergio@laszal.xyz
// date: 14/8/2023

package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
	"path/filepath"
)

func main() {
	var (
		recursive bool
		verbose   bool
	)

	// Define command-line flags
	flag.BoolVar(&recursive, "R", false, "Change files and directories recursively")
	flag.BoolVar(&verbose, "v", false, "Output a diagnostic for every file processed")

	// Parse command-line flags
	flag.Parse()

	// Ensure the correct number of arguments
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: chmod [-R] [-v] <mode> <file>")
		return
	}

	// Extract mode and file path from command-line arguments
	modeStr := flag.Args()[0]
	filePath := flag.Args()[1]

	// Parse the numeric mode specified by the user
	mode, err := parseNumericMode(modeStr)
	if err != nil {
		fmt.Println("Error parsing mode:", err)
		return
	}

	// Change permissions of the specified file
	if err := chmod(filePath, mode, recursive, verbose); err != nil {
		fmt.Println("Error:", err)
	}
}

// chmod changes the permissions of a file or directory
func chmod(filePath string, newMode os.FileMode, recursive bool, verbose bool) error {
	// Display a diagnostic message if verbose flag is set
	if verbose {
		fmt.Printf("Changing permissions of %s to %04o\n", filePath, newMode)
	}

	// Change permissions of the specified file
	if err := os.Chmod(filePath, newMode); err != nil {
		return err
	}

	// Recursively change permissions of directories if recursive flag is set
	if recursive {
		return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				// Display a diagnostic message if verbose flag is set
				if verbose {
					fmt.Printf("Changing permissions of %s to %04o\n", path, newMode)
				}

				// Change permissions of the directory
				if err := os.Chmod(path, newMode); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return nil
}

// parseNumericMode parses the numeric mode specified by the user
func parseNumericMode(modeStr string) (os.FileMode, error) {
	modeInt, err := strconv.ParseUint(modeStr, 8, 32)
	if err != nil {
		return 0, err
	}
	return os.FileMode(modeInt), nil
}
