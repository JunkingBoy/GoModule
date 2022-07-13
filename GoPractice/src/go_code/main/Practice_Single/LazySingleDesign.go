package Practice_Single

import "sync"

/*
设置结构体
建立私有变量用来返回结构体
提供获取该结构体实例的函数
 */
/* 设置结构体 */
type Tool struct {
    // 定义属性
    values int
}

/* 设置私有变量用于获取该结构体实例 */
/*
这样做方便切换实例化的结构体
 */
var instance *Tool

/* 进行加锁操作 */
var lock sync.Mutex

/* 定义构建私有变量实例的函数 */
func GetInstance() *Tool {
    // 先上锁
    lock.Lock()

    // 异常处理
    defer lock.Unlock()

    if instance == nil {
        instance = new(Tool)
    }
    return instance
}
