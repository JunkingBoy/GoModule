package main

import "io"

/*
声明一个类型
 */
type device struct {
}

/*
实现io.writer和write()方法
 */
func (d *device) Write(p []byte) (n int, err error) {
    return 0, nil
}

/*
实现io.closer的close()方法
 */
func (d *device) Close() error {
    return nil
}

/*
函数调用
 */
func main() {
    // 将device实例赋值给io.WriteCloser关闭器
    var wc io.WriteCloser = new(device)

    // 写入数据
    wc.Write(nil)

    // 关闭
    wc.Close()

    // 声明写入器
    var writeOnly io.Writer = new(device)

    // 写入数据
    writeOnly.Write(nil)
}
