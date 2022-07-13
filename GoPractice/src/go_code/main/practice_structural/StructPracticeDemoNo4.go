package main

import "fmt"

/*
写一个People结构体
两个成员：
1、name
2、指向自己
 */
type People struct {
	Name string
	Child *People
}

func main() {
	//初始化结构体的成员
	relation := &People{
		Name: "外祖父",
		Child: &People{
			Name: "父亲",
			Child: &People{
				Name: "自己",
				/*
				因为需要初始化的结构体成员可以自定义，所以可以套娃
				 */
			},
		},
	}
	fmt.Printf("%T\n", relation)
}