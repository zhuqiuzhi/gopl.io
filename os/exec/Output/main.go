package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("go", "version")
	version, err := cmd.Output()
	fmt.Printf("+ %s\n", strings.Join(cmd.Args, " "))
	if err != nil {
		fmt.Printf("%s", cmdErr(err))
		return
	}
	fmt.Printf("%s", version)
}

func cmdErr(err error) string {
	if ee, ok := err.(*exec.ExitError); ok && len(ee.Stderr) > 0 {
		return fmt.Sprintf("%s: %s", err, ee.Stderr)
	}
	return fmt.Sprint(err)
}
