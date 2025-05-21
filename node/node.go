// Package node contains definition for nodes in go-linkedlist.
package node

import (
    "fmt"
)

// Node in singly linked List.
type Node[T comparable] struct {
    Data T
    Next *Node[T]
}

// Return a string representing node in linked list.
func (node Node[T]) String() string {
    return fmt.Sprintf("%v", node.Data)
}
