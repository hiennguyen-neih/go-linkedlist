package main

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/golist"
)

func main() {
    fmt.Println("Example golist ListOf")
    int_list := golist.GoList[int]{}
    str_list := golist.ListOf("a", "b", "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)

    fmt.Println("Example golist Append")
    int_list.Append(1,2,3,4,3)
    str_list2 := golist.Append(str_list, "d", "e", "f")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist All and Any")
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

    fmt.Println("Example golist Delete")
    int_list = golist.ListOf(1,2,3,4,3)
    str_list = golist.ListOf("a","b","c","d","e")
    int_list.Delete(3)
    str_list2 = golist.Delete(str_list, "c")
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist DropLast")
    int_list = golist.ListOf(1,2,3,4)
    str_list = golist.ListOf("a","b","c","d")
    int_list.DropLast()
    str_list2 = golist.DropLast(str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist DropWhile")
    int_list = golist.ListOf(1,2,3,4,5,4,3,2,1)
    str_list = golist.ListOf("a","b","c","d","c","b","a")
    int_list.DropWhile(func(n int) bool {
        return n < 5
    })
    str_list2 = golist.DropWhile(func(s string) bool {
        return s != "d"
    }, str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Duplicate")
    int_list = golist.Duplicate(5, 0)
    str_list = golist.Duplicate(4, "X")
    fmt.Printf("int_list: %v\n", int_list)
    fmt.Printf("str_list: %v\n", str_list)

    fmt.Println("Example golist Filter")
    int_list = golist.ListOf(1,2,3,4,5,6)
    str_list = golist.ListOf("a","b","a","c","a","d")
    int_list.Filter(func(n int) bool {
        return (n % 2) == 0
    })
    str_list2 = golist.Filter(func(s string) bool {
        return s != "a"
    }, str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Foldl, Foldr and ForEach")
    int_list = golist.ListOf(1,2,3,4,5)
    sum := golist.Foldl(func(n, s int) int {
        fmt.Printf("%v ", n)
        return n + s
    }, 0, int_list)
    fmt.Printf("sum: %v\n", sum)
    fac := golist.Foldr(func(n, p int) int {
        fmt.Printf("%v ", n)
        return n * p
    }, 1, int_list)
    fmt.Printf("fac: %v\n", fac)
    golist.ForEach(func(n int) {
        fmt.Printf("%v ", n * 2)
    }, int_list)
    fmt.Println()

    fmt.Println("Example golist Join")
    int_list = golist.ListOf(1,2,3,4)
    int_list.Join(0)
    str_list = golist.ListOf("a","b","c","d")
    str_list2 = golist.Join("X", str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Map")
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

    fmt.Println("Example golist Last, Max, Min, Member, Nth")
    str_list = golist.ListOf("a","b","c","d","c","b")
    str_last := golist.Last(str_list)
    str_max := golist.Max(str_list)
    str_min := golist.Min(str_list)
    str_mem := golist.Member("d", str_list)
    str_nth := golist.Nth(2, str_list)
    fmt.Printf("str_last: %v\n", str_last)
    fmt.Printf("str_max:  %v\n", str_max)
    fmt.Printf("str_min:  %v\n", str_min)
    fmt.Printf("str_mem:  %v\n", str_mem)
    fmt.Printf("str_nth:  %v\n", str_nth)

    fmt.Println("Example golist Merge")
    int_list = golist.ListOf(1,2,3,4)
    int_list2 := golist.ListOf(5,6,7,8)
    str_list = golist.ListOf("a","b","c")
    str_list2 = golist.ListOf("d","e","f")
    str_list3 := golist.Merge(str_list, str_list2)
    int_list.Merge(int_list2)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)
    fmt.Printf("str_list3: %v\n", str_list3)

    fmt.Println("Example golist NthTail")
    int_list = golist.ListOf(1,2,3,4,5)
    str_list = golist.ListOf("a","b","c","d","e")
    int_list.NthTail(3)
    str_list2 = golist.NthTail(2, str_list)
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Reverse")
    int_list = golist.ListOf(1,2,3,4,5)
    str_list = golist.ListOf("a","b","c","d")
    str_list2 = golist.Reverse(str_list)
    int_list.Reverse()
    fmt.Printf("int_list:  %v\n", int_list)
    fmt.Printf("str_list:  %v\n", str_list)
    fmt.Printf("str_list2: %v\n", str_list2)
}
