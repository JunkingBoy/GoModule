package main

import (
    "bufio"
    "bytes"
    "fmt"
)

func main() {
    // 声明字节数组
    data := []byte("JunkingBoy")

    // 新建newreader对象
    re := bytes.NewReader(data)

    // 构建readerwriter对象
    r := bufio.NewReader(re)

    //由于read方法返回的是字节切片长度，所以新建字节切片对象
    var buf [128]byte

    // 调用reader对象的read方法
    n, err := r.Read(buf[:])

    fmt.Println(string(buf[:n]), n, err)
}
