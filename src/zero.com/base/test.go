package main

import "fmt"

type User struct {
    Age  int
    Name string
}

func main() {
    var user1 = User{Age: 13, Name: "zero"}
    fmt.Printf("user1:  %+v, \t内存地址：%p\n", user1, &user1)


    //array
    arr := [2]User{user1,user1} //产生了两个副本
    fmt.Printf("arr[0]: %+v, \t内存地址：%p\n", arr[0], &arr[0])
    fmt.Printf("arr[1]: %+v, \t内存地址：%p\n", arr[1], &arr[1])

    for i, v := range arr {
        fmt.Printf("     %d: %+v, \t内存地址：%p\n", i, v, &v)
    }

}

func passV(b *User) {
    b.Age++
    b.Name = "zero007"
    fmt.Printf("传入修改后的User: %+v, 内存地址：%p, 指针的内存地址: %p\n",
        *b, b, &b)
}
