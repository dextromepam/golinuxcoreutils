// whoami.go
// author: sergio@laszal.xyz
// date 14/8/2023

package main

import(
	"fmt"
	"os"
	"os/user"
	"strconv"
)


func main() {
	uid := os.Getuid()

	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(u.Username)
}
