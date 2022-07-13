package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "reflect"
)

func main()  {
    //Fetch("https://www.baidu.com")
    //getInformation("https://www.baidu.com")
    //testGet("http://bms.spocoo.com")
    browser("http://bms.spocoo.com")
}

func Fetch(url string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        //抛出异常
        panic(err)
    }

    //模拟请求头
    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("Accept-Encoding", "gzip, deflate")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Cookie", "BIDUPSID=C85DFD56A41E07B7D1539FCD84BC148E; PSTM=1618884940; BD_UPN=12314753; __yjs_duid=1_2b4965cfd510f929fa5128f6cdfd27b81618890321484; BAIDUID=71AABA0A49D9B9B79C316CA00467EC5D:FG=1; BDUSS=9vek1ZYlR0dXRMTkZtQ0ZFQng4Vk5IR3FWN0t2SWlOTE41QURLZX5NWDJ1MWRoSUFBQUFBJCQAAAAAAAAAAAEAAABkitK6v6FraW5nQm95AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPYuMGH2LjBha2; BDUSS_BFESS=9vek1ZYlR0dXRMTkZtQ0ZFQng4Vk5IR3FWN0t2SWlOTE41QURLZX5NWDJ1MWRoSUFBQUFBJCQAAAAAAAAAAAEAAABkitK6v6FraW5nQm95AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPYuMGH2LjBha2; MCITY=-340%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BDSFRCVID=LtPOJeC62iPZqnTHHUSgh-Xz_nxCX5rTH6aoGeaEwQwcldjtwlPyEG0Psf8g0KAbl_9MogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF=tbuj_D-KfC03jtOY5-__-4_tbh_X5-RLfKJ3al7F5l8-hCLz3xTWDUkqbmcvaUueL5TaVC5-WKOxOKQphnrDMJDsyM7haxCeWjnfLf3N3KJmstP9bT3v5tjBM-cu2-biWb7M2MbdJUJP_IoG2Mn8M4bb3qOpBtQmJeTxoUJ25DnJhbLGe6KhDjv0jGtOq-rJHDbJBbFEKJF_Hn7zeTO6Qx4pbt-qJJbCLIrH_ln_5j6SjhCR0n_KLpkT-f6nBT5KaKTBWhnOX-tWDp3lWR0bejKkQN3TBnkO5bRiLRoL-R6xDn3oyTbJXp0njb3ly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCO_CI5tDD-bP365ITVhPDthMoJbto0atoyQJ0QLhQh8pcNLTDK5UuL2tvnBPTk2Rca_l7OyUjasKOIMlO1j4_e-4rhWPRwWDn7WhTgBJ-MSp5jDh3ob6ksD-Rt5tRfJGcy0hvctb3cShPm0MjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2b6QhDN_JtTL8fn4sLKDhHJnKHR6P2DTMhnOMetoCWMT-0bFHhq3kblvKOhn1bbbsqqkhBUtOhqJg2Hn7_JjOQRnRMK_R-q5bQh0tMtvabMQxtNR30Dnjtpvh8f7bhMJobUPUDUJ9LUkJLgcdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj0DKLtK_hMIDxj6_3MbjHhxnf5-Rb56uX3b7EfhRMVh7_bf--D4FrWlQXLUv7QG7MBT69Mf7_JhOyjJ5xy5K_hPJ-5xQfL6v2Wnb_bP3SV-THQT3mKP5bbN3i34jQXb5uWb3cWKOJ8UbSjh3PBTD02-nBat-OQ6npaJ5nJq5nhMJmb67JD-50exbH55uqfRke3f; BDSFRCVID_BFESS=LtPOJeC62iPZqnTHHUSgh-Xz_nxCX5rTH6aoGeaEwQwcldjtwlPyEG0Psf8g0KAbl_9MogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF_BFESS=tbuj_D-KfC03jtOY5-__-4_tbh_X5-RLfKJ3al7F5l8-hCLz3xTWDUkqbmcvaUueL5TaVC5-WKOxOKQphnrDMJDsyM7haxCeWjnfLf3N3KJmstP9bT3v5tjBM-cu2-biWb7M2MbdJUJP_IoG2Mn8M4bb3qOpBtQmJeTxoUJ25DnJhbLGe6KhDjv0jGtOq-rJHDbJBbFEKJF_Hn7zeTO6Qx4pbt-qJJbCLIrH_ln_5j6SjhCR0n_KLpkT-f6nBT5KaKTBWhnOX-tWDp3lWR0bejKkQN3TBnkO5bRiLRoL-R6xDn3oyTbJXp0njb3ly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCO_CI5tDD-bP365ITVhPDthMoJbto0atoyQJ0QLhQh8pcNLTDK5UuL2tvnBPTk2Rca_l7OyUjasKOIMlO1j4_e-4rhWPRwWDn7WhTgBJ-MSp5jDh3ob6ksD-Rt5tRfJGcy0hvctb3cShPm0MjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2b6QhDN_JtTL8fn4sLKDhHJnKHR6P2DTMhnOMetoCWMT-0bFHhq3kblvKOhn1bbbsqqkhBUtOhqJg2Hn7_JjOQRnRMK_R-q5bQh0tMtvabMQxtNR30Dnjtpvh8f7bhMJobUPUDUJ9LUkJLgcdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj0DKLtK_hMIDxj6_3MbjHhxnf5-Rb56uX3b7EfhRMVh7_bf--D4FrWlQXLUv7QG7MBT69Mf7_JhOyjJ5xy5K_hPJ-5xQfL6v2Wnb_bP3SV-THQT3mKP5bbN3i34jQXb5uWb3cWKOJ8UbSjh3PBTD02-nBat-OQ6npaJ5nJq5nhMJmb67JD-50exbH55uqfRke3f; delPer=0; BD_CK_SAM=1; PSINO=7; BAIDUID_BFESS=15889219ADC39404455A42D375683B63:FG=1; H_PS_PSSID=34837_34442_34430_34068_34741_34584_34517_26350_34867_34791_34670; H_PS_645EC=fbb6juVtfpuBzU18jx%2Bt8o1SviSgBFKlerhx8LfJ8VcTTzRopxz7rZfDs4w9%2Fh2roxjd; BD_HOME=1; BA_HECTOR=0484808k0h8ga401eu1gmvnqo0q")
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

    //开始模拟浏览器客户端发送请求
    client := http.Client{}

    response, err := client.Do(req)
    if err != nil {
        //返回空数组
        return []byte{}, err
    }

    var reader io.ReadCloser
    //判断返回的数据编码
    if response.Header.Get("Content-Encoding") == "gzip" {
       reader, err = gzip.NewReader(response.Body)
       if err != nil {
           fmt.Println(err)
           return nil, err
       }else {
           fmt.Println(response)
           reader = response.Body
       }
    }

    //如果成功，延后执行关闭流的动作
    defer response.Body.Close()

    //完全成功的话借助ioUtil包下的方法进行读取内容
    /*
       ioutil的readall方法的返回值也是byte和err，所以可以直接返回
    */
    body, err := ioutil.ReadAll(reader)
    if err != nil {
        return []byte{}, err
    }

    fmt.Println(body)

    return body, err
}

