package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "status")
	stdoutStderr, _ := cmd.CombinedOutput()
	fmt.Printf("+ %s\n", strings.Join(cmd.Args, " "))
	fmt.Printf("%s", stdoutStderr)
}
