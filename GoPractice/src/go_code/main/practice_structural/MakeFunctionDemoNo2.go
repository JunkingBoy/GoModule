package main

/*
创建一个猫的基类
 */
type AnotherCat struct {
	Color string
	Name string
}

/*
创建一个黑猫的子类，内嵌基类
 */
type BlackCat struct {
	AnotherCat
}

/*
创建一个构造基类的函数
 */
func NewCat(color string) *AnotherCat {
	//返回时实例化基类
	return &AnotherCat{
		Color: color,
	}
}

/*
创建一个构造子类的函数
 */
func NewBlackCat(name string) *BlackCat {
	//返回时实例化子类并且初始化黑猫子类的成员
	Bcat := &BlackCat{}
	Bcat.Name = "Luci"
	Bcat.Color = "Red"
	//返回构造实例的子类实例
	return Bcat
}