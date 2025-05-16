package main

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/golist"
)

func main() {
    fmt.Println("Test golist ListOf")
    int_list := golist.GoList[int]{}
    str_list := golist.ListOf("a", "b", "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)

    fmt.Println("Test golist Append")
    int_list.Append(1,2,3,4,3)
    str_list2 := golist.Append(str_list, "d", "e", "f")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Test golist All and Any")
    int_list = golist.ListOf(1,2,3,4)
    str_list = golist.ListOf("a","b","c","d")
    int_bool := golist.All(func(i int) bool {
        return i < 4
    }, int_list)
    str_bool := golist.Any(func(s string) bool {
        return s == "b"
    }, str_list)
    fmt.Printf("int_list All < 4: %v\n", int_bool)
    fmt.Printf("str_list Any = b: %v\n", str_bool)

    fmt.Println("Test golist Delete")
    int_list = golist.ListOf(1,2,3,4,3)
    str_list = golist.ListOf("a","b","c","d","e")
    int_list.Delete(3)
    str_list2 = golist.Delete(str_list, "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Test golist DropLast")
    int_list = golist.ListOf(1,2,3,4)
    str_list = golist.ListOf("a","b","c","d")
    int_list.DropLast()
    str_list2 = golist.DropLast(str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Test golist Map")
    int_list = golist.ListOf(1,2,3,4)
    str_list = golist.ListOf("a","b","c","d")
    int_list.Map(func(n int) int {
        return n * 2
    })
    str_list2 = golist.Map(func(s string) string {
        return strings.ToUpper(s)
    }, str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)
}
