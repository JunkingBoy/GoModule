package main

import (
    "errors"
    "fmt"
    "os"
)

/*
声明日志写入接口,后面的类型都是通过该接口的方法实现日志写入
*/
type LogWriter interface {
    Write(data interface{}) error
}

/*
定义一个命令行写入器类型
要通过日志写入接口的方法进行日志写入
 */
type consoleWriter struct {
}

// 实现日志write()方法
func (f *consoleWriter) Write(data interface{}) error {
    // 序列化字符串
    str := fmt.Sprintf("%v\n", data)

    // 将数据以字节数组的形式写入命令行
    _, err := os.Stdout.Write([]byte(str))

    return err
}

// 创建命令行写入器实例函数
func newConsoleWriter() *consoleWriter {
    return &consoleWriter{}
}

/*
定义一个文件写入器类型
 */
type fileWriter struct {
    // 定义属性为文件io类型
    file *os.File
}

// 文件写入器写入文件名
func (f *fileWriter) SetFile(fileName string) (err error) {
    // 文件已打开则关闭前一个文件
    if nil != f.file {
        f.file.Close()
    }

    // 创建文件并且保存句柄
    f.file, err = os.Create(fileName)

    // 创建过程出错则返回错误信息
    return err
}

// 文件写入器实现write函数
func (f *fileWriter) Write(data interface{}) error {
    // 日志为空判断
    if nil == f.file {
        // 输出对象
        return errors.New("文件不存在!")
    }

    // 序列化数据
    str := fmt.Sprintf("%v\n", data)

    // 将数据以字节数组的形式写入文件--->调用f实现的函数
    _, err := f.file.Write([]byte(str))

    return err
}

// 提供文件写入器实例化函数
func newFileWriter() *fileWriter {
    return &fileWriter{}
}

/*
定义日志器类型
 */
type LoggerTool struct {
    // 日志写入函数
    writerList []LogWriter
}

// 定义该日志写入器的方法
// 1.注册日志写入器--->将实现了LogWriter接口的对象接收并加入日志器中
func (l *LoggerTool) RegisterWriter(Writer LogWriter) {
    // 调用接口方法，添加日志
    l.writerList = append(l.writerList, Writer)
}

// 2.将一个任意类型的数据写入日志
func (l *LoggerTool) Log(data interface{}) {
    // 遍历所有注册的写入器
    for _, writer := range l.writerList {
        // 日志输出到每一个写入器
        writer.Write(data)
    }
}

// 提供实例化日志器的函数
func NewLogger() *LoggerTool {
    return &LoggerTool{}
}

// 提供创建日志器的方法
func createLoggerTool() *LoggerTool {
    // 创建日志器--->调用实例化方法创建对象
    l := NewLogger()

    // 创建命令行写入器
    cw := newConsoleWriter()

    // 注册命令行写入器到日志器中
    l.RegisterWriter(cw)

    // 创建文件写入器
    fw := newFileWriter()

    // 设置文件名
    if err := fw.SetFile("学习地址路径大全.log"); err != nil {
        fmt.Println(err)
    }

    // 注册文件写入器到日志器中
    l.RegisterWriter(fw)

    return l
}

func main() {
    // 准备日志器
    l := createLoggerTool()

    // 写一个日志
    l.Log("HelloWorld!")
}
