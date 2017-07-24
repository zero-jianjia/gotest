package main

import "fmt"

func main() {
    //array()
    //slice()
    testMap()

    //slice1()
}

func slice1() {
    type unknow  struct {
        age int
    }

    a1 := make([]unknow, 0, 10)
    //a1 := make([]unknow, 0, 10)
    a1 = append(a1, unknow{100})
    u := a1[0]
    u.age =14
    fmt.Println(a1[0])
    //u := (a1[0])
    //fmt.Println(&u)
    //fmt.Println(ok)
    //if ok{
    //    (u).age = 14
    //}

    //a1[0].age = 14
    fmt.Println(&(a1[0]))
}

func set() {
    var null struct{}

    set := make(map[string]struct{})
    set["a"] = null
}

func testMap1() {
    type user struct{ name string }

    m := map[int]user{
        1: {"user1"},
    }
    // 当 map 因扩张而重新哈希时，各键值项存储位置都会发生改变。
    // 因此，map被设计成 not addressable。
    // 类似 m[1].name 这种期望透过原 value指针修改成员的行为自然会被禁止。

    // m[1].name = "Tom" // Error: cannot assign to m[1].name

    //正确做法是完整替换 value 或使用指针。
    u := m[1]
    u.name = "Tom"
    m[1] = u // 替换 value。

    m2 := map[int]*user{
        1: &user{"user1"},
    }
    m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。

}

func testMap() {
    a := map[string]string{
        "000":"aaa",
        "001":"bbb",
        "002":"ccc",
    }
    for k, v := range a {
        fmt.Println(k, v)
    }

    monthdays := make(map[string]int)
    monthdays["Feb"] = 29

    v, exist := monthdays["Fef"]
    fmt.Println(v, exist)
    delete(monthdays, "Feb")

    fmt.Println("---------------------------------")
    montys := make(map[string][]string)
    //montys["Fef"]=make([]string,0,100)
    //vv := montys["Fef"]
    //if len(vv)==0{
    //    montys["Fef"] = make([]string,0, 512)
    //    vv = montys["Fef"]
    //}
    montys["Fef"] = append(montys["Fef"],"af")
    montys["Fef"] = append(montys["Fef"],"af123")
    fmt.Println(montys)

}

func array() {
    var arr[5] int
    arr[0] = 41
    arr[1] = 13
    fmt.Printf("The first element is %d\n", arr[0])

    numbers := [6]int{1, 2, 3, 5}//补零
    fmt.Printf("The first element is %d\n", numbers[4])

    numbers2 := [...]int{1, 2, 3, 5}//补零
    fmt.Printf("The length is is %d\n", len(numbers2))
}

func slice() {
    a := [...]int{0, 1, 2, 3, 4 }

    s1 := a[2:4]
    fmt.Println(s1) // 2 3

    s2 := a[:4]
    fmt.Println(s2) // 0 1 2 3

    s3 := a[:]
    fmt.Println(s3) // 0 1 2 3 4

    s4 := s3[:]
    fmt.Println(s4) // 0 1 2 3 4

    s5 := append(s3, 5, 6)
    fmt.Println(s5) // 0 1 2 3 4 5 6

    var s = make([]int, 6)
    size := copy(s, s4);
    fmt.Println(size, s)//5 [0 1 2 3 4 0]
}


