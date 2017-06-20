package main

import (
    "fmt"
    "errors"
)

func main() {
    test1()//panic error!
    test2()//defer panic
    test4()
}

//标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。
//通过判断错误对象实例来确定具体错误类型。
//如何区别使用 panic 和 error 两种方式？惯例是：导致关键流程出现不可修复性错误的使用 panic，其他使用 error。
var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
    if y == 0 { return 0, ErrDivByZero }
    return x / y, nil
}

func test4() {
    switch z, err := div(10, 0); err {
    case nil:
        println(z)
    case ErrDivByZero:
        panic(err)
    }
}


//使用延迟匿名函数或下面这样都是有效的。
func except() {
    recover()
}
func test3() {
    //defer func() {
    //    recover()
    //}()
    defer except()
    panic("test panic")
}

func test2() {
    //延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
    defer func() {
        fmt.Println(recover())
    }()

    defer func() {
        panic("defer panic")
    }()

    panic("test panic")
}

func test1() {//使用 panic 抛出错误，recover 捕获错误。
    defer func() {
        if err := recover(); err != nil {
            println(err.(string)) // 将 interface{} 转型为具体类型。
        }
    }()

    panic("panic error!")
}