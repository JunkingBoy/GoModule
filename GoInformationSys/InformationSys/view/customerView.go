package main

import (
    "GoInformationSys/InformationSys/model"
    "GoInformationSys/InformationSys/service"
    "fmt"
)

/*
定义结构体面向用户
 */
type customerView struct {
    // 接收用户输入
    key string
    // 是否循环展示主菜单--->推出时就不需要
    load bool
    // 添加字段customerService
    customerService *service.CustomerService
}

// 显示客户信息
func (this *customerView) list() {
    // 获取已经在客户列表的数据(在切片中)--->调用customerService实现的list函数
    customers := this.customerService.List()
    // 显示内容
    fmt.Println("---------------------------Customers List---------------------------")
    fmt.Println("ID\t\tNAME\t\tGENDER\t\tAGE\t\tPHONE\t\tEMAIL")
    //// 获取customer对象
    //customer := this.customerService.GetSelfInfo()
    // 循环获取拿到的客户列表信息并且打印
    for i := 0; i < len(customers); i++ {
        //// 设置num
        //customer.SetNum(i)
        // 调用customer包下的getinfo函数获取信息
        fmt.Println(customers[i].GetInfo())
    }
    fmt.Printf("\n-------------------------Customers List Finish-------------------------\n\n")
}

// 添加客户的函数
func (this *customerView) add() {
    fmt.Println("---------------------Add Customer---------------------")
    // 姓名
    name := ""
    for {
        fmt.Println("Name:")
        fmt.Scanln(&name)
        if "" != name {
            break
        }
        fmt.Println("Empty Name, Please Try Again!")
    }

    // 性别
    sex := ""
    for {
        fmt.Println("Sex:")
       fmt.Scanln(&sex)
        if "" != sex && sex == "Man" || sex == "Woman" || sex == "man" || sex == "woman" {
            break
        }
        fmt.Println("Sex Error, Please Try Again!")
    }

    // 年龄
    age := 0
    for {
        fmt.Println("Age:")
        fmt.Scanln(&age)
        if age > 0 && age < 120 {
            break
        }
        fmt.Println("Age Error, Please Try Again!")
    }

    // 电话号码
    phone := ""
    for {
        fmt.Println("Phone Number:")
        fmt.Scanln(&phone)
        if "" != phone && len(phone) == 11 {
            break
        }
        fmt.Println("Phone Number Error, Please Try Again!")
    }

    // 邮箱
    email := ""
    for {
        fmt.Println("Email:")
        fmt.Scanln(&email)
        if "" != email {
            break
        }
        fmt.Println("Empty Email, Please Try Again!")
    }

    //构建一个新的Customer实例--->调用customer包下提供的构造函数
    //注意: id号，没有让用户输入，id是唯一的，需要系统分配
    customer := model.NewCustomerByPart(name, sex, age, phone, email)

    judge := this.customerService.Add(customer)
    // 调用service包下定义的添加函数
    if judge {
        fmt.Println("---------------------Add Successfully---------------------")
    }else {
        fmt.Println("---------------------Add Fail, Can not Find The Customer By Id---------------------")
    }
}

// 获取用户输入的id并且对客户列表的id进行删除
func (this *customerView) delete() {
    fmt.Println("---------------------Delete Customer---------------------")
    fmt.Print("Select the customer number to be deleted(-1Exit):")
    id := -1
    fmt.Scanln(&id)
    if id != -1 && id != 1 {
        fmt.Println("Confirm deletion(Y/N):")
        choose := ""
        fmt.Scanln(&choose)
        switch choose {
        case "Y":
            // 调用customerService下的delete函数
            judge := this.customerService.Delete(id)
            if judge {
                fmt.Println("---------------------Delete Successfully---------------------")
            }else {
                fmt.Println("---------------------Delete Fail, Can not Find The Customer By Id---------------------")
            }
        case "y":
            // 调用customerService下的delete函数
            judge := this.customerService.Delete(id)
            if judge {
                fmt.Println("---------------------Delete Successfully---------------------")
            }else {
                fmt.Println("---------------------Delete Fail, Can not Find The Customer By Id---------------------")
            }
        case "N":
            break
        case "n":
            break
        default:
            fmt.Println("Unknow Input,Please Try Again!")
        }
    }else if id == 1 {
        fmt.Println("Can Not Delete This Data!")
    }
}

// 退出软件
func (this *customerView) exit() {
    fmt.Println("Confirm Exit(Y/N):")
    fmt.Scanln(&this.key)
    switch this.key {
    case "Y":
        break
    case "y":
        break
    case "N":
        break
    case "n":
        break
    default:
        fmt.Print("Error Input, Confirm Exit(Y/N)")
    }

    if this.key == "Y" || this.key == "y" {
        this.load = false
    }
}

// 显示主菜单
func (this *customerView) mainMenu() {
    for {
        fmt.Println("-----------------Information System-----------------")
        fmt.Println("                 1 Add Customer")
        fmt.Println("                 2 Modify Customer")
        fmt.Println("                 3 Delete Customer")
        fmt.Println("                 4 Customers List")
        fmt.Println("                 5 Exit System")
        fmt.Print("Please Input What You Want(1-5):")

        fmt.Scanln(&this.key)
        switch this.key {
        case "1":
            this.add()
        case "2":
            fmt.Println("Modify Customer Information")
        case "3":
            this.delete()
        case "4":
            this.list()
        case "5":
            this.exit()
        default:
            fmt.Println("Input Error, Please Try Again!")
            continue
        }

        // 结构体customerView的load属性不为true直接退出
        if !this.load {
            break
        }
    }
    fmt.Println("Exit System!")
}

// 函数调用
func main() {
    //在main函数中，创建一个customerView,并运行显示主菜单..
    customerView := customerView{
        key : "",
        load : true,
    }
    //这里完成对customerView结构体的customerService字段的初始化
    customerView.customerService = service.NewCustomerService()
    if customerView.load == true {
        //显示主菜单..
        customerView.mainMenu()
    }
}
