// pwd.go
// author: sergio@laszal.xyz
// date 15/8/2023

package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pwd)
}
