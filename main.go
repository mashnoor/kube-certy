package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	cmd := exec.Command("echo", "-e", "hello")
	stdout, err := cmd.Output()

	if err == nil {
		fmt.Println(string(stdout))
	} else {
		fmt.Println(err)
	}
}
