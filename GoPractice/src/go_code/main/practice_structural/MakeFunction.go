package main

/*
创建一个猫的结构体
 */
type Cat struct {
	Color string
	Name string
}

/*
实例化猫的名字的函数
 */
func NewCatName(name string) *Cat {
	//返回值中构造结构体实例
	return &Cat{
		Name: name,
	}
}

/*
实例化猫的颜色的函数
 */
func NewCatColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}