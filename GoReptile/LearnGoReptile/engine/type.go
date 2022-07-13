package engine

/*
这是专门存放爬取到的结果的结构体的
因为需要将爬取到的结果封装起来传递到引擎当中让引擎进行工作，所以要将结果封装成为一个结构体以进行传递
*/
//Request代表任务
type Request struct {
    Url string
    //需要一个函数，这个函数用于标识解析地址的方法。并且获取到一个返回值，该返回值也是一个结构体
    ParseFunc func([]byte) ParseResult
}

/*
这是一个返回值的结构体，当我们处理完了指定的url以后将结果封装到结构体当中进行返回
*/
type ParseResult struct {
    //返回一个Request数组
    request []Request
    //返回可能会获取到的内容--->使用interface进行接收
    content []interface{}
}
