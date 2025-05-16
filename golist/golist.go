// Package golist contains functions and methods for singly linked list in Go.
package golist

import (
    "fmt"
    "strings"
)

/*
 * Define data types
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
    current := list.Head
    for current != nil {
        if fun(current.data) == false {
            return false
        }
        current = current.next
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise return false.
func Any[T comparable](fun func(T) bool, list GoList[T]) bool {
    current := list.Head
    for current != nil {
        if fun(current.data) == true {
            return true
        }
        current = current.next
    }
    return false
}

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) (result GoList[T]) {    
    current := list.Head
    for current != nil {
        result.Append(current.data)
        current = current.next
    }
    result.Append(values...)
    return
}

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    current := list.Head
    for current != nil {
        if current.data != value {
            result.Append(current.data)
            current = current.next
        } else {
            current = current.next
            break
        }
    }
    for current != nil {
        result.Append(current.data)
        current = current.next
    }
    return
}

// Drop the last element in list and return as new list.
func DropLast[T comparable](list GoList[T]) (result GoList[T]) {
    current := list.Head
    for current != nil {
        if current.next.next == nil {
            result.Append(current.data)
            return
        }
        result.Append(current.data)
        current = current.next
    }
    return
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    current := list.Head
    for current != nil {
        if fun(current.data) == false {
            break
        }
        current = current.next
    }
    for current != nil {
        result.Append(current.data)
        current = current.next
    }
    return
}

// Returns a list containing n copies of term element.
func Duplicate[T comparable](n int, element T) (result GoList[T]) {
    for i := 0; i < n; i++ {
        result.Append(element)
    }
    return
}

// Return new list contains elements in input that fun returns true.
func Filter[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    current := list.Head
    for current != nil {
        if fun(current.data) == true {
            result.Append(current.data)
        }
        current = current.next
    }
    return
}

// Applying function to every elements in list and return as new list.
func Map[T comparable](fun func(T) T, list GoList[T]) (result GoList[T]) {
    current := list.Head
    for current != nil {
        result.Append(fun(current.data))
        current = current.next
    }
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
        current := list.Head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
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
    current := list.Head
    for current.next != nil {
        if current.next.data == value {
            current.next = current.next.next
            return
        }
        current = current.next
    }
}

// Drop the last element in list.
func (list *GoList[T]) DropLast() {
    current := list.Head
    for current != nil {
        if current.next.next == nil {
            current.next = nil
            break
        }
        current = current.next
    }
}

// Drop elements in list while fun returns true.
func (list *GoList[T]) DropWhile(fun func(T) bool) {
    current := list.Head
    for current != nil {
        if fun(current.data) == true {
            list.Head = current.next
            current = current.next
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
    current := list.Head
    for current.next != nil {
        if fun(current.next.data) == false {
            current.next = current.next.next
        }
        current = current.next
    }
}

// Applying function to every elements in list.
func (list *GoList[T]) Map(fun func(T) T) {
    current := list.Head
    for current != nil {
        current.data = fun(current.data)
        current = current.next
    }
}

// Return a string representing singly linked list.
func (list GoList[T]) String() string {
    var builder strings.Builder
    current := list.Head
    for current != nil {
        fmt.Fprintf(&builder, "%v", current.data)
        builder.WriteString(" -> ")
        current = current.next
    }
    builder.WriteString("nil")
    return builder.String()
}
