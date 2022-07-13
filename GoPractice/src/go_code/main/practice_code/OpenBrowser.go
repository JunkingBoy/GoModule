package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseCmd(cmd string) {
	cmd = strings.TrimSpace(cmd)
	switch {
	case strings.HasPrefix(cmd, ""):

	case strings.HasPrefix(cmd, ""):

	case cmd == "help":
		help()
	}
}

func help() {
	fmt.Println(
		`open:
					open windows [-url url] 使用windows默认浏览器打开网址，如果不输入url默认打开我的博客地址
					open linux [-url url] 使用linux浏览器打开网址，如果不输入url默认打开我的博客地址`,
		)
}

func nextLine() string {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	return string(bytes)
}

func main() {
	cmd := nextLine()
	parseCmd(cmd)
}