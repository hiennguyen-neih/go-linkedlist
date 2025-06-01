# go-linkedlist

![License](https://img.shields.io/github/license/hiennguyen-neih/go-linkedlist)
![Go Version](https://img.shields.io/badge/go-1.18+-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/hiennguyen-neih/go-linkedlist/golist.svg)](https://pkg.go.dev/github.com/hiennguyen-neih/go-linkedlist/golist)

Linked list library for Go programming language (golang).

* [Go singly linked-list](./golist/)

## Install

```bash
go get github.com/hiennguyen-neih/go-linkedlist
```

## Documentation

```bash
go doc -all github.com/hiennguyen-neih/go-linkedlist/golist
```

## Import

```go
import "github.com/hiennguyen-neih/go-linkedlist/golist"
```

## Usage

### Declare variables

```go
list1 := golist.GoList[int]{}
list2 := golist.New("a", "b", "c", "d", "e")
list3 := golist.FromSlice([]int{1, 2, 3, 4, 5})
```

### For loop

```go
list := golist.New("a", "b", "c", "d", "e")
for node := list.Head; node != nil; node = node.Next {
    // do somethings with node.Data
    fmt.Println(node.Data)
}
```
### Example
```go
list := golist.New(1, 2, 3, 4, 5)
list.Map(func(n int) int { return n*2 }).Join(0).Reverse()
fmt.Println(list) // [ 10 -> 0 -> 8 -> 0 -> 6 -> 0 -> 4 -> 0 -> 2 ]
```
