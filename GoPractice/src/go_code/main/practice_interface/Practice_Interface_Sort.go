package main

import (
    "fmt"
    "sort"
)

/*
将[]string定义为类型
Go的接口是隐式的，所以定义类型需要实现接口当中的方法
接口实现不受限于结构体，任何类型都可以实现接口。要排序的字符串切片 []string 是系统定制好的类型，无法让这个类型去实现 sort.Interface 排序接口。因此，需要将 []string 定义为自定义的类型
 */
type MyStringList []string

// 实现sort.Interface接口获取元素数量的方法
func (m MyStringList) Len() int {
    return len(m)
}

// 实现sort.Interface比较元素的方法
func (m MyStringList) Less(i, j int) bool {
    return m[i] < m[j]
}

// 实现sort.Interface交换元素的方法
func (m MyStringList) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func main() {
    // 准备一个内容被打乱的字符串切片
    /*
    由于将 []string 定义成 MyStringList 类型，字符串切片初始化
     */
    names := MyStringList{
        "3, 这是三",
        "4, 这是四",
        "5, 这是五",
        "1, 这是一",
        "2, 这是二",
    }

    // 使用sort包进行排序
    sort.Sort(names)

    // 遍历打印结果
    for _, v := range names {
        fmt.Printf("%s\n", v)
    }
}
