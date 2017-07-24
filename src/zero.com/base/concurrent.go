package main

import (
    "sync"
    "runtime"
    "time"
    "fmt"
)

func main() {
    //concurrent1()
    //concurrent2()
    //concurrent3()
    concurrent4()
}

func loop(done chan bool) {
    for i := 0; i < 10; i++ {
        fmt.Print(i)
    }
    done <- true
}
//默认地， go所有的goroutines只能在一个线程里跑
//为了达到真正的并行，runtime.GOMAXPROCS(2)
//貌似上述结论有问题
func concurrent4(){
    done := make(chan bool)
    runtime.GOMAXPROCS(2)
    go loop(done)
    go loop(done)
    fmt.Println("start")
    <-done
    <-done
}


//runtime.Gosched()
//和yield 作用类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执行

//调用 runtime.Goexit 将立即终止当前 goroutine 执行，
// 调度器确保所有已注册 defer 延迟调用被执行。
//输出：B.defer
//     A.defer
func concurrent2() {
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

func concurrent1() {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    for i := 0; i < 2; i++ {
        go func(id int) {
            defer wg.Done()
            println("Hello, World!", id)
        }(i)
    }

    wg.Wait()
}

func concurrent3() {
    a := make(chan bool)

    wg := new(sync.WaitGroup)
    wg.Add(2)

    go gofunc(wg)
    go gofunc(wg)

    go func() {
        fmt.Println("waiting")
        wg.Wait()
        fmt.Println(11)
        a <- true
    }()

    <- a
    fmt.Println("over")
}

func gofunc(wg *sync.WaitGroup) {
    time.Sleep(2 * time.Second)
    wg.Done()
}