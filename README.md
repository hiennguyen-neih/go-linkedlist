# go-linkedlist

![License](https://img.shields.io/github/license/hiennguyen-neih/go-linkedlist)
![Go Version](https://img.shields.io/badge/go-1.18+-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/hiennguyen-neih/go-linkedlist.svg)](https://pkg.go.dev/github.com/hiennguyen-neih/go-linkedlist)

Linked list library for Go programming language (golang).

* [Go singly linked-list](./golist/)
* [Go double linked-list](./golist2/)

## Install

```bash
go get github.com/hiennguyen-neih/go-linkedlist
```

## GoList (singly linked-list)

### Import

```go
import "github.com/hiennguyen-neih/go-linkedlist/golist"
```

### Documentation

[Go Reference](https://pkg.go.dev/github.com/hiennguyen-neih/go-linkedlist/golist)

```bash
go doc -all github.com/hiennguyen-neih/go-linkedlist/golist
```

### Usage

#### Declare variables

```go
list1 := golist.GoList[int]{}
list2 := golist.New("a", "b", "c", "d", "e")
list3 := golist.FromSlice([]int{1, 2, 3, 4, 5})
```

#### For loop

```go
list := golist.New("a", "b", "c", "d", "e")
for node := list.Head; node != nil; node = node.Next {
    // do somethings with node.Data
    fmt.Println(node.Data)
}
```
#### Example
```go
list := golist.New("c", "d", "e")
list = golist.AppendHead(list, "a", "b")
list = golist.Join(list, "x")
list = golist.Map(list, func(s string) string {
    return strings.ToUpper(s)
})
fmt.Println(list) // ["A"->"X"->"B"->"X"->"C"->"X"->"D"->"X"->"E"]
```

## GoList2 (doubly linked-list)

### Import

```go
import "github.com/hiennguyen-neih/go-linkedlist/golist2"
```

### Documentation

[Go Reference](https://pkg.go.dev/github.com/hiennguyen-neih/go-linkedlist/golist2)

```bash
go doc -all github.com/hiennguyen-neih/go-linkedlist/golist2
```

### Usage

#### Declare variables

```go
list1 := golist2.GoList2[int]{}
list2 := golist2.New("a", "b", "c", "d", "e")
list3 := golist2.FromSlice([]int{1, 2, 3, 4, 5})
```

#### For loop

```go
list := golist2.New(1.0, 2.1, 3.2, 4.3, 5.4)
for node := list.Head; node != nil; node = node.Next {
    // do somethings with node.Data
    fmt.Println(node.Data)
}
```
#### Example
```go
list2 := golist2.New(1, 2, 3)
list2 = golist2.Append(list2, 4, 5, 6)
list2 = golist2.FilterMap(list2, func(n int) (bool, int) {
    return n % 2 == 0, n * 2
})
list2 = golist2.Reverse(list2)
fmt.Println(list2)  // [12<->8<->4]
```
