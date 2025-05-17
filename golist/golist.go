// Package golist contains functions and methods for singly linked list in Go.
package golist

import (
    "fmt"
    "strings"
)

/*
 * Define structs and interfaces
 */

// Node in singly linked List.
type node[T comparable] struct {
    data T
    next *node[T]
}

// Struct of Go singly linked list.
type GoList[T comparable] struct {
    Head *node[T]
}

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

/*
 * Exported functions
 */

// Create new singly linked list.
func ListOf[T comparable](values ...T) (list GoList[T]) {
    list.Append(values...)
	return
}

// Return true if fun returns true for all elements in list, otherwise return false.
func All[T comparable](fun func(T) bool, list GoList[T]) bool {
    curr := list.Head
    for curr != nil {
        if fun(curr.data) == false {
            return false
        }
        curr = curr.next
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise return false.
func Any[T comparable](fun func(T) bool, list GoList[T]) bool {
    curr := list.Head
    for curr != nil {
        if fun(curr.data) == true {
            return true
        }
        curr = curr.next
    }
    return false
}

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) (result GoList[T]) {    
    curr := list.Head
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    result.Append(values...)
    return
}

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    curr := list.Head
    for curr != nil {
        if curr.data != value {
            result.Append(curr.data)
            curr = curr.next
        } else {
            curr = curr.next
            break
        }
    }
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    return
}

// Drop the last element in list and return as new list.
func DropLast[T comparable](list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for curr != nil {
        if curr.next.next == nil {
            result.Append(curr.data)
            return
        }
        result.Append(curr.data)
        curr = curr.next
    }
    return
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for curr != nil {
        if fun(curr.data) == false {
            break
        }
        curr = curr.next
    }
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    return
}

// Returns a list containing n copies of term element.
func Duplicate[T comparable](n int, elem T) (result GoList[T]) {
    for i := 0; i < n; i++ {
        result.Append(elem)
    }
    return
}

// Return new list contains elements in input that fun returns true.
func Filter[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    curr := list.Head
    for curr != nil {
        if fun(curr.data) == true {
            result.Append(curr.data)
        }
        curr = curr.next
    }
    return
}

// Execute fun with input is elements in list from left to right and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldl[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    curr := list.Head
    for curr != nil {
        acc0 = fun(curr.data, acc0)
        curr = curr.next
    }
    return acc0
}

// Execute fun with input is elements in list from right to left and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldr[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    reverse := Reverse(list)
    curr := reverse.Head
    for curr != nil {
        acc0 = fun(curr.data, acc0)
        curr = curr.next
    }
    return acc0
}

// Applying function for each element in list.
func ForEach[T comparable](fun func(T), list GoList[T]) {
    curr := list.Head
    for curr != nil {
        fun(curr.data)
        curr = curr.next
    }
}

// Insert sep between each element in list and return as new list.
func Join[T comparable](sep T, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for curr != nil {
        result.Append(curr.data)
        if curr.next != nil {
            result.Append(sep)
        }
        curr = curr.next
    }
    return
}

// Return the last element in list.
func Last[T comparable](list GoList[T]) (last T) {
    // last = nil
    curr := list.Head
    for curr != nil {
        if curr.next == nil {
            last = curr.data
            return
        } else {
            curr = curr.next
        }
    }
    return
}

// Applying function to every elements in list and return as new list.
func Map[T comparable](fun func(T) T, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for curr != nil {
        result.Append(fun(curr.data))
        curr = curr.next
    }
    return
}

// Return maximum element in list.
// This function only works with list of numbers or strings.
func Max[T Ordered](list GoList[T]) (max T) {
    curr := list.Head
    max = curr.data
    for curr != nil {
        if curr.data > max {
            max = curr.data
        }
        curr = curr.next
    }
    return
}

// Return true if elem in list, otherwise return false.
func Member[T comparable](elem T, list GoList[T]) bool {
    curr := list.Head
    for curr != nil {
        if curr.data == elem {
            return true
        }
        curr = curr.next
    }
    return false
}

// Return a list that is merged of list1 and list2
func Merge[T comparable](list1 GoList[T], list2 GoList[T]) (result GoList[T]) {
    curr := list1.Head
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    curr = list2.Head
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    return
}

