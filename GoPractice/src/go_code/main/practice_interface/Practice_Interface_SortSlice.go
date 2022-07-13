package main

import (
    "fmt"
    "sort"
)

/*
定义分类
 */
type herokind int

/*
定义枚举常数
 */
const (
    none herokind = iota
    tank
    assassin
    mage
)

/*
定义类型
 */
type hero struct {
    name string
    kind herokind
}

/*
调用
 */
func main() {
    heros := []*hero{
        {"吕布", tank},
        {"李白", assassin},
        {"妲己", mage},
        {"貂蝉", assassin},
        {"关羽", tank},
        {"诸葛亮", mage},
    }

    // 调用sort.Slice函数
    sort.Slice(heros, func(i, j int) bool {
        if heros[i].kind != heros[j].kind {
            return heros[i].kind < heros[j].kind
        }
        return heros[i].name < heros[j].name
    })

    // 循环输出切片中的值
    for _, v := range heros {
        fmt.Printf("%v\n", v)
    }
}
