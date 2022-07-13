package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	/*调用包内的方法将键值对保存进去*/
	scene.Store("name", "JunkingBoy")
	scene.Store("age", 22)
	scene.Store("long", 18)
	/*
	键值对保存到 sync.Map 中，sync.Map 将键和值以 interface{} 类型进行保存。
	所以可以是任意类型的value
	 */

	/*使用包内方法从键中取值*/
	fmt.Println(scene.Load("long"))

	/*删除对应的键值对*/
	scene.Delete("age")

	//遍历键值对--->调用sync包下的方法
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate", key, value)
		return true
	})
}