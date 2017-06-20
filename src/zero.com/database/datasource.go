package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func main() {
    db, err := sql.Open("mysql", "dsp_user:123456@tcp(10.210.228.92:3307)/jianjia?charset=utf8")
    checkErr(err)

    rows, err := db.Query("SELECT id FROM adunit limit 5")

    //for rows.Next() {
    //    var id string
    //
    //    rows.Columns()
    //    err = rows.Scan(&id)
    //    checkErr(err)
    //    fmt.Println(id)
    //}

    column, _ := rows.Columns()  //读出查询出的列字段名
    values := make([][]byte, len(column))   //values是每个列的值，这里获取到byte里
    scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
    for i := range values {    //让每一行数据都填充到[][]byte里面
        scans[i] = &values[i]
    }
    results := make(map[int]map[string]string) //最后得到的map
    i := 0
    for rows.Next() { //循环，让游标往下移动
        if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
            fmt.Println(err)
            return
        }
        row := make(map[string]string) //每行数据
        for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
            key := column[k]
            row[key] = string(v)
        }
        results[i] = row //装入结果集中
        i++
    }
    for k, v := range results { //查询出来的数组
        fmt.Println(k, v)
    }

    db.Close()

}

func checkErr(errMasg error) {
    if errMasg != nil {
        panic(errMasg)
    }
}
