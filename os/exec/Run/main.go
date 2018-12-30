package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("curl", "www.baidu.com")
	err := cmd.Run()
	fmt.Printf("+ %s\n", strings.Join(cmd.Args, " "))
	if err != nil {
		fmt.Println(err)
		return
	}
}
