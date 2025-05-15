// Package golist contains methods and functions for singly linked list in Go
package golist

import (
	"fmt"
    "strings"
)

/*
 * Define data types
 */

// Node in Linked List
type node[T comparable] struct {
    data T
    next *node[T]
}

// Struct of singly linked list
type LinkedList[T comparable] struct {
    Node *node[T]
}

/*
 * Exported functions
 */

// Create new linked list
func ListOf[T comparable](values ...T) LinkedList[T] {
	list := LinkedList[T]{}
    list.Append(values...)
	return list
}

// Method add values into last of current list
func (list *LinkedList[T]) Append(values ...T) {
    for _, value := range values {
        newNode := &node[T]{data: value}

        if list.Node == nil {
            list.Node = newNode
            continue
        }
        
        current := list.Node
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// Function add values into last of current list and return as new list
func Append[T comparable](list LinkedList[T], values ...T) LinkedList[T] {
    result := LinkedList[T]{}
    
    current := list.Node
    for current != nil {
        result.Append(current.data)
        current = current.next
    }

    result.Append(values...)
    return result
}

// Method delete node in list by value
func (list *LinkedList[T]) Delete(value T) {
    if list.Node == nil {
        return
    }

    // Ignore if it's first node
    if list.Node.data == value {
        list.Node = list.Node.next
        return
    }

    current := list.Node
    for current.next != nil {
        if current.next.data == value {
            current.next = current.next.next
            return
        }
        current = current.next
    }
}

// Function delete node in list by value and return new list
func Delete[T comparable](list LinkedList[T], value T) LinkedList[T] {
    if list.Node == nil {
        return list
    }

    result := LinkedList[T]{}
    current := list.Node
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

// Method ToString
func (list *LinkedList[T]) ToString() string {
    var builder strings.Builder
    current := list.Node
    for current != nil {
        fmt.Fprintf(&builder, "%v", current.data)
        builder.WriteString(" -> ")
        current = current.next
    }
    builder.WriteString("nil")
    return builder.String()
}
