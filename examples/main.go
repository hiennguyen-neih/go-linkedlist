package main

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/golist"
)

func main() {
    fmt.Println("Test golist ListOf")
    int_list := golist.ListOf[int]()
    str_list := golist.ListOf("a", "b", "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)

    fmt.Println("Test golist Append")
    int_list.Append(1,2,3,4,3)
    str_list2 := golist.Append(str_list, "d", "e", "f")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Test golist Delete")
    int_list.Delete(3)
    str_list2 = golist.Delete(str_list, "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Test golist Map")
    int_list.Map(func(n int) int {
        return n * 2
    })
    str_list2 = golist.Map(str_list, func(s string) string {
        return strings.ToUpper(s)
    })
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)
}
