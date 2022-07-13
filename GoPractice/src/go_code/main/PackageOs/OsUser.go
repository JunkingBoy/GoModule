package main

import (
    "log"
    "os/user"
)

/*
通过os/user包下的函数获取到user对应的信息
 */
func main() {
    // 调用user包下的current函数
    u, _ := user.Current()

    // 获取user结构体下的属性
    log.Println("用户名:", u.Username)
    log.Println("用户id:", u.Uid)
    log.Println("用户主目录:", u.HomeDir)
    log.Println("用户主组id:", u.Gid)

    // 用户所在所有组的id--->循环获取
    s, _ := u.GroupIds()
    log.Println("用户所在所有组:", s)
}
