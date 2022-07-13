package main

/*
实例化一个通过字符串映射函数切片的map--->输入字符串找到映射的函数
 */
var event = make(map[string][]func(interface{}))
/*
通过事件名（string）关联回调列表（[]func(interface{}）
同一个事件名称可能存在多个事件回调，因此使用回调列表保存。回调的函数声明为 func(interface{})。
 */

/*
注册事件：--->形参为
1、事件名
2、回调函数
提供给外部的通过事件名注册响应函数的入口。
 */
func RegisterEvent(name string, callBack func(interface{})) {
	//事件的列表
	list := event[name]
	/*
	通过事件名（name）进行查询，返回回调列表（[]func(interface{}）
	 */

	//在列表切片中添加函数
	list = append(list, callBack)
	/*
	为同一个事件名称在已经注册的事件回调的列表中再添加一个回调函数
	 */

	//将修改的事件列表切片保存回去
	event[name] = list
	/*
	将修改后的函数列表设置到 map 的对应事件名中
	 */
}

//调用事件
func CallEvent(name string, param interface{}) {
	//找到事件map映射
	list := event[name]
	
	//遍历列表找到函数
	for _, callback := range list {
		//传入参数调用回调
		callback(param)
	}
}