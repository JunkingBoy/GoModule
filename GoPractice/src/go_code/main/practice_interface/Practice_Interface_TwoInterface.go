package main

import "fmt"

/* 定义飞行动物接口 */
type Flyer interface {
    Fly()
}

/* 定义行走动物接口 */
type Walker interface {
    Walk()
}

/* 定义鸟类型 */
type bird struct {
}

/* 定义鸟类型的方法 */
func (b *bird) Fly() {
    fmt.Println("bird behavior : fly")
}

func (b *bird) Walk()  {
    fmt.Println("bird behavior : walk")
}

/* 定义🐖类型 */
type pig struct {
}

/* 定义🐖类型的方法 */
func (p *pig) Walk() {
    fmt.Println("pig behavior : walk")
}

// 调用上述方法，使用类型断言
func main() {
    // 创建动物的名字到实例的映射
    /*
    一个 map，映射对象名字和对象实例，实例是鸟和猪
     */
    animals := map[string]interface{}{
        "bird" : new(bird),
        "pig" : new(pig),
    }

    // 遍历map
    for name, obj := range animals {
        // 使用类型断言判断值得类型
        /*
        使用类型断言获得 f，类型为 Flyer 及 isFlyer 的断言成功的判定
         */
        f, isFlyer := obj.(Flyer)
        /*
        使用类型断言获得 w，类型为 Walker 及 isWalker 的断言成功的判定
         */
        w, isWalker := obj.(Walker)

        // 打印结果
        fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

        // 判断结果
        if isFlyer {
            f.Fly()
        }

        if isWalker {
            w.Walk()
        }
    }
}
