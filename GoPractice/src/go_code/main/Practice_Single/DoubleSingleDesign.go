package Practice_Single

import "sync"

var lockFirst sync.Mutex

/*
第一次判断不加锁
第二次加锁保证线程安全
一旦对象建立，获取对象就不用加锁了
 */
func GetInstanceFirst() *Tool {
    if instance == nil {
        // 第一次创建对象，进行加锁操作
        lockFirst.Lock()

        // 第一次创建对象
        if instance == nil {
            instance = new(Tool)
        }

        // 第一次创建对象后就施放对象
        lockFirst.Unlock()
    }

    // 返回线程对象
    return instance
}
