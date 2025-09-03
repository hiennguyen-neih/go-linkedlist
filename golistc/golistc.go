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

// Returns true if fun returns true for all node data in list, otherwise returns
// false.
func All[T any](list GoListC[T], fun func(T) bool) bool {
    node := list.Head
    for {
        if !fun(node.Data) {
            return false
        }

        node = node.Next
        if node == list.Head {
            break
        }
    }
    return true
}

// Returns true if fun returns true for at least 1 node data in list, otherwise
// returns false.
func Any[T any](list GoListC[T], fun func(T) bool) bool {
    node := list.Head
    for {
        if fun(node.Data) {
            return true
        }

        node = node.Next
        if node == list.Head {
            break
        }
    }
    return false
}

// Appends values into last of input list.
func Append[T any](list GoListC[T], values ...T) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    for {
        result.append(node.Data)

        node = node.Next
        if node == list.Head {
            break
        }
    }
    for _, value := range values {
        result.append(value)
    }
    return result
}

// Appends values into head of input list.
func AppendHead[T any](list GoListC[T], values ...T) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    for _, value := range values {
        result.append(value)
    }

    for {
        result.append(node.Data)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Returns a list containing the nodes of input list in reverse order.
func Reverse[T any](list GoListC[T]) GoListC[T] {
    var head *node.Node[T]
    for {
        curr := list.Head
        node := &node.Node[T]{Data: curr.Data, Next: head}
        head = node

        curr = curr.Next
        if curr == list.Head {
            break
        }
    }
    return GoListC[T]{Head: head}
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

// Do reverse the list.
func (list *GoListC[T]) reverse() *GoListC[T] {
    prev := list.Tail
    node := list.Head
    for {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
        if node == list.Head {
            break
        }
    }
    list.Head = prev
    list.Tail = node
    return list
}
