// mv.go
// Author: sergio@laszal.xyz
// Date: 14th August 2023

package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
)

func main() {
	// Define command-line flags
	var (
		recursive bool
		verbose   bool
	)

	flag.BoolVar(&recursive, "R", false, "Move directories recursively")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Parse()

	// Get non-option arguments (source and destination)
	nonOptionArgs := flag.Args()

	if len(nonOptionArgs) != 2 {
		fmt.Println("Usage: mv [-R] [-v] <source> <destination>")
		return
	}

	// Perform the move operation
	err := move(nonOptionArgs[0], nonOptionArgs[1], recursive, verbose)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

// move moves a source file or directory to a destination.
func move(source, destination string, recursive, verbose bool) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		if !recursive {
			return fmt.Errorf("'%s' is a directory. Use -R flag to move directories recursively", source)
		}

		destination = filepath.Join(destination, filepath.Base(source))
		if err := os.MkdirAll(destination, srcInfo.Mode()); err != nil {
			return err
		}

		entries, err := os.ReadDir(source)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			entrySource := filepath.Join(source, entry.Name())
			entryDest := filepath.Join(destination, entry.Name())

			if err := move(entrySource, entryDest, true, verbose); err != nil {
				return err
			}
		}

		if err := os.Remove(source); err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Moved directory: '%s' -> '%s'\n", source, destination)
		}
	} else {
		err := os.Rename(source, destination)
		if err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Moved file: '%s' -> '%s'\n", source, destination)
		}
	}

	return nil
}
