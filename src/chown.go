// chown.go
// author: sergio@laszal.xyz
// date: 15/8/2023

package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var (
		recursive bool
		verbose   bool
	)

	// Define command-line flags
	flag.BoolVar(&recursive, "R", false, "Change files and directories recursively")
	flag.BoolVar(&verbose, "v", false, "Output a diagnostic for every file processed")

	flag.Parse()

	// Ensure proper command-line arguments
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: chown [-R] [-v] <mode> <file>")
		return
	}

	// Extract user input and file path
	filePath := flag.Args()[1]
	inputStr := flag.Args()[0]

	// Parse user and group IDs
	uid, gid := parseID(inputStr)

	// Perform chown operation
	chown(filePath, uid, gid, verbose, recursive)
}

func chown(fileName string, uid int, gid int, verbose bool, recursive bool) {
	// Change ownership/group of the specified file
	if err := os.Chown(fileName, uid, gid); err != nil {
		fmt.Println("Error changing ownership/group:", err)
		os.Exit(1)
	}

	// Display diagnostic message if verbose mode is enabled
	if verbose {
		fmt.Printf("Ownership of '%s' retained\n", fileName)
	}

	// Apply chown recursively if specified
	if recursive {
		err := filepath.Walk(fileName, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Change ownership/group of directories
			if info.IsDir() {
				if err := os.Chown(path, uid, gid); err != nil {
					return err
				}

				// Display diagnostic message if verbose mode is enabled
				if verbose {
					fmt.Printf("Ownership of '%s' retained\n", path)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking directory:", err)
		}
	}
}

func parseID(input string) (int, int) {
	var (
		uid     int
		gid     int
	)

	// Split user input into user and group parts
	userInfoList := strings.Split(input, ":")

	// Get user ID
	u, err := user.Lookup(userInfoList[0])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	uid, _ = strconv.Atoi(u.Uid)

	// Get group ID if specified, otherwise use user's primary group
	if len(userInfoList) > 1 {
		g, err := user.LookupGroup(userInfoList[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		gid, _ = strconv.Atoi(g.Gid)
	} else {
		gid, _ = strconv.Atoi(u.Gid)
	}

	return uid, gid
}
