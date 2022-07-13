package main

import (
    "errors"
    "fmt"
    "strconv"
    "time"
)

/* 构建音乐类型结构体 */
type Music struct {
    Id string
    Name string
    Artist string
    Source string
    Type string
}

//type MusicEntry []string

/* 实现音乐库的管理类型 */
/*
使用了一个数组切片作为基础存储结构，其他的操作其实都只是对这个数组切片的包装
 */
type MusicManager struct {
    musics []Music
}

// 创建音乐库对象
func NewMusicManager() *MusicManager {
    return &MusicManager{make([]Music, 0)}
}

// 获取音乐库长度
func (m *MusicManager) Len() int {
    return len(m.musics)
}

/* 定义音乐库类型下的函数 */
// 获取音乐库类型下的函数
func (m *MusicManager) list() int {
    return len(m.musics)
}

// 根据索引获取音乐
func (m *MusicManager) Get(index int) (music *Music, err error) {
    // 判断形参，手动返回错误内容
    if index < 0 || index >= len(m.musics) {
        return nil, errors.New("索引超过音乐库下标值!")
    }

    // 返回正常的索引结果
    return &m.musics[index], nil
}

// 根据名称查找音乐
func (m *MusicManager) Find(name string) *Music {
    // 判断音乐库是否为空，如果为空则直接返回空
    if len(m.musics) == 0 {
        return nil
    }

    // 循环取值对比
    for _, m := range m.musics {
        if m.Name == name {
            return &m
        }
    }

    return nil
}

// 添加音乐函数
func (m *MusicManager) Add(music *Music) {
    m.musics = append(m.musics, *music)
}

// 删除音乐库中音乐
func (m *MusicManager) Remove(index int) *Music {
    // 判断索引
    if index < 0 || index > len(m.musics) {
        return nil
    }

    // 根据索引创建删除对象
    removedMusic := &m.musics[index]

    // 从数组切片当中删除
    /*
    1、删除中间元素
    2、删除仅有的一个元素
    3、删除最后一个元素
     */
    if index < len(m.musics)-1 {
        // 中间元素
        m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
    }else if index == 0 {
        // 仅有一个元素
        m.musics = make([]Music, 0)
    }else {
        // 删除最后一个元素
        m.musics = m.musics[:index-1]
    }

    return removedMusic
}

/* 设计一个简单的播放接口 */
type Player interface {
    Play(source string)
}

// 设置一个MP3Player结构体
type MP3Player struct {
    stat int
    progress int // 进度
}

// 设置mp3player实现player接口下的函数
func (p *MP3Player) Play(source string) {
    fmt.Println("听MP3音乐", source)
    // 设置音乐进度
    p.progress = 0
    // 模拟正在播放
    for p.progress < 100 {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(".")
        p.progress += 10
    }

    fmt.Println("\n结束播放", source)
}

// 声明play函数
func Play(source, mtype string) {
    // 声明一个player类型的变量
    var p Player
    // 设置断言
    switch mtype {
    case "MP3":
        p = &MP3Player{}
    default:
        fmt.Println("未知类型!", mtype)
        return
    }

    // 播放音乐
    p.Play(source)
}

// 创建MusicManager结构体变量
var lib *MusicManager
var id int = 1
var ctrl, signal chan int

// 构建命令行窗口函数
func handleLibCommands(tokens []string) {
    // 对形参进行断言
    switch tokens[1] {
    case "list":
        // 循环获取音乐库的值输出
        for i := 0; i < lib.Len(); i++ {
            e, _ := lib.Get(i)
            fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
        }
    case "add":
        {
            if len(tokens) == 6 {
                id++
                lib.Add(&Music{strconv.Itoa(id),
                    tokens[2],
                tokens[3],
                tokens[4],
                tokens[5]})
            }else {
                fmt.Println("USAGE: lib add <name><artist><source><type>")
            }
        }
    case "remove":
        if len(tokens) == 3 {
            //lib.RemoveByName(tokens[2])
        }else {
            fmt.Println("USAGE: lib remove <name>")
        }
    default:
        fmt.Println("Unrecognized lib command:", tokens[1])
    }
}

// 构建命令行参数2
func handlePlayCommand(tokens []string) {
    if len(tokens) != 2 {
        fmt.Println("USAGE:play<name>")
        return
    }

    e := lib.Find(tokens[1])
    if e == nil {
        fmt.Println("音乐", tokens[1], "不存在!")
        return
    }

    Play(e.Source, e.Type)
}
