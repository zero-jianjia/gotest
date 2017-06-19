package main

import "fmt"

func main() {
    //array()
    //slice()
    testMap()
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

    v, exist := monthdays["Feb"]
    fmt.Println(v, exist)
    delete(monthdays, "Feb")
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


