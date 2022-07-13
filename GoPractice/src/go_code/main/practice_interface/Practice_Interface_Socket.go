package main

// 定义结构体
type Socket struct {

}

// 定义该结构体实现了的方法
func (s *Socket) Write(p []byte) (n int, err error) {
    return 0, nil
}

func (s *Socket) Close() error {
    return nil
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// 使用Writer的代码, 并不知道Socket和Closer的存在
/*
个人理解：
以java的接口为例。我们在实现接口的时候需要重写接口当中的全部的方法。这样才可以实现接口
也就是说在实现一个接口的同时我们就了解到了接口当中的全部的方法信息。这是强耦合性
在Go中定义接口的实现时候结构体本身和该结构体实现的其他方法--->注意看下面方法的入参
 */
func usingWriter(writer Writer) {
    writer.Write(nil)
}

func usingCloser(closer Closer) {
    closer.Close()
}

func main() {
    // 实例化Socket结构体
    s := new(Socket)

    // 使用usingWriter方法注意看入参
    usingWriter(s)

    // 使用usingCloser方法--->我们传入的是实现了方法里面定义的接口的对象，并不是接口对象
    usingCloser(s)
}
