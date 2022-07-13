package main

import "fmt"

/* 定义具备刷脸特性的接口 */
type ContainCanUseFaceId interface {
    CanUseFaceId()
}

/* 定义具备偷窃特性的接口 */
type ContainStolen interface {
    Stolen()
}

/* 定义刷脸支付类型 */
type Alipay struct {
}

// 刷脸支付类型实现刷脸特性接口定义的函数
func (a *Alipay) CanUseFaceId() {
}

/* 定义现金支付类型 */
type Cash struct {
}

// 现金支付类型实现偷窃特性接口定义的函数
func (c *Cash) Stolen() {
}

/* 定义根据接口类型进行类型断言的函数 */
func print(payMethod interface{}) {
    switch payMethod.(type) {
    case ContainCanUseFaceId: // 这里判断的是接口的类型，所以case后+的是接口名称
        fmt.Println("支持刷脸支付!")
    case ContainStolen:
        fmt.Println("可能被偷窃!")
    default:
        fmt.Println("未知什么类型的接口!")
    }
}

// 调用print方法
func main() {
    print(new(Cash))
    print(new(Alipay))
    print(new(Cash))
}
