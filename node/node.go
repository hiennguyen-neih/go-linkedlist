// Package node contains definition for nodes in go-linkedlist.
package node

import (
    "fmt"
)

// Node in singly linked List.
type Node[T any] struct {
    Data T
    Next *Node[T]    // Pointer to next node of the list.
}

// Return a string representing node in linked list.
func (node Node[T]) String() string {
    var data any = node.Data
    if str, ok := data.(string); ok {
        return fmt.Sprintf("%q", str)
    }
    return fmt.Sprintf("%v", node.Data)
}
