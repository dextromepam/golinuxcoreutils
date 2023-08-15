// touch.go
// date 15/8/2023
// author sergio@laszal.xyz


package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Give a file")
		os.Exit(1)
	}

	fd, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %s\n", err)
		os.Exit(1)
	}
	defer fd.Close()

	os.Exit(0)
}
