package main

import "fmt"

/* 声明一个字典结构 */
type Dictionary struct {
    // 设置一个map属性，key和value都是interface{}类型
    data map[interface{}]interface{}
}

/* 定义该类型的函数 */
func (d *Dictionary) Get(key interface{}) interface{} {
    return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
    d.data[key] = value
}

/* 设置回调函数，返回需要遍历的数据 */
// 定义回调，类型为 func(k,v interface{})bool，意思是返回键值数据（k、v）。bool 表示遍历流程控制，返回 true 时继续遍历，返回 false 时终止遍历。
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
    // 回调函数返回值为空，直接返回
    if callback == nil {
        return
    }

    // 循环取值
    for k, v := range d.data {
        // 根据 callback 的返回值，决定是否继续遍历
        if !callback(k, v) {
            return
        }
    }
}

/* 清空数据函数 */
func (d *Dictionary) Clear() {
    d.data = make(map[interface{}]interface{})
}

/* 创建一个字典 */
func NewDictionary() *Dictionary {
    d := &Dictionary{}

    // 初始化map-->先清理
    d.Clear()
    return d
}

func main() {
    // 创建字典实例
    dict := NewDictionary()

    // 添加数据
    dict.Set("大锤", 80)
    dict.Set("小锤", 60)
    dict.Set("瞎几把锤", 24)

    // 获取值并打印
    favorite := dict.Get("小锤")
    fmt.Println("价格：", favorite)

    // 遍历所有字典元素
    dict.Visit(func(key, value interface{}) bool {
        // 转换值得类型
        if value.(int) > 50 {
            // 输出内容
            fmt.Println(key, "贵的一批!")
            return true
        }

        // 默认输出
        fmt.Println(key, "不是很贵!")

        return true
    })
}
