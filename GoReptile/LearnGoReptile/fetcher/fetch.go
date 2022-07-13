package fetcher

import (
    "io/ioutil"
    "net/http"
)

/*
这个函数用户访问url并且获取到希望获取到的标签的信息
*/
func Fetch(url string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        //抛出异常
        panic(err)
    }

    //如果成功，延后执行关闭流的动作
    defer req.Body.Close()

    //模拟请求头
    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("Accept-Encoding", "gzip, deflate")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Cookie", "thinkphp_show_page_trace=0|0; thinkphp_show_page_trace=0|4; _ati=9948700605832; think_var=zh-cn; PHPSESSID=23b62311e542faaf12e72eb0d37766a9; thinkphp_show_page_trace=0|0; token=SBZruCjjaGJ3EqljjeVkqJucVONDtEiKM9Us5sYqZCtq3AeDcToNTWZvy35uMhgrpBaIv4wmPuFkkMJMBgLVBHclYunBv_CjtfIy_siOOJYLugnaRI0I0pL27kN8IIUm; users=-5PYsAYAeYYOh41Hi2idlIjgX8y1R6I1RsRb3LW-eDPstfTSSmzqgZOL43PQhRUY_K3yquAsSSIWsk40_rbAG7sL-hvHBKkUtMVeLUKBq4dxfDCSEGc0LaxwFPduJNiMofuhg3M7ErA0-v_2JU5yj0XtcqB7FZybF2qcaCLDh4m4r3HQp8g5sSrTa6LNR3P5OaNWGi355tE2WTKG3XSFJYo3uJEoMBSZGurYrKYh_sF2xGCDG_lq32teO0vg-uT8ligtNhFXW8ges-VajW72t0-o7AeXpumJHnd99RkCy0Zls151W-Jy4_5vWKuBA66Dvkbuy03LmzxIfsolvvQRABREX6Tl5YoO-E7T2_9C6BP1C5a3dX9NvndzaidPtND4sZ2IPUAVPMafparuOQUEfkTCCNSlg13cvytLqDlXhJ0jL9QVyrObRUKbUWApEZrj6kuP2a2gKc5fkOEZSZa0JA")

    //开始模拟浏览器客户端发送请求
    client := http.Client{}

    response, err := client.Do(req)
    if err != nil {
        //返回空数组
        return []byte{}, err
    }

    //完全成功的话借助ioUtil包下的方法进行读取内容
    /*
       ioutil的readall方法的返回值也是byte和err，所以可以直接返回
    */
    return ioutil.ReadAll(response.Body)
}