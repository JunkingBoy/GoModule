package main

import (
    "github.com/skip2/go-qrcode"
    "image/color"
    "log"
)

/*
自定义一个二维码以及二维码内容
 */
func main() {
    qr, err := qrcode.New("https://www.cnblogs.com/JunkingBoy/", qrcode.Medium)
    if err != nil {
        log.Fatal(err)
    }else {
        qr.BackgroundColor = color.Black
        qr.ForegroundColor = color.White
        qr.WriteFile(256,"./学习资料.png")
    }
}
