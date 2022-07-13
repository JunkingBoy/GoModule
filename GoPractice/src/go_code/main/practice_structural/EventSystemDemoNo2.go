package main

import (
	"fmt"
)

/*
添加一个结构体
 */
type Actor struct {

}

/*
为这个结构体添加一个事件处理函数
相当于类的内部方法 
*/
func (print *Actor) onEvent(param interface{}) {
	//打印内容
	fmt.Println("actor event is:", param)
}

/*
定义一个全局事件
为全局事件响应函数。有时需要全局进行侦听或者处理一些事件，这里使用普通函数实现全局事件的处理。
 */
func GlobalEvent(param interface{}) {
	fmt.Println("global event is:", param)
}

func main() {
	//实例化结构体
	ac := new(Actor)

	//注册一个回调函数
	RegisterEvent("OnSkill", ac.onEvent)
	/*
	注册一个 OnSkill 事件，实现代码由 ac 的 OnEvent 进行处理。也就是 Actor的OnEvent() 方法。
	 */

	//在注册的回调函数上在注册一个全局事件
	RegisterEvent("Onskill", GlobalEvent)
	/*
	模拟同一个名称对应不同的回调函数的场景
	注册一个 OnSkill 事件，实现代码由 GlobalEvent 进行处理
	注册的是同一个名字的事件，但前面注册的事件不会被覆盖，而是被添加到事件系统中，关联 OnSkill 事件的函数列表中。
	 */

	//调用事件
	CallEvent("OnSkill", 100)
	/*
	模拟处理事件，通过 CallEvent() 函数传入两个参数，第一个为事件名，第二个为处理函数的参数
	 */
}