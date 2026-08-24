// Package golist contains functions and methods for singly circular linked
// list in Go.
package golistc

import (
    "fmt"
    "strings"

    "github.com/google/go-cmp/cmp"
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
    Head *node.Node[T] // First node of the list.
    Tail *node.Node[T] // Last node of the list.
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
    if node == nil {
        return result
    }
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

// Returns a list that is concatenated of all input lists.
func Concat[T any](lists ...GoListC[T]) GoListC[T] {
    var result GoListC[T]
    for _, list := range lists {
        node := list.Head
        for {
            result.append(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }
    return result
}

// Returns a copy of input list where the first node data that matching value
// is removed.
func Delete[T any](list GoListC[T], value T) GoListC[T] {
    var result GoListC[T]
    if list.Head == nil {
        return result
    }
    node := list.Head
    for {
        if !cmp.Equal(node.Data, value) {
            result.append(node.Data)
            node = node.Next
        } else {
            node = node.Next
            break
        }
        if node == list.Head {
            break
        }
    }
    for node != list.Head {
        result.append(node.Data)
        node = node.Next
    }
    return result
}

// Deletes node at the specific index of list. If index is out of bound, the
// original list is returned. Negative index indicate an offset from the end
// of list.
func DeleteAt[T any](list GoListC[T], index int) GoListC[T] {
    var result GoListC[T]
    len := Len(list)
    if list.Head == nil {
        return result
    }
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index >= len {
        for node := list.Head; ; {
            result.append(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
        }
    } else {
        node := list.Head
        for i := 0; ; {
            if i == index {
                i++
                node = node.Next
                continue
            }
            i++
            result.append(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }
    return result
}

// Drops the last node of input list. If input list is an empty list, returns
// an empty list.
func DropLast[T any](list GoListC[T]) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil {
        return result
    }
    for {
        if node.Next == list.Tail {
            result.append(node.Data)
            break
        }
        result.append(node.Data)
        node = node.Next
    }
    return result
}

// Drops nodes from list while fun returns true.
func DropWhile[T any](list GoListC[T], fun func(T) bool) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil {
        return result
    }
    for {
        if !fun(node.Data) {
            break
        }
        node = node.Next
        if node == list.Head {
            return result
        }
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

// Returns a list containing n copies of term elem. If n is negative or equal
// 0, return empty list.
func Duplicate[T any](n int, elem T) GoListC[T] {
    var result GoListC[T]
    for i := 0; i < n; i++ {
        result.append(elem)
    }
    return result
}

// Returns true if all corresponding nodes in both list1 and list2 have the
// same value, otherwise return false.
func Equal[T any](list1, list2 GoListC[T]) bool {
    node1 := list1.Head
    node2 := list2.Head
    for {
        if !cmp.Equal(node1.Data, node2.Data) {
            return false
        }
        node1 = node1.Next
        node2 = node2.Next
        if (node1 == list1.Head && node2 != list2.Head) || (node1 != list1.Head && node2 == list2.Head) {
            return false
        } else if node1 == list1.Head && node2 == list2.Head {
            break
        }
    }
    return true
}

// Returns a list contains node data from input list for which fun returns true.
func Filter[T any](list GoListC[T], fun func(T) bool) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil {
        return result
    }
    for {
        if fun(node.Data) {
            result.append(node.Data)
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Calls fun on successive nodes of list to update or remove nodes from list.
// Input fun must return (bool, value). The functions returns a list that nodes
// data are value in which fun returns (true, value).
func FilterMap[T any](list GoListC[T], fun func(T) (bool, T)) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil {
        return result
    }
    for {
        if keep, value := fun(node.Data); keep {
            result.append(value)
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Returns position of first node of list that match with value. If there is
// no matching node, returns -1.
func Find[T any](list GoListC[T], value T) int {
    i := 0
    node := list.Head
    if node == nil {
        return -1
    }
    for {
        if cmp.Equal(node.Data, value) {
            return i
        }
        i++
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return -1
}

// Calls fun(data, acc) on successive nodes of list from left to right (from
// start of list to end of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldl[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    for node := list.Head; ; {
        acc0 = fun(node.Data, acc0)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return acc0
}

// Calls fun(data, acc) on successive nodes of list from right to left (from
// end of list to start of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldr[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    reverse := Reverse(list)
    for node := reverse.Head; ; {
        acc0 = fun(node.Data, acc0)
        node = node.Next
        if node == reverse.Head {
            break
        }
    }
    return acc0
}

// Calls fun(data) for each node in list, ignoring the return value. This
// function is used for its side effects and the evaluation order is defined
// to be the same as the order of the nodes in the list.
func ForEach[T any](list GoListC[T], fun func(T)) {
    for node := list.Head; ; {
        fun(node.Data)
        node = node.Next
        if node == list.Head {
            break
        }
    }
}

// Returns a list with val is inserted at specific index. index is capped at
// list length. Negative index indicate an offset from the end of list.
func InsertAt[T any](list GoListC[T], index int, val T) GoListC[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index > len {
        panic("InsertAt, index is out of bound!")
    }

    var result GoListC[T]
    if index == len {
        for node := list.Head; ; {
            result.append(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
        }
        result.append(val)
    } else {
        i := 0
        for node := list.Head; ; {
            if i == index {
                result.append(val)
            }
            result.append(node.Data)
            i++
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }

    return result
}

// Inserts sep between each node in list. This function has no effect on an
// empty list or a singleton list.
func Join[T any](list GoListC[T], sep T) GoListC[T] {
    var result GoListC[T]
    for node := list.Head; ; {
        result.append(node.Data)
        if node.Next != list.Head {
            result.append(sep)
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Returns the length of list.
func Len[T any](list GoListC[T]) int {
    len := 0
    node := list.Head
    if node == nil {
        return len
    }
    for {
        len += 1
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return len
}

// Calls fun(data) to every nodes in list and returns a list contains returned
// values of that fun.
func Map[T any](list GoListC[T], fun func(T) T) GoListC[T] {
    var result GoListC[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; ; {
        result.append(fun(node.Data))
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Returns a list containing the nodes of input list in reverse order.
func Reverse[T any](list GoListC[T]) GoListC[T] {
    var result GoListC[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; ; {
        result.appendHead(node.Data)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Takes nodes data in list while fun returns true, returning the longest
// prefix in which all nodes data satisfy the predicate.
func TakeWhile[T any](list GoListC[T], fun func(T) bool) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil {
        return result
    }
    for {
        if fun(node.Data) {
            result.append(node.Data)
        } else {
            break
        }
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
        if node == nil {
            break
        }
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

// Do append value into tail of list.
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

// Do append value into head of list.
func (list *GoListC[T]) appendHead(value T) *GoListC[T] {
    node := &node.Node[T]{Data: value}
    if list.Head == nil {
        list.Head = node
        list.Tail = node
        node.Next = node
    } else {
        node.Next = list.Head
        list.Head = node
        list.Tail.Next = node
    }
    return list
}

// Do reverse the list.
// func (list *GoListC[T]) reverse() *GoListC[T] {
//     prev := list.Tail
//     node := list.Head
//     for {
//         next := node.Next
//         node.Next = prev
//         prev = node
//         node = next
//         if node == list.Head {
//             break
//         }
//     }
//     list.Head = prev
//     list.Tail = node
//     return list
// }
