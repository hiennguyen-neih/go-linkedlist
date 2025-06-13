// Package golist2 contains functions and methods for doubly linked list in Go.
package golist2

import (
    "fmt"
    "strings"
    // "github.com/google/go-cmp/cmp"
    "github.com/hiennguyen-neih/go-linkedlist/node"
    // "github.com/hiennguyen-neih/go-linkedlist/constraints"
)

/*
 *******************************************************************************
 * Define structs and interfaces
 *******************************************************************************
 */

// Struct of Go doubly linked list.
type GoList2[T any] struct {
    Head *node.Node2[T]    // First node of the list.
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly linked list from input values.
func New[T any](values ...T) GoList2[T] {
    var list GoList2[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input slice into new singly linked list.
func FromSlice[T any](values []T) GoList2[T] {
    var list GoList2[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input singly linked list into new slice.
func ToSlice[T any](list GoList2[T]) []T {
    var result []T
    for node := list.Head; node != nil; node = node.Next {
        result = append(result, node.Data)
    }
    return result
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

// Returns a string representing the singly linked list.
func (list GoList2[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[ ")
    for node := list.Head; node != nil; node = node.Next {
        var data any = node.Data
        if str, ok := data.(string); ok {
            fmt.Fprintf(&builder, "%q", str)
        } else {
            fmt.Fprintf(&builder, "%v", node.Data)
        }
        if node.Next != nil {
            builder.WriteString(" <-> ")
        }
    }
    builder.WriteString(" ]")
    return builder.String()
}

/*
 *******************************************************************************
 * Internal functions and methods
 *******************************************************************************
 */

// Do append value into head of list.
func (list *GoList2[T]) appendHead(value T) *GoList2[T] {
    node := &node.Node2[T]{Data: value, Next: list.Head}
    if list.Head != nil {
        list.Head.Prev = node
    }
    list.Head = node
    return list
}

// Do reverse the list.
func (list *GoList2[T]) reverse() *GoList2[T] {
    var prev *node.Node2[T]
    for node := list.Head; node != nil; node = node.Prev {
        node.Prev, node.Next = node.Next, node.Prev
        prev = node
    }
    list.Head = prev
    return list
}
