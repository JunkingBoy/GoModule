package Practice_Single

import "sync"

/*
使用sync包下的类型进行对象创建
 */
var once sync.Once

// 通过once变量执行接口当中的do方法创建对象
func GetInstanceOnce() *Tool {
    once.Do(func() {
        instance = new(Tool)
    })

    // 返回对象
    return instance
}
