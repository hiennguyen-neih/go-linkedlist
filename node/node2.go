// Package node contains definition for nodes in go-linkedlist.
package node

import (
    "fmt"
)

// Node in singly linked List.
type Node2[T any] struct {
    Prev *Node2[T]    // Pointer to previous node of the list.
    Data T
    Next *Node2[T]    // Pointer to next node of the list.
}

// Return a string representing node in linked list.
func (node Node2[T]) String() string {
    var data any = node.Data
    if str, ok := data.(string); ok {
        return fmt.Sprintf("%q", str)
    }
    return fmt.Sprintf("%v", node.Data)
}
