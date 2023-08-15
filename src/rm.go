// rm.go
// Author: sergio@laszal.xyz
// Date: 14th August 2023

package main

import (
	"fmt"
	"path/filepath"
	"os"
	"flag"
)

func main() {
	// Define command-line flags
	var (
		recursive bool
		verbose   bool
	)

	flag.BoolVar(&recursive, "R", false, "Remove directories recursively")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Parse()

	// Get non-option arguments (targets to be removed)
	nonOptionArgs := flag.Args()

	if len(nonOptionArgs) == 0 {
		fmt.Println("Usage: rm [-R] [-v] <target> [<target> ...]")
		return
	}

	// Perform the removal operation
	err := removeTargets(nonOptionArgs, recursive, verbose)

	if err != nil {
		fmt.Println("Error:", err)
	}
}

// removeTargets removes the specified targets (files or directories).
func removeTargets(targets []string, recursive bool, verbose bool) error {
	for _, target := range targets {
		err := remove(target, recursive, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}

// remove removes a single target (file or directory).
func remove(target string, recursive bool, verbose bool) error {
	targetInfo, err := os.Stat(target)
	if err != nil {
		return err
	}

	if targetInfo.IsDir() {
		if !recursive {
			return fmt.Errorf("'%s' is a directory. Use -R flag to remove directories recursively", target)
		}

		entries, err := os.ReadDir(target)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			entryPath := filepath.Join(target, entry.Name())

			if err := remove(entryPath, true, verbose); err != nil {
				return err
			}
		}

		if err := os.Remove(target); err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Removed directory: '%s'\n", target)
		}
	} else {
		if err := os.Remove(target); err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Removed file: '%s'\n", target)
		}
	}

	return nil
}
