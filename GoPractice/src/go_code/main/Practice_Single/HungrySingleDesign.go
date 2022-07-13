package Practice_Single

/*
新建结构体
新建单例常量
使用init函数初始化
使用全局变量初始化
 */
type config struct {
}

// 设置单例变量
var cfg *config

// 通过init函数进行初始化
func init() {
    cfg = new(config)
}

// 提供获取实例的函数
func NewConfig() *config {
    return cfg
}

// 提供全局变量获取--->所谓全局变量就是将实例声明为一个包内可使用的变量通过方法获取该变量
var cfg1 = new(config)

// 提供获取变量cfg1的函数
func NewConfigGal() *config {
    return cfg1
}
