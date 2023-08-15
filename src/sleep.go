// sleep.go
// author: sergio@laszal.xyz
// date: 15/8/2023

package main

import(
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	
	if len(os.Args) != 2 {
		fmt.Println("Usage: sleep <time>")
		return
	}

	timeSec, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}	

	duration := time.Duration(timeSec) * time.Second
	
	time.Sleep(duration)
	

	
}
