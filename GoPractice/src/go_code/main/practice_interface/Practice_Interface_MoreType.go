package main

// 定义一个接口，该接口有一个服务需要的开启和输出日志的方法
type Service interface {
    // 开启服务
    Start()

    // 输出日志
    Log(string2 string)
}

// 定义一个日志类型
type Logger struct {
}

// 类型实现接口中的函数
func (g *Logger) Log(l string) {
}

// 定义另一个类型，该类型的属性是实现了接口当中方法的类型
type GameService struct {
    Logger // 嵌入日志类型
}

// 再用外类型去实现接口当中的另一个函数
func (g *GameService) Start() {
}

// 方法调用
func main() {
    // 创建类型对象
    var s Service = new(GameService)

    // 调用接口方法
    s.Start()
    s.Log("HelloWorld!")
}
