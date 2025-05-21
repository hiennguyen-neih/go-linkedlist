package main

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/golist"
)

func main() {
    var int_list1 = golist.GoList[int]{}
    var int_list2 = golist.GoList[int]{}
    var int_list3 = golist.GoList[int]{}
    var str_list1 = golist.GoList[string]{}
    var str_list2 = golist.GoList[string]{}
    var str_list3 = golist.GoList[string]{}

    fmt.Println("Example golist New")
    int_list1 = golist.GoList[int]{}
    str_list1 = golist.New("a", "b", "c")
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list1: %v\n", str_list1)

    fmt.Println("Example golist Append")
    int_list1.Append(1,2,3,4,3)
    str_list2 = golist.Append(str_list1, "d", "e", "f")
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist All and Any")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_bool := golist.All(func(i int) bool {
        return i < 4
    }, int_list1)
    str_bool := golist.Any(func(s string) bool {
        return s == "b"
    }, str_list1)
    fmt.Printf("int_list1 All < 4: %v\n", int_bool)
    fmt.Printf("str_list1 Any = b: %v\n", str_bool)

    fmt.Println("Example golist Delete")
    int_list1 = golist.New(1,2,3,4,3)
    str_list1 = golist.New("a","b","c","d","e")
    int_list1.Delete(3)
    str_list2 = golist.Delete(str_list1, "c")
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist DropLast")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_list1.DropLast()
    str_list2 = golist.DropLast(str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist DropWhile")
    int_list1 = golist.New(1,2,3,4,5,4,3,2,1)
    str_list1 = golist.New("a","b","c","d","c","b","a")
    int_list1.DropWhile(func(n int) bool {
        return n < 5
    })
    str_list2 = golist.DropWhile(func(s string) bool {
        return s != "d"
    }, str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Duplicate, Seq, Sum")
    int_list1 = golist.Duplicate(5, 0)
    int_list2 = golist.Seq(2, 9, 2)
    int_sum := golist.Sum(int_list2)
    str_list1 = golist.Duplicate(4, "X")
    str_sum := golist.Sum(str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("int_list2: %v\n", int_list2)
    fmt.Printf("int_sum: %v\n", int_sum)
    fmt.Printf("str_list1: %v\n", str_list1)
    fmt.Printf("str_sum: %v\n", str_sum)

    fmt.Println("Example golist Filter")
    int_list1 = golist.New(1,2,3,4,5,6)
    str_list1 = golist.New("a","b","a","c","a","d")
    int_list1.Filter(func(n int) bool {
        return (n % 2) == 0
    })
    str_list2 = golist.Filter(func(s string) bool {
        return s != "a"
    }, str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Foldl, Foldr and ForEach")
    int_list1 = golist.New(1,2,3,4,5)
    sum := golist.Foldl(func(n, s int) int {
        fmt.Printf("%v ", n)
        return n + s
    }, 0, int_list1)
    fmt.Printf("sum: %v\n", sum)
    fac := golist.Foldr(func(n, p int) int {
        fmt.Printf("%v ", n)
        return n * p
    }, 1, int_list1)
    fmt.Printf("fac: %v\n", fac)
    golist.ForEach(func(n int) {
        fmt.Printf("%v ", n * 2)
    }, int_list1)
    fmt.Println()

    fmt.Println("Example golist Join")
    int_list1 = golist.New(1,2,3,4)
    int_list1.Join(0)
    str_list1 = golist.New("a","b","c","d")
    str_list2 = golist.Join("X", str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Map")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_list1.Map(func(n int) int {
        return n * 2
    })
    str_list2 = golist.Map(func(s string) string {
        return strings.ToUpper(s)
    }, str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Last, Length, Max, Min, Member, Nth, Prefix, Suffix, Search")
    str_list1 = golist.New("a","b","c","d","c","b")
    str_list2 = golist.New("a","b","c")
    str_list3 = golist.New("d","c","b")
    str_last := golist.Last(str_list1)
    str_len := golist.Length(str_list1)
    str_max := golist.Max(str_list1)
    str_min := golist.Min(str_list1)
    str_mem := golist.Member("d", str_list1)
    str_nth := golist.Nth(2, str_list1)
    str_pre := golist.Prefix(str_list2, str_list1)
    str_suf := golist.Suffix(str_list3, str_list1)
    pos, val := golist.Search(func(s string) bool {
        return s == "d"
    }, str_list1)
    fmt.Printf("pos: %v\nval: %v\n", pos, val)
    fmt.Printf("str_last: %v\n", str_last)
    fmt.Printf("str_len:  %v\n", str_len)
    fmt.Printf("str_max:  %v\n", str_max)
    fmt.Printf("str_min:  %v\n", str_min)
    fmt.Printf("str_mem:  %v\n", str_mem)
    fmt.Printf("str_nth:  %v\n", str_nth)
    fmt.Printf("str_pre:  %v\n", str_pre)
    fmt.Printf("str_suf:  %v\n", str_suf)
    fmt.Printf("pos: %v - val: %v\n", pos, val)

    fmt.Println("Example golist Merge")
    int_list1 = golist.New(1,2,3,4)
    int_list2 = golist.New(5,6,7,8)
    str_list1 = golist.New("a","b","c")
    str_list2 = golist.New("d","e","f")
    str_list3 = golist.Merge(str_list1, str_list2)
    int_list1.Merge(int_list2)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list3: %v\n", str_list3)

    fmt.Println("Example golist NthTail")
    int_list1 = golist.New(1,2,3,4,5)
    str_list1 = golist.New("a","b","c","d","e")
    int_list1.NthTail(3)
    str_list2 = golist.NthTail(2, str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Reverse")
    int_list1 = golist.New(1,2,3,4,5)
    str_list1 = golist.New("a","b","c","d")
    str_list2 = golist.Reverse(str_list1)
    int_list1.Reverse()
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Split, SplitWith")
    int_list1 = golist.New(0,1,2,3,4,5,6,7,8,9)
    int_list2, int_list3 = golist.SplitWith(func(n int) bool {
        return n % 2 == 0
    }, int_list1)
    str_list1 = golist.New("a","b","c","d","e","f")
    str_list2, str_list3 = golist.Split(4, str_list1)
    fmt.Printf("int_list2: %v\n", int_list2)
    fmt.Printf("int_list3: %v\n", int_list3)
    fmt.Printf("str_list2: %v\n", str_list2)
    fmt.Printf("str_list3: %v\n", str_list3)

    fmt.Println("Example golist Sublist")
    int_list1 = golist.New(1,2,3,4,5)
    int_list1.Sublist(1, 3)
    str_list1 = golist.New("a","b","c","d","e")
    str_list2 = golist.Sublist(str_list1, 3, 5)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list2: %v\n", str_list2)

    fmt.Println("Example golist Subtract")
    int_list1 = golist.New(1,2,3,2,1,2)
    int_list2 = golist.New(2,1,2)
    int_list1.Subtract(int_list2)
    str_list1 = golist.New("a","b","c","b","a","b")
    str_list2 = golist.New("b","a","b")
    str_list3 = golist.Subtract(str_list1, str_list2)
    fmt.Printf("int_list1: %v\n", int_list1)
    fmt.Printf("str_list3: %v\n", str_list3)
}
