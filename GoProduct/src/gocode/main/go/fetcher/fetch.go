package fetcher

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

/*
这个函数用户访问url并且获取到希望获取到的标签的信息
 */
func Fetch(url string) ([]byte, error) {
    resp, err := http.Get("http://bms.spocoo.com")
    if err != nil {
        //抛出异常
        panic(err)
    }

    //如果成功，延后执行关闭流的动作
    defer resp.Body.Close()

    //判断状态码
    if resp.StatusCode != http.StatusOK {
        //打印错误的code码
        fmt.Printf("Error status code:%d", resp.StatusCode)
    }

    //完全成功的话借助ioUtil包下的方法进行读取内容
    /*
    ioutil的readall方法的返回值也是byte和err，所以可以直接返回
     */
    return ioutil.ReadAll(resp.Body)
}