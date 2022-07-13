package main

import (
    "fmt"
    "math"
)

/* 创建一个算术表达式接口 */
type Expr interface {
    // 每个表达式都必须都需要返回，所以在接口当中定义该方法
    Eval(env Env)float64
}

/* 声明可能的数据类型 */
// Var表示一个变量。如:x
type Var string

// literal表示数字常量。如:3.14
type literal float64

// unary表示一元操作符表达式。如:-x
type unary struct {
    op rune // "+","-"中的一个
    x Expr
}

// binary表示二元操作符表达式。如:x+y
type binary struct {
    op rune // "+","-","*","/"中的一个
    x, y Expr
}

// call表示函数调用表达式
type call struct {
    fn string // pow\sin\sqrt函数中的一个
    args []Expr
}

// 添加上下文把变量映射到数值
type Env map[Var]float64

// Var和Eval类型的数据实现接口当中的方法
func (v Var) Eval(env Env) float64 {
    return env[v]
}

func (l literal) Eval(env Env) float64 {
    return float64(l)
}

// `unary`和`binary`的`Eval`方法首先对它们的操作数递归求值，然后应用`op`操作。不把除以 0 或者无穷大当做错误
func (u unary) Eval(env Env) float64 {
    // switch进行类型断言
    switch u.op {
    case '+':
        return +u.x.Eval(env)
    case '-':
        return -u.x.Eval(env)
    }

    // 抛出异常
    panic(fmt.Sprintf("不支持的操作:%q", u.op))
}

func (b binary) Eval(env Env) float64 {
    // 进行类型断言
    switch b.op {
    case '+':
        return b.x.Eval(env) + b.y.Eval(env)
    case '-':
        return b.x.Eval(env) - b.y.Eval(env)
    case '*':
        return b.x.Eval(env) * b.y.Eval(env)
    case '/':
        return b.x.Eval(env) / b.y.Eval(env)
    }

    // 抛出异常
    panic(fmt.Sprintf("不支持的流操作:%q", b.op))
}

// `call`方法先对`pow`、`sin`或者`sqrt`函数的参数求值，再调用`math`包中的对应函数
func (c call) Eval(env Env) float64 {
    switch c.fn {
    case "pow":
        return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
    case "sin":
        return math.Sin(c.args[0].Eval(env))
    case "sqrt":
        return math.Sqrt(c.args[0].Eval(env))
    }

    // 抛出异常
    panic(fmt.Sprintf("不支持的呼叫方法:%s", c.fn))
}
