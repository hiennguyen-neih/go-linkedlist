package main

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/golist"
    "github.com/hiennguyen-neih/go-linkedlist/golist2"
    "github.com/hiennguyen-neih/go-linkedlist/golistc"
)

func main() {
    // GoList (singly list)
    fmt.Println("Examples for GoList (singly list)")
    list := golist.New("c", "d", "e")
    fmt.Println(list)   // ["c"->"d"->"e"]
    list = golist.AppendHead(list, "a", "b")
    fmt.Println(list)   // ["a"->"b"->"c"->"d"->"e"]
    list = golist.Join(list, "x")
    fmt.Println(list)   // ["a"->"x"->"b"->"x"->"c"->"x"->"d"->"x"->"e"]
    list = golist.Map(list, func(s string) string {
        return strings.ToUpper(s)
    })
    fmt.Println(list)   // ["A"->"X"->"B"->"X"->"C"->"X"->"D"->"X"->"E"]
    fmt.Println()

    // GoList2 (doubly list)
    fmt.Println("Examples for GoList2 (doubly list)")
    list2 := golist2.New(1, 2, 3)
    fmt.Println(list2)  // [1<->2<->3]
    list2 = golist2.Append(list2, 4, 5, 6)
    fmt.Println(list2)  // [1<->2<->3<->4<->5]
    list2 = golist2.FilterMap(list2, func(n int) (bool, int) {
        return n % 2 == 0, n * 2
    })
    fmt.Println(list2)  // [4<->8<->12]
    list2 = golist2.Reverse(list2)
    fmt.Println(list2)  // [12<->8<->4]
    fmt.Println()

    // GoListC (singly circular list)
    fmt.Println("Examples for GoListC (singly circular list)")
    listc := golistc.New(1.9, 2.8, 3.7, 4.6)
    fmt.Println(listc)
}
