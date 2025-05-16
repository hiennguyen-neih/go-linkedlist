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

// Create new linked list.
func ListOf[T comparable](values ...T) GoList[T] {
	list := GoList[T]{}
    list.Append(values...)
	return list
}

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

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) GoList[T] {
    result := GoList[T]{}
    
    current := list.Head
    for current != nil {
        result.Append(current.data)
        current = current.next
    }

    result.Append(values...)
    return result
}

// Applying function to every elements in list
func (list *GoList[T]) Map(fun func(T) T) {
    current := list.Head
    for current != nil {
        current.data = fun(current.data)
        current = current.next
    }
}

// Applying function to every elements in list and return as new list
func Map[T comparable](list GoList[T], fun func(T) T) GoList[T] {
    result := GoList[T]{}

    current := list.Head
    for current != nil {
        result.Append(fun(current.data))
        current = current.next
    }

    return result
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

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) GoList[T] {
    if list.Head == nil {
        return list
    }

    result := GoList[T]{}
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

    return result
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
