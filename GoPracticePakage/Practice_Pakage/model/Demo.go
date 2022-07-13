package model

import "fmt"

/* 定义一个不可以被外包访问的常量 */
const canNotFind = "Lucifer"

/* 定义一个可以被外包访问的常量 */
const CanFind = "JunkingBoy"

/* 设置包外不可以访问的结构体 */
type canNotFindStruct struct {
}

/* 设置包外可以访问的结构体 */
type CanFindStruct struct {
}

/* 设置包外不可访问的接口 */
type canFindInter interface {
}

/* 设置包外可访问的接口 */
type CanFindInter interface {
}

/* 在该包下定义一个打印的函数 */
func PrintStr() {
    fmt.Println("WelCome To Lucifer's Home!")
}
