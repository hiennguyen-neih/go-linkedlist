// Package golist contains functions and methods for singly circular linked
// list in Go.
package golistc

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

// Struct of Go singly linked list.
type GoListC[T any] struct {
    Head *node.Node[T]    // First node of the list.
    Tail *node.Node[T]    // Last node of the list.
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly linked list from input values.
func New[T any](values ...T) GoListC[T] {
    var list GoListC[T]
    for _, val := range values {
        list.append(val)
    }
    return list
}

// Convert input slice into new singly linked list.
func FromSlice[T any](values []T) GoListC[T] {
    var list GoListC[T]
    for _, val := range values {
        list.append(val)
    }
    return list
}

// Convert input singly linked list into new slice.
func ToSlice[T any](list GoListC[T]) []T {
    var result []T
    node := list.Head
    for {
        result = append(result, node.Data)

        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

// Returns a string representing the singly linked list.
func (list GoListC[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[")
    node := list.Head
    for {
        var data any = node.Data
        if str, ok := data.(string); ok {
            fmt.Fprintf(&builder, "%q", str)
        } else {
            fmt.Fprintf(&builder, "%v", node.Data)
        }
        builder.WriteString("=>")

        node = node.Next
        if node == list.Head {
            break
        }
    }
    builder.WriteString("]")
    return builder.String()
}

/*
 *******************************************************************************
 * Internal functions and methods
 *******************************************************************************
 */

// Do append value into head of list.
func (list *GoListC[T]) append(value T) *GoListC[T] {
    node := &node.Node[T]{Data: value}
    if list.Head != nil {
        node.Next = list.Head
        list.Tail.Next = node
        list.Tail = node
    } else {
        list.Head = node
        list.Tail = node
        node.Next = node
    }
    return list
}
