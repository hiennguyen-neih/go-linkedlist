# go-linkedlist
Linked list library for Go programming language (golang).

* [Go singly linked-list](./golist/)

## Install
```Go
$ go get github.com/hiennguyen-neih/go-linkedlist
```

## Document
```Go
$ go doc -all github.com/hiennguyen-neih/go-linkedlist/golist
```

## Import
```Go
import "github.com/hiennguyen-neih/go-linkedlist/golist"
```

## Usage
### Declare variables
```Go
list1 := golist.GoList[int]{}
list2 := golist.New("a", "b", "c", "d")
```
### For loop
```Go
for node := list.Head; node != nil; node = node.Next {
    // do-somethings
}
```
