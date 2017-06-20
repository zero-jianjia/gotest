package main

import (
    "sync"
    "runtime"
)
func main() {
    concurrent1()
    concurrent2()
}

//runtime.Gosched()
//和yield 作用类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执行

//调用 runtime.Goexit 将立即终止当前 goroutine 执行，
// 调度器确保所有已注册 defer 延迟调用被执行。
//输出：B.defer
//     A.defer
func concurrent2()  {
    wg := new(sync.WaitGroup)
    wg.Add(1)

    go func() {
        defer wg.Done()
        defer println("A.defer")

        func() {
            defer println("B.defer")
            runtime.Goexit() // 终止当前 goroutine
            println("B") // 不会执行
        }()

        println("A") // 不会执行
    }()

    wg.Wait()
}

func concurrent1()  {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    for i := 0; i < 2; i++ {
        go func(id int) {
            defer wg.Done()
            println("Hello, World!",id)
        }(i)
    }

    wg.Wait()
}