package main

import (
	"fmt"
	"strings"
)

func makeSure(sure string, name string) string {
	if strings.HasSuffix(name, sure) == false {
		return name + sure
	}else {
		return name
	}
}

func main() {
	func (data int) {
		fmt.Println("Hello", data)
	}(100)

	f := func(data string) {
		fmt.Println("Hello", data)
	}

	f("JunkingBoy")
}