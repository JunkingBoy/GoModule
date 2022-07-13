package main

func main() {
	/*声明一个Map*/
	mapInt := make(map [int] int)

	/*开启一段并发代码，不断地写入value*/
	go func() {
		for {
			mapInt[1] = 1
		}
	}()

	/*开启另一段并发代码，读取mapInt*/
	go func() {
		for {
			_ =mapInt[1]
		}
	}()

	/*输出*/
	for {

	}
}