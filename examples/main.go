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

    // New
    fmt.Println("Example golist New")
    int_list1 = golist.GoList[int]{}
    str_list1 = golist.New("a", "b", "c")
    fmt.Printf("int_list1: %v\n", int_list1)    // [  ]
    fmt.Printf("str_list1: %v\n", str_list1)    // [ a -> b -> c ]
    fmt.Println()

    // Append - AppendHead
    fmt.Println("Example golist Append and AppendHead")
    int_list1.Append(1,2,3,4,3)
    int_list2.AppendHead(5,4,3,2,1)
    str_list2 = golist.Append(str_list1, "d", "e", "f")
    str_list3 = golist.AppendHead(str_list1, "x", "y", "z")
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 1 -> 2 -> 3 -> 4 -> 3 ]
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 5 -> 4 -> 3 -> 2 -> 1 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ a -> b -> c -> d -> e -> f ]
    fmt.Printf("str_list3: %v\n", str_list3)    // [ x -> y -> z -> a -> b -> c ]
    fmt.Println()

    // All - Any
    fmt.Println("Example golist All and Any")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_bool := golist.All(int_list1, func(i int) bool {
        return i < 4
    })
    str_bool := golist.Any(str_list1, func(s string) bool {
        return s == "b"
    })
    fmt.Printf("int_list1 All < 4: %v\n", int_bool) // false
    fmt.Printf("str_list1 Any = b: %v\n", str_bool) // true
    fmt.Println()

    // Concat - Merge - UMerge
    fmt.Println("Example golist Concat, Merge and UMerge")
    int_list1 = golist.New(2,8,6,8)
    int_list2 = golist.New(1,3,3,7)
    str_list1 = golist.New("a","b","c")
    str_list2 = golist.New("d","e","f")
    fmt.Println("Merge:")
    int_list3 = golist.Merge(int_list1, int_list2)
    fmt.Printf("int_list3: %v\n", int_list3)        // [ 1 -> 2 -> 3 -> 3 -> 6 -> 7 -> 8 -> 8 ]
    fmt.Println("UMerge:")
    int_list3 = golist.UMerge(int_list1, int_list2)
    fmt.Printf("int_list3: %v\n", int_list3)        // [ 1 -> 2 -> 3 -> 6 -> 7 -> 8 ]
    fmt.Println("Concat before int_list2 Map:")
    str_list3 = golist.Concat(str_list1, str_list2)
    int_list1.Concat(int_list2)
    fmt.Printf("int_list1: %v\n", int_list1)        // [ 2 -> 8 -> 6 -> 8 -> 1 -> 3 -> 3 -> 7 ]
    fmt.Printf("int_list2: %v\n", int_list2)        // [ 1 -> 3 -> 3 -> 7 ]
    fmt.Printf("str_list3: %v\n", str_list3)        // [ a -> b -> c -> d -> e -> f ]
    fmt.Println("Concat after int_list2 Map:")
    int_list2.Map(func(n int) int { return n * n })
    fmt.Printf("int_list1: %v\n", int_list1)        // [ 2 -> 8 -> 6 -> 8 -> 1 -> 9 -> 9 -> 49 ]
    fmt.Printf("int_list2: %v\n", int_list2)        // [ 1 -> 9 -> 9 -> 49 ]
    fmt.Println()

    // Delete
    fmt.Println("Example golist Delete")
    int_list1 = golist.New(1,2,3,4,3)
    str_list1 = golist.New("a","b","c","d","e")
    int_list1.Delete(3)
    str_list2 = golist.Delete(str_list1, "c")
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 1 -> 2 -> 4 -> 3 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ a -> b -> d -> e ]
    fmt.Println()

    // DropLast
    fmt.Println("Example golist DropLast")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_list1.DropLast()
    str_list2 = golist.DropLast(str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 1 -> 2 -> 3 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ a -> b -> c ]
    fmt.Println()

    // DropWhile - TakeWhile
    fmt.Println("Example golist DropWhile and TakeWhile")
    int_list1 = golist.New(1,2,3,4,5,4,3,2,1)
    int_list2 = golist.New(9,8,7,6,5,4,3,2,1)
    str_list1 = golist.New("a","b","c","d","c","b","a")
    int_list1.DropWhile(func(n int) bool {
        return n < 5
    })
    int_list2.TakeWhile(func(n int) bool {
        return n >= 5
    })
    str_list2 = golist.DropWhile(str_list1, func(s string) bool {
        return s != "d"
    })
    str_list3 = golist.TakeWhile(str_list1, func(s string) bool {
        return s != "d"
    })
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 5 -> 4 -> 3 -> 2 -> 1 ]
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 9 -> 8 -> 7 -> 6 -> 5 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ d -> c -> b -> a ]
    fmt.Printf("str_list3: %v\n", str_list3)    // [ a -> b -> c ]
    fmt.Println()

    // Duplicate - Seq - Sum
    fmt.Println("Example golist Duplicate, Seq and Sum")
    int_list1 = golist.Duplicate(5, 0)
    int_list2 = golist.Seq(2, 9, 2)
    int_sum := golist.Sum(int_list2)
    str_list1 = golist.Duplicate(4, "X")
    str_sum := golist.Sum(str_list1)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 0 -> 0 -> 0 -> 0 -> 0 ]
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 2 -> 4 -> 6 -> 8 ]
    fmt.Printf("int_sum: %v\n", int_sum)        // 20
    fmt.Printf("str_list1: %v\n", str_list1)    // [ X -> X -> X -> X ]
    fmt.Printf("str_sum: %v\n", str_sum)        // XXXX
    fmt.Println()

    // Filter - FilterMap
    fmt.Println("Example golist Filter and FilterMap")
    int_list1 = golist.New(1,2,3,4,5,6)
    int_list2 = golist.New(1,2,3,4,5,6)
    str_list1 = golist.New("a","b","a","c","a","d")
    int_list1.Filter(func(n int) bool {
        return (n % 2) == 0
    })
    int_list2.FilterMap(func(n int) (bool, int) {
        return (n % 2) != 0, n * 2
    })
    str_list2 = golist.Filter(str_list1, func(s string) bool {
        return s != "a"
    })
    str_list3 = golist.FilterMap(str_list1, func (s string) (bool, string) {
        return s != "a", strings.ToUpper(s)
    })
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 2 -> 4 -> 6 ]
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 2 -> 6 -> 10 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ b -> c -> d ]
    fmt.Printf("str_list3: %v\n", str_list3)    // [ B -> C -> D ]
    fmt.Println()

    // Foldl - Foldr - ForEach
    fmt.Println("Example golist Foldl, Foldr and ForEach")
    int_list1 = golist.New(1,2,3,4,5)
    sum := golist.Foldl(int_list1, 0, func(n, s int) int {
        fmt.Printf("%v ", n)                    // 1 2 3 4 5
        return n + s
    })
    fmt.Printf("| sum: %v\n", sum)                // 15
    fac := golist.Foldr(int_list1, 1, func(n, f int) int {
        fmt.Printf("%v ", n)                    // 5 4 3 2 1
        return n * f
    })
    fmt.Printf("| fac: %v\n", fac)                // 120
    golist.ForEach(int_list1, func(n int) {
        fmt.Printf("%v ", n * 2)                // 2 4 6 8 10
    })
    fmt.Println("\n")

    // Join
    fmt.Println("Example golist Join")
    int_list1 = golist.New(1,2,3,4)
    int_list1.Join(0)
    str_list1 = golist.New("a","b","c","d")
    str_list2 = golist.Join(str_list1, "X")
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 1 -> 0 -> 2 -> 0 -> 3 -> 0 -> 4 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ a -> X -> b -> X -> c -> X -> d ]
    fmt.Println()

    // Map - MapFoldl - MapFoldr
    fmt.Println("Example golist Map, MapFoldl and MapFoldr")
    int_list1 = golist.New(1,2,3,4)
    str_list1 = golist.New("a","b","c","d")
    int_list1.Map(func(n int) int {
        return n * 2
    })
    str_list2 = golist.Map(str_list1, func(s string) string {
        return strings.ToUpper(s)
    })
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 2 -> 4 -> 6 -> 8 ]
    int_list2, sum = golist.MapFoldl(int_list1, 0, func(n, s int) (int, int) {
        fmt.Printf("%v ", n)                    // 2 4 6 8
        return n * 2, s + n
    })
    fmt.Printf("| sum: %v\n", sum)              // 20
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 4 -> 8 -> 12 -> 16 ]
    int_list3, fac = golist.MapFoldr(int_list1, 1, func(n, f int) (int, int) {
        fmt.Printf("%v ", n)                    // 8 6 4 2
        return n * n, f * n
    })
    fmt.Printf("| fac: %v\n", fac)              // 384
    fmt.Printf("int_list3: %v\n", int_list3)    // [ 64 -> 36 -> 16 -> 4 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ A -> B -> C -> D ]
    fmt.Println()

    // Find - Last - Len - Max - Min - Member - Nth - Prefix - Suffix - Search
    fmt.Println("Example golist Find, Last, Len, Max, Min, Member, Nth, Prefix, Suffix and Search")
    str_list1 = golist.New("a","b","c","d","c","b")
    str_list2 = golist.New("a","b","c")
    str_list3 = golist.New("d","c","b")
    str_find := golist.Find(str_list1, "c")
    str_last := golist.Last(str_list1)
    str_len := golist.Len(str_list1)
    str_max := golist.Max(str_list1)
    str_min := golist.Min(str_list1)
    str_mem := golist.Member(str_list1, "d")
    str_nth := golist.Nth(str_list1, 2)
    str_pre := golist.Prefix(str_list2, str_list1)
    str_suf := golist.Suffix(str_list3, str_list1)
    pos, val := golist.Search(str_list1, func(s string) bool {
        return s > "c"
    })
    fmt.Printf("str_find: %v\n", str_find)      // 2
    fmt.Printf("str_last: %v\n", str_last)      // b
    fmt.Printf("str_len:  %v\n", str_len)       // 6
    fmt.Printf("str_max:  %v\n", str_max)       // d
    fmt.Printf("str_min:  %v\n", str_min)       // a
    fmt.Printf("str_mem:  %v\n", str_mem)       // true
    fmt.Printf("str_nth:  %v\n", str_nth)       // c
    fmt.Printf("str_pre:  %v\n", str_pre)       // true
    fmt.Printf("str_suf:  %v\n", str_suf)       // true
    fmt.Printf("pos: %v - val: %v\n", pos, val) // 3 - d
    fmt.Println()

    // NthTail
    fmt.Println("Example golist NthTail")
    int_list1 = golist.New(1,2,3,4,5)
    str_list1 = golist.New("a","b","c","d","e")
    int_list1.NthTail(3)
    str_list2 = golist.NthTail(str_list1, 2)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 4 -> 5 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ c -> d -> e ]
    fmt.Println()

    // Reverse
    fmt.Println("Example golist Reverse")
    int_list1 = golist.New(1,2,3,4,5)
    str_list1 = golist.New("a","b","c","d")
    str_list2 = golist.Reverse(str_list1)
    int_list1.Reverse()
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 5 -> 4 -> 3 -> 2 -> 1 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ d -> c -> b -> a ]
    fmt.Println()

    // Sort - USort
    fmt.Println("Example golist Sort and USort")
    int_list1 = golist.New(2,5,1,2,7,3,9,4,8,6,4)
    int_list2 = golist.Sort(int_list1)
    int_list3 = golist.USort(int_list1)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 2 -> 5 -> 1 -> 2 -> 7 -> 3 -> 9 -> 4 -> 8 -> 6 -> 4 ]
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 1 -> 2 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 ]
    fmt.Printf("int_list3: %v\n", int_list3)    // [ 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 ]
    fmt.Println()

    // Split - SplitWith - Partition
    fmt.Println("Example golist Split, SplitWith and Partition")
    int_list1 = golist.New(1,2,3,4,5,6,7,8,9,0)
    int_list2, int_list3 = golist.SplitWith(int_list1, func(n int) bool {
        return n < 5
    })
    str_list1 = golist.New("a","b","c","d","e","f")
    str_list2, str_list3 = golist.Split(str_list1, 4)
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 1 -> 2 -> 3 -> 4 ]
    fmt.Printf("int_list3: %v\n", int_list3)    // [ 5 -> 6 -> 7 -> 8 -> 9 -> 0 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ a -> b -> c -> d ]
    fmt.Printf("str_list3: %v\n", str_list3)    // [ e -> f ]
    int_list2, int_list3 = golist.Partition(int_list1, func(n int) bool {
        return n % 2 == 0
    })
    fmt.Printf("int_list2: %v\n", int_list2)    // [ 2 -> 4 -> 6 -> 8 -> 0 ]
    fmt.Printf("int_list3: %v\n", int_list3)    // [ 1 -> 3 -> 5 -> 7 -> 9 ]
    fmt.Println()

    // Sublist
    fmt.Println("Example golist Sublist")
    int_list1 = golist.New(1,2,3,4,5)
    int_list1.Sublist(1, 3)
    str_list1 = golist.New("a","b","c","d","e")
    str_list2 = golist.Sublist(str_list1, 3, 5)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 2 -> 3 -> 4 ]
    fmt.Printf("str_list2: %v\n", str_list2)    // [ d -> e ]
    fmt.Println()

    // Subtract
    fmt.Println("Example golist Subtract")
    int_list1 = golist.New(1,2,3,2,1,2)
    int_list2 = golist.New(2,1,2)
    int_list1.Subtract(int_list2)
    str_list1 = golist.New("a","b","c","b","a","b")
    str_list2 = golist.New("b","a","b")
    str_list3 = golist.Subtract(str_list1, str_list2)
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 3 -> 1 -> 2 ]
    fmt.Printf("str_list3: %v\n", str_list3)    // [ c -> a -> b ]
    fmt.Println()

    // Method chaining
    fmt.Println("Example golist method chaining")
    int_list1 = golist.New(1,2,3,4)
    int_list1.Join(8).Map(func(n int) int { return n * 2 }).Concat(golist.New(1,2,3)).Reverse()
    fmt.Printf("int_list1: %v\n", int_list1)    // [ 3 -> 2 -> 1 -> 8 -> 16 -> 6 -> 16 -> 4 -> 16 -> 2 ]
    fmt.Println()
}
