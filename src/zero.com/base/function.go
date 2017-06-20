package main

import "fmt"

func main() {
    println(test001("sum: %d", 1, 2, 3))
    println(add(test002()))
    println(test003(1, 2))
    println(test004(1, 2)) // 输出: 103
}

//多个 defer 注册，按 FILO 次序执行。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
func test006(x int) {

    defer println("a")
    defer println("b")

    defer func() {
        println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
    }()

    defer println("c")
    // 输出：
    // c
    // b
    // a
    // panic: runtime error: integer divide by zero
}

func test005(x, y int) (z int) {//test005(1, 2)) // 输出: 203
    defer func() {
        println(z) // 输出: 203
    }()

    z = x + y
    return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (ret)
}


func test004(x, y int) (z int) {
    //命名返回参数允许 defer 延迟调用通过闭包读取和修改。
    defer func() {
        z += 100
    }()
    z = x + y
    return
}

func test003(x, y int) (z int) {
    z = x + y
    return
}

func test002() (int, int) {
    return 1, 2
}

func add(x, y int) int {
    return x + y
}

func test001(s string, n ...int) string {
    var x int
    for _, i := range n {
        x += i
    }
    return fmt.Sprintf(s, x)
}