// Return minimum element in list.
// This functions only works with list of numbers or strings.
func Min[T Ordered](list GoList[T]) (min T) {
    curr := list.Head
    min = curr.data
    for curr != nil {
        if curr.data < min {
            min = curr.data
        }
        curr = curr.next
    }
    return
}

// Return the nth element in list. Note that list index count from 0.
func Nth[T comparable](n int, list GoList[T]) (elem T) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.next
    }
    elem = curr.data
    return
}

// Return sublist from the nth element as new list.
func NthTail[T comparable](n int, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.next
    }
    for curr != nil {
        result.Append(curr.data)
        curr = curr.next
    }
    return
}

// Return result is reverse of input list
func Reverse[T comparable](list GoList[T]) (result GoList[T]) {
    var head *node[T]
    curr := list.Head
    for curr != nil {
        node := &node[T]{data: curr.data, next: head}
        head = node
        curr = curr.next
    }
    result = GoList[T]{Head: head}
    return
}

/*
 * Exported methods
 */

 // Append all values into last of list.
func (list *GoList[T]) Append(values ...T) {
    for _, value := range values {
        newNode := &node[T]{data: value}
        if list.Head == nil {
            list.Head = newNode
            continue
        }
        curr := list.Head
        for curr.next != nil {
            curr = curr.next
        }
        curr.next = newNode
    }
}

// Delete first node in list with value of input.
func (list *GoList[T]) Delete(value T) {
    if list.Head == nil {
        return
    }
    // Ignore if it's first node
    if list.Head.data == value {
        list.Head = list.Head.next
        return
    }
    curr := list.Head
    for curr.next != nil {
        if curr.next.data == value {
            curr.next = curr.next.next
            return
        }
        curr = curr.next
    }
}

// Drop the last element in list.
func (list *GoList[T]) DropLast() {
    curr := list.Head
    for curr != nil {
        if curr.next.next == nil {
            curr.next = nil
            break
        }
        curr = curr.next
    }
}

// Drop elements in list while fun returns true.
func (list *GoList[T]) DropWhile(fun func(T) bool) {
    curr := list.Head
    for curr != nil {
        if fun(curr.data) == true {
            list.Head = curr.next
            curr = curr.next
        } else {
            break
        }
    }
}

// Only keep elements in list that fun return true.
func (list *GoList[T]) Filter(fun func(T) bool) {
    if list.Head == nil {
        return
    }
    if fun(list.Head.data) == false {
        list.Head = list.Head.next
    }
    curr := list.Head
    for curr.next != nil {
        if fun(curr.next.data) == false {
            curr.next = curr.next.next
        }
        curr = curr.next
    }
}

// Insert sep between each element in list.
func (list *GoList[T]) Join(sep T) {
    curr := list.Head
    for curr != nil {
        if curr.next != nil {
            node := &node[T]{data: sep, next: curr.next}
            curr.next = node
            curr = node.next
        } else {
            curr = nil
        }
    }
}

// Applying function to every elements in list.
func (list *GoList[T]) Map(fun func(T) T) {
    curr := list.Head
    for curr != nil {
        curr.data = fun(curr.data)
        curr = curr.next
    }
}

// Merge list2 into last of list1
func (list1 *GoList[T]) Merge(list2 GoList[T]) {
    curr1 := list1.Head
    for curr1 != nil {
        if curr1.next == nil {
            curr1.next = list2.Head
            return
        }
        curr1 = curr1.next
    }
}

// Return sublist from the nth element.
func (list *GoList[T]) NthTail(n int) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.next
    }
    list.Head = curr
}

// Reverse the input list
func (list *GoList[T]) Reverse() {
    var prev *node[T]
    curr := list.Head
    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    list.Head = prev
}

// Return a string representing singly linked list.
func (list GoList[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[ ")
    curr := list.Head
    for curr != nil {
        fmt.Fprintf(&builder, "%v", curr.data)
        if curr.next != nil {
            builder.WriteString(" -> ")
        }
        curr = curr.next
    }
    builder.WriteString(" ]")
    return builder.String()
}
