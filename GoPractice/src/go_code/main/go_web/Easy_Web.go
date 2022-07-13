package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

/* 资源访问函数 */
func index(w http.ResponseWriter, r *http.Request) {
    // 设置访问路由--->读取内容
    content, _ := ioutil.ReadFile("./Practice.html")
    // 输出内容
    w.Write(content)
}
