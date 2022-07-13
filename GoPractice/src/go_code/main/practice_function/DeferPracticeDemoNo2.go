package main

import (
	"fmt"
	"sync"
)

var (
	//声明一个map
	valueKey = make(map[string]int)
	valueKeyGuard sync.Mutex
)

/*
声明一个函数，读取map里面的键和值，返回值，并且对竞争的资源加锁
 */
func readValue(key string) int {
	//加上资源锁
	valueKeyGuard.Lock()

	v := valueKey[key]

	//资源解锁
	valueKeyGuard.Unlock()

	return v
}

/*
使用defer简化上面的函数
 */
func readValue2(key string) int {
	//资源加锁
	valueKeyGuard.Lock()

	//使用defer让释放所最后执行
	defer valueKeyGuard.Unlock()

	//返回value
	return valueKey[key]
}

func main() {
	valueKey["name"] = 1
	valueKey["long"] = 18
	//调用readValue方法
	fmt.Println(readValue("name"))
	fmt.Println(readValue2("long"))
}