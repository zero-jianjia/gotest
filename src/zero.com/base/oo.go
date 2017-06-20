package main

func main() {
}

//支持匿名结构，可用作结构成员或定义变量。
func oo2()  {
    type File struct {
        name string
        size int
        attr struct {
                 perm int
                 owner int
             }
    }

    f := File{
        name: "test.txt",
        size: 1025,
        // attr: {0755, 1}, // Error: missing type in composite literal
    }

    f.attr.owner = 1
    f.attr.perm = 0755

    var attr = struct {
        perm int
        owner int
    }{2, 0755}

    f.attr = attr
}

//顺序初始化必须包含全部字段，否则会出错。
func oo1()  {
    type User struct {
        name string
        age int
    }

    u1 := User{"Tom", 20}
    //u2 := User{"Tom"} // Error: too few values in struct initializer
    _ : u1
}