func browser(url string) error {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        //抛出异常
        panic(err)
    }

    //模拟请求头
    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("Accept-Encoding", "gzip, deflate")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Cookie", "BIDUPSID=C85DFD56A41E07B7D1539FCD84BC148E; PSTM=1618884940; BD_UPN=12314753; __yjs_duid=1_2b4965cfd510f929fa5128f6cdfd27b81618890321484; BAIDUID=71AABA0A49D9B9B79C316CA00467EC5D:FG=1; BDUSS=9vek1ZYlR0dXRMTkZtQ0ZFQng4Vk5IR3FWN0t2SWlOTE41QURLZX5NWDJ1MWRoSUFBQUFBJCQAAAAAAAAAAAEAAABkitK6v6FraW5nQm95AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPYuMGH2LjBha2; BDUSS_BFESS=9vek1ZYlR0dXRMTkZtQ0ZFQng4Vk5IR3FWN0t2SWlOTE41QURLZX5NWDJ1MWRoSUFBQUFBJCQAAAAAAAAAAAEAAABkitK6v6FraW5nQm95AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPYuMGH2LjBha2; MCITY=-340%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BDSFRCVID=LtPOJeC62iPZqnTHHUSgh-Xz_nxCX5rTH6aoGeaEwQwcldjtwlPyEG0Psf8g0KAbl_9MogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF=tbuj_D-KfC03jtOY5-__-4_tbh_X5-RLfKJ3al7F5l8-hCLz3xTWDUkqbmcvaUueL5TaVC5-WKOxOKQphnrDMJDsyM7haxCeWjnfLf3N3KJmstP9bT3v5tjBM-cu2-biWb7M2MbdJUJP_IoG2Mn8M4bb3qOpBtQmJeTxoUJ25DnJhbLGe6KhDjv0jGtOq-rJHDbJBbFEKJF_Hn7zeTO6Qx4pbt-qJJbCLIrH_ln_5j6SjhCR0n_KLpkT-f6nBT5KaKTBWhnOX-tWDp3lWR0bejKkQN3TBnkO5bRiLRoL-R6xDn3oyTbJXp0njb3ly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCO_CI5tDD-bP365ITVhPDthMoJbto0atoyQJ0QLhQh8pcNLTDK5UuL2tvnBPTk2Rca_l7OyUjasKOIMlO1j4_e-4rhWPRwWDn7WhTgBJ-MSp5jDh3ob6ksD-Rt5tRfJGcy0hvctb3cShPm0MjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2b6QhDN_JtTL8fn4sLKDhHJnKHR6P2DTMhnOMetoCWMT-0bFHhq3kblvKOhn1bbbsqqkhBUtOhqJg2Hn7_JjOQRnRMK_R-q5bQh0tMtvabMQxtNR30Dnjtpvh8f7bhMJobUPUDUJ9LUkJLgcdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj0DKLtK_hMIDxj6_3MbjHhxnf5-Rb56uX3b7EfhRMVh7_bf--D4FrWlQXLUv7QG7MBT69Mf7_JhOyjJ5xy5K_hPJ-5xQfL6v2Wnb_bP3SV-THQT3mKP5bbN3i34jQXb5uWb3cWKOJ8UbSjh3PBTD02-nBat-OQ6npaJ5nJq5nhMJmb67JD-50exbH55uqfRke3f; BDSFRCVID_BFESS=LtPOJeC62iPZqnTHHUSgh-Xz_nxCX5rTH6aoGeaEwQwcldjtwlPyEG0Psf8g0KAbl_9MogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF_BFESS=tbuj_D-KfC03jtOY5-__-4_tbh_X5-RLfKJ3al7F5l8-hCLz3xTWDUkqbmcvaUueL5TaVC5-WKOxOKQphnrDMJDsyM7haxCeWjnfLf3N3KJmstP9bT3v5tjBM-cu2-biWb7M2MbdJUJP_IoG2Mn8M4bb3qOpBtQmJeTxoUJ25DnJhbLGe6KhDjv0jGtOq-rJHDbJBbFEKJF_Hn7zeTO6Qx4pbt-qJJbCLIrH_ln_5j6SjhCR0n_KLpkT-f6nBT5KaKTBWhnOX-tWDp3lWR0bejKkQN3TBnkO5bRiLRoL-R6xDn3oyTbJXp0njb3ly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCO_CI5tDD-bP365ITVhPDthMoJbto0atoyQJ0QLhQh8pcNLTDK5UuL2tvnBPTk2Rca_l7OyUjasKOIMlO1j4_e-4rhWPRwWDn7WhTgBJ-MSp5jDh3ob6ksD-Rt5tRfJGcy0hvctb3cShPm0MjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2b6QhDN_JtTL8fn4sLKDhHJnKHR6P2DTMhnOMetoCWMT-0bFHhq3kblvKOhn1bbbsqqkhBUtOhqJg2Hn7_JjOQRnRMK_R-q5bQh0tMtvabMQxtNR30Dnjtpvh8f7bhMJobUPUDUJ9LUkJLgcdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj0DKLtK_hMIDxj6_3MbjHhxnf5-Rb56uX3b7EfhRMVh7_bf--D4FrWlQXLUv7QG7MBT69Mf7_JhOyjJ5xy5K_hPJ-5xQfL6v2Wnb_bP3SV-THQT3mKP5bbN3i34jQXb5uWb3cWKOJ8UbSjh3PBTD02-nBat-OQ6npaJ5nJq5nhMJmb67JD-50exbH55uqfRke3f; delPer=0; BD_CK_SAM=1; PSINO=7; BAIDUID_BFESS=15889219ADC39404455A42D375683B63:FG=1; H_PS_PSSID=34837_34442_34430_34068_34741_34584_34517_26350_34867_34791_34670; H_PS_645EC=fbb6juVtfpuBzU18jx%2Bt8o1SviSgBFKlerhx8LfJ8VcTTzRopxz7rZfDs4w9%2Fh2roxjd; BD_HOME=1; BA_HECTOR=0484808k0h8ga401eu1gmvnqo0q")
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

    //开始模拟浏览器客户端发送请求
    client := http.Client{}

    response, err := client.Do(req)
    if err != nil {
        //打印错误信息
        log.Println(err)
    }

    var reader io.ReadCloser
    //判断返回的数据编码
    if response.Header.Get("Content-Encoding") == "gzip" {
        reader, err = gzip.NewReader(response.Body)
        if err != nil {
            log.Println(err)
        }else {
            fmt.Println(response)
            reader = response.Body
        }
    }

    //使用buf流读取resp.body信息
    buf := bytes.NewBuffer(make([]byte, 0, 512))
    length, _ := buf.ReadFrom(reader)
    fmt.Println(len(buf.Bytes()))
    fmt.Println(length)
    fmt.Println(string(buf.Bytes()))

    return nil
}

func getInformation(url string) error {
    req, err := http.Get(url)
    //失败，抛出异常
    if err != nil {
        //抛出异常
        panic(err)
    }
    //成功，延后执行关闭流的动作
    defer req.Body.Close()

    fmt.Println(req)

    return nil
}

func testGet(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        log.Println(err)
        panic(err)
    }

    //请求正常，延迟关闭响应流
    defer resp.Body.Close()

    //循环打印响应头
    headers := resp.Header
    //循环获取响应头报文信息
    for k, v := range headers {
        fmt.Printf("%v, %v\n", k, v)
    }

    //打印响应头信息
    fmt.Printf("是否压缩：%t\n", resp.Uncompressed)
    fmt.Println(reflect.TypeOf(resp.Body)) //*http.gzipReader
    fmt.Println(resp.Close)

    //使用buf流读取resp.body信息
    buf := bytes.NewBuffer(make([]byte, 0, 512))
    length, _ := buf.ReadFrom(resp.Body)
    fmt.Println(len(buf.Bytes()))
    fmt.Println(length)
    fmt.Println(string(buf.Bytes()))

    return nil
}