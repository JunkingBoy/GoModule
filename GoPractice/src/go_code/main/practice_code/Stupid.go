package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("cmd", "/C", "shutdown", "-s", "-t", "0")

    /* 执行指令 */
    err := cmd.Start()
    if err != nil {
        fmt.Println(err.Error())
    }
}
