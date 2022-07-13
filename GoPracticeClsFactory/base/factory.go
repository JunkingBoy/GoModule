package base

/* 创建一个接口 */
type Class interface {
    Do()
}

/* 定义一个变量，保存接口注册好的信息 */
var (
    // 保存注册好的工厂信息
    /*
    map的key是字符串，func() Class是一个返回值为Class接口对象的函数作为map的value
    提供给工厂方注册使用，所谓的“工厂”，就是一个定义为func() Class的普通函数，调用此函数，创建一个类实例，实现的工厂内部结构体会实现Class接口
     */
    factoryByName = make(map[string]func() Class)
)

/* 注册器，类生成工厂 */
func Register(name string, factory func() Class) error {
    if "" != name {
        // 设置工厂名称
        factoryByName[name] = factory
    }
    return nil
}

/* 根据名称创建对应的类接口对象 */
func Create(name string) Class {
    if "" != name {
        /*
        在已经注册的信息中查找名字对应的工厂函数，找到后，在下一行调用并返回接口
         */
        if f, ok := factoryByName[name]; ok {
            return f()
        }else {
            panic("未找到该名称！")
        }
    }
    return nil
}
