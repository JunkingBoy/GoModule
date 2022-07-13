package main

import (
    "encoding/json"
    "fmt"
)

/*
定义屏幕尺寸结构体
 */
type Screen struct {
    size float32 //尺寸
    ResX, ResY int //X和Y轴分辨率
}

/*
定义电池容量
 */
type Battery struct {
    Capacity int //电池容量
}

/*
定义获取JsonData的函数
 */
func getJsonData() []byte {
    //定义匿名结构体
    raw := &struct {
        Screen
        Battery
        HasTouchID bool
    }{
        Screen: Screen{
            size: 8.0,
            ResX: 1920,
            ResY: 1080,
        },
        Battery: Battery{
            2910,
        },
        HasTouchID: true,
    }
    //循环获取结构体当中的内容
    jsonData, _ := json.Marshal(raw)
    return jsonData
}

func main() {
    jsonData := getJsonData()
    fmt.Printf(string(jsonData))

    //定义另一个空的匿名结构体
    screenAndTouch := struct {
        Screen
        HasTouchId bool
    }{}

    //使用json包下的方法解析
    json.Unmarshal(jsonData, &screenAndTouch)
    fmt.Printf("%+v\n", screenAndTouch)
}