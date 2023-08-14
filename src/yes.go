// yes.go
// author sergio@laszal.xyz
// date 14/8/2023

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var argString string
	
	// Display yes by default
	if len(os.Args) == 1 {
		argString = "yes"
	} else {
		argString = strings.Join(os.Args[1:], " ")
	}
	
	for {
		fmt.Println(argString)
	}
}
