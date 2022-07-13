package main

import (
    "GoPracticePakage/Practice_Pakage/model"
    "fmt"
)

func main() {
    // 新建model对象
    p := model.NewPerson("Lucifer")
    p.SetAge(22)
    p.SetSalary(12000)
    fmt.Println(p)
    fmt.Println(p.Name, "年龄=", p.GetAge(), "薪资=", p.GetSalary())
    model.PrintStr()
    c := model.CanFind
    fmt.Println(c)
}
