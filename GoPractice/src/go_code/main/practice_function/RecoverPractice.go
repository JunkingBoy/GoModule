package main

import (
	"fmt"
	"runtime"
)

/*
定义一个结构体，里面的元素定义为抛出异常时需要传递的上下文信息
 */
type panicContent struct {
	function string //当前所在函数
}
/*
声明描述错误的结构体，保存执行错误的函数。
 */

/*
实现一个run函数，形参为一个函数或者是闭包
 */
func Run(entry func()) {
	//defer标记宕机处理
	defer func() { //--->非常重要：defer 将闭包延迟执行，当 panic 触发崩溃时，ProtectRun() 函数将结束运行，此时 defer 后的闭包将会发生调用
		//发生宕机时，获取panic传递的上下文并打印
		err := recover() //recover() 获取到 panic 传入的参数

		switch err.(type) { //错误类型断言
		case runtime.Error:
			//运行时异常
			fmt.Println("runtime Error:", err)
		default:
			fmt.Println("another unknow Error:", err)
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前：")

	//调用run函数，手动触发错误
	Run(func() {
		fmt.Println("手动触发宕机前:")

		//使用panic结构体传递上下文--->一个指针对象，指向结构体
		panic(&panicContent{ //--->这是将一个结构体附带信息传递给panic
			"手动触发panic",
			/*
			使用 panic 手动触发一个错误，并将一个结构体附带信息传递过去
			recover 就会获取到这个结构体信息，并打印出来
			 */
		})

		fmt.Println("手动触发宕机后。")
	})
	/*
	这不是一个运行时异常，宕机被恢复了，所以可以执行下面一个空指针错误的方法
	 */

	//手动触发空指针错误
	Run(func() {
		fmt.Println("赋值宕机前：")

		var a *int
		*a = 1
		/*
		模拟代码中空指针赋值造成的错误
		由 Runtime 层抛出错误，被 ProtectRun() 函数的 recover() 函数捕获到
		 */

		fmt.Println("赋值宕机后。")
	})
	/*
	这是一个运行时异常
	 */

	fmt.Println("运行后。")
}