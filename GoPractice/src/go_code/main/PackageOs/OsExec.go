package main

import (
    "fmt"
    "os/exec"
)

/*
调用os包下的exec包下的函数对操作系统进行操作
 */
func main() {
    f, err := exec.LookPath("main")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
}
