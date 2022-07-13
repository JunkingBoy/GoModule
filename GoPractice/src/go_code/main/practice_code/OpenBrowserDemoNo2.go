package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(`explorer`,`https://www.cnblogs.com/JunkingBoy/`)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}