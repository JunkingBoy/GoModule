package parse

import (
    "fmt"
    "regexp"
)

/*
正则表达式进行匹配内容的函数
封装正则表达式为常量
*/
const regexpStr = `<div class="([^"]+)"><h5>([^"]+)</h5></div>`
func ParseContent(content []byte) {
    //测试匹配的标签为
    /*<a href="/tag/漫画" class="tag">漫画</a>-->使用正则表达式进行匹配*/
    /*<div class="sp01"><h5>登录您的帐户</h5></div>*/
    re := regexp.MustCompile(regexpStr)

    //获取正则表达式的子串--->返回一个三维数组。需要获取到每一个索引下的内容。找到所有的满足上诉正则表达式规则的标签
    /*
       这样写就可以获取到全部的标签和标枪当中的指定位置的内容的正则表达式
       由于现在是没有模拟浏览器进行访问，所以每一次的访问都需要跳转到登录的界面，而且在实际的爬取界面当中并没有那么多的url。所以通过拼串的方式捕获url先占时不做
    */
    matches := re.FindAllSubmatch(content, -1)

    //将获取到的结果作为结构体进行返回
    //result := engine.ParseResult{}

    //使用range循环获取到变量当中的值
    for _, m := range matches {
        //先打印结果进行测试
        /*这里可以测试通过拼串的方式获取到url*/
        //--->这样就获取到了该页面当中的所有的url，这个对于后期的有大量的内联框架的页面十分有效
        //fmt.Printf("url:%s\n", "域名" + string(m[2]))
        fmt.Printf("%s : %s\n", m[0], m[2])
    }
}
