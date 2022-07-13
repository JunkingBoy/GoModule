package main

import "fmt"

func main()  {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("%v, %v, %v, %q\n", i, f, b, s)
	var name = "Lucifer"
	fmt.Println(name)
	var arr = 1
	fmt.Println(arr)
	var bool = false
	fmt.Println(bool)
	var deci = 3.1415926535
	fmt.Println(deci)
	var arrs = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrs)
	array := -1
	fmt.Println(array)
	myName := "Lucifer"
	fmt.Println(myName)
	var intValueNo1, intValueNo2, intValueNo3, intValueNo4 int
	intValueNo1, intValueNo2, intValueNo3, intValueNo4 = 1, 2, 3, 4
	fmt.Println(intValueNo1, intValueNo2, intValueNo3, intValueNo4)
	var intVal, stringValue, boolVlaue = 1, "你好", true
	fmt.Println(intVal, stringValue, boolVlaue)
	intValue, stringValueNo1, boolValueNo1 := 1, "Junking", true
	fmt.Println(intValue, stringValueNo1, boolValueNo1)
}