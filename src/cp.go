// cp.go
// Author: sergio@laszal.xyz
// Date: 14th August 2023

package main

import (
	"fmt"
	"path/filepath"
	"os"
	"io"
	"flag"
)

func main() {
	// Define command-line flags
	var (
		verbose   bool
		recursive bool
	)

	flag.BoolVar(&recursive, "R", false, "Copy directories recursively")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Parse()

	// Get non-option arguments (source and destination)
	nonOptionArgs := flag.Args()

	if len(nonOptionArgs) != 2 {
		fmt.Println("Usage: cp [-R] [-v] <source> <destination>")
		return
	}

	// Perform the copy operation
	err := copy(nonOptionArgs[0], nonOptionArgs[1], verbose, recursive)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Copy performs the file or directory copy operation.
func copy(source, destination string, verbose bool, recursive bool) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		if err := os.MkdirAll(destination, srcInfo.Mode()); err != nil {
			return err
		}

		entries, err := os.ReadDir(source)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			srcPath := filepath.Join(source, entry.Name())
			dstPath := filepath.Join(destination, entry.Name())

			if recursive {
				if err := copy(srcPath, dstPath, true, recursive); err != nil {
					return err
				}
			} else {
				if err := copyFile(srcPath, dstPath, verbose); err != nil {
					return err
				}
			}
		}
	} else {
		if err := copyFile(source, destination, verbose); err != nil {
			return err
		}
	}

	return nil
}

// copyFile copies a single file from source to destination.
func copyFile(sourceFile string, destinationFile string, verbose bool) error {
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	err = dst.Sync()
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Copied: '%s' -> '%s'\n", sourceFile, destinationFile)
	}

	return nil
}
