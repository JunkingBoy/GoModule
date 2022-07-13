package main

import (
    "fmt"
    "sort"
)

/*
声明分类
将 int 声明为 HeroKind 英雄类型，后面会将这个类型当做枚举来使用
 */
type HeroKind int

/*
定义类型
 */
type Hero struct {
    Name string
    Kind HeroKind
}

/*
定义指针切片类型
 */
type Heros []*Hero

/*
定义常量--->类似枚举类
 */
const (
    None HeroKind = iota
    Tank
    Assassin
    Mage
)

/*
实现sort.Interface的方法
 */
// 取接口元素数量
func (s Heros) Len() int {
    return len(s)
}

// 比较接口元素
func (s Heros) Less(i, j int) bool {
    // 先判断分类是否一致，分类不一致优先对分类进行排序
    if s[i].Kind != s[j].Kind {
        // 对分类进行排序
        return s[i].Kind < s[j].Kind
    }

    // 默认按照名字升序排序
    return s[i].Name < s[j].Name
}

// 交换元素位置
func (s Heros) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

/*
调用接口
 */
func main() {
    // 准备数据
    heros := Heros{
        &Hero{"吕布", Tank},
        &Hero{"李白", Assassin},
        &Hero{"妲己", Mage},
        &Hero{"貂蝉", Assassin},
        &Hero{"关羽", Tank},
        &Hero{"诸葛亮", Mage},
    }

    // 调用sort包排序
    sort.Sort(heros)

    // 遍历列表打印结果
    for _, v := range heros {
        fmt.Printf("%v\n", v)
    }
}
