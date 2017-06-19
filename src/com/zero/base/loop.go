package main

import "fmt"

func main() {
    loopTest()
}

func loopTest() {
    var b int = 5
    var a int

    /* for 循环 */
    //这里的a是局部变量，不是上面的a
    for a := 0; a < 3; a++ {
        fmt.Printf("a 的值为: %d\n", a)
    }

    for a < b {
        a++
        fmt.Printf("a 的值为: %d\n", a)
    }

    numbers := [6]int{1, 2, 3, 5}
    for i, x := range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
    }

    //for true {
    //    fmt.Printf("这是无限循环。\n")
    //}
}
