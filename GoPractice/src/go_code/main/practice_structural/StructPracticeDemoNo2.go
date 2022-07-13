package main

/*
创建一个player结构体
 */
type Player struct {
	Name string
	HP float64
	MP float64
}

func main() {
	//使用new关键字对结构体进行实例化
	playerNo1 := new(Player)
	playerNo1.Name = "JunkingBoy"
	playerNo1.HP = 100000000
	playerNo1.MP = 100000000
}