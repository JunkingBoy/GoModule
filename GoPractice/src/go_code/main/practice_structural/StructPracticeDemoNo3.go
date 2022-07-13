package main

/*
创建一个命令行结构体
 */
type Command struct {
	Name string //指令名称
	Var *int //指令绑定的变量
	Comment string //注释
}

//定义指令绑定的变量
var version int = 1

/*
使用函数封装实例化结构体、对成员进行赋值、绑定等操作
调用这个函数自动的对结构体进行赋值操作并且返回传入的参数
 */
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name: name,
		Var: varref,
		Comment: comment,
	}
}

func main() {
	//使用取地址对结构体进行实例化
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
}