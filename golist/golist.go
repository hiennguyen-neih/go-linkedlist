// Package golist contains functions and methods for singly linked list in Go.
package golist

import (
    "fmt"
    "strings"
    "github.com/hiennguyen-neih/go-linkedlist/node"
    "github.com/hiennguyen-neih/go-linkedlist/constraints"
)

/*
 *******************************************************************************
 * Define structs and interfaces
 *******************************************************************************
 */

// Struct of Go singly linked list.
type GoList[T comparable] struct {
    Head *node.Node[T]
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly linked list.
func New[T comparable](values ...T) (list GoList[T]) {
    list.AppendHead(values...)
    return
}

// Deprecated: Use golist.New instead.
func ListOf[T comparable](values ...T) (list GoList[T]) {
    list = New(values...)
    return
}

// Return true if fun returns true for all elements in list, otherwise return
// false.
func All[T comparable](fun func(T) bool, list GoList[T]) bool {
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Data) {
            return false
        }
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise
// return false.
func Any[T comparable](fun func(T) bool, list GoList[T]) bool {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return true
        }
    }
    return false
}

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) (result GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
    }
    for _, value := range values {
        result.doAppendHead(value)
    }
    result.Reverse()
    return
}

// Return new list is input list that append values into head of it.
func AppendHead[T comparable](list GoList[T], values ...T) (result GoList[T]) {
    for _, value := range values {
        result.doAppendHead(value)
    }
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
    }
    result.Reverse()
    return
}

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    node := list.Head
    for node != nil {
        if node.Data != value {
            result.doAppendHead(node.Data)
            node = node.Next
        } else {
            node = node.Next
            break
        }
    }
    for node != nil {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    result.Reverse()
    return
}

// Drop the last element in list and return as new list.
func DropLast[T comparable](list GoList[T]) (result GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        if node.Next.Next == nil {
            result.doAppendHead(node.Data)
            result.Reverse()
            return
        }
        result.doAppendHead(node.Data)
    }
    result.Reverse()
    return
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    node := list.Head
    for node != nil {
        if !fun(node.Data) {
            break
        }
        node = node.Next
    }
    for node != nil {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    result.Reverse()
    return
}

// Returns a list containing n copies of term element.
func Duplicate[T comparable](n int, elem T) (result GoList[T]) {
    for i := 0; i < n; i++ {
        result.doAppendHead(elem)
    }
    return
}

// Return new list contains elements in input that fun returns true.
func Filter[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.doAppendHead(node.Data)
        }
    }
    result.Reverse()
    return
}

// Execute fun with input is elements in list from left to right and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldl[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    for node := list.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Execute fun with input is elements in list from right to left and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldr[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Applying function for each element in list.
func ForEach[T comparable](fun func(T), list GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        fun(node.Data)
    }
}

// Insert sep between each element in list and return as new list.
func Join[T comparable](sep T, list GoList[T]) (result GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
        if node.Next != nil {
            result.doAppendHead(sep)
        }
    }
    result.Reverse()
    return
}

// Return the last element in list.
func Last[T comparable](list GoList[T]) (last T) {
    for node := list.Head; node != nil; node = node.Next {
        if node.Next == nil {
            last = node.Data
            return
        }
    }
    return
}

// Return length of list.
func Length[T comparable](list GoList[T]) (len int) {
    len = 0
    for node := list.Head; node != nil; node = node.Next {
        len += 1
    }
    return
}

// Applying function to every elements in list and return as new list.
func Map[T comparable](fun func(T) T, list GoList[T]) (result GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(fun(node.Data))
    }
    result.Reverse()
    return
}

// Return maximum element in list.
// This function only works with list of numbers or strings.
func Max[T constraints.Ordered](list GoList[T]) (max T) {
    node := list.Head
    max = node.Data
    for node != nil {
        if node.Data > max {
            max = node.Data
        }
        node = node.Next
    }
    return
}

// Return true if elem in list, otherwise return false.
func Member[T comparable](elem T, list GoList[T]) bool {
    for node := list.Head; node != nil; node = node.Next {
        if node.Data == elem {
            return true
        }
    }
    return false
}

// Return a list that is merged of list1 and list2
func Merge[T comparable](list1 GoList[T], list2 GoList[T]) (result GoList[T]) {
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        result.doAppendHead(node1.Data)
    }
    for node2 := list2.Head; node2 != nil; node2 = node2.Next {
        result.doAppendHead(node2.Data)
    }
    result.Reverse()
    return
}

// Return minimum element in list.
// This functions only works with list of numbers or strings.
func Min[T constraints.Ordered](list GoList[T]) (min T) {
    node := list.Head
    min = node.Data
    for node != nil {
        if node.Data < min {
            min = node.Data
        }
        node = node.Next
    }
    return
}

// Return the nth element in list. Note that list index count from 0.
func Nth[T comparable](n int, list GoList[T]) (elem T) {
    node := list.Head
    for i := 0; i < n; i++ {
        node = node.Next
    }
    elem = node.Data
    return
}

// Return sublist from the nth element as new list.
func NthTail[T comparable](n int, list GoList[T]) (result GoList[T]) {
    node := list.Head
    for i := 0; i < n; i++ {
        node = node.Next
    }
    for node != nil {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    result.Reverse()
    return
}

// Return true if list1 is prefix of list2, otherwise return false.
func Prefix[T comparable](list1 GoList[T], list2 GoList[T]) bool {
    node1 := list1.Head
    node2 := list2.Head
    for node1 != nil {
        if node2 == nil || node1.Data != node2.Data {
            return false
        }
        node1 = node1.Next
        node2 = node2.Next
    }
    return true
}

// Return result is reverse of input list
func Reverse[T comparable](list GoList[T]) (result GoList[T]) {
    var head *node.Node[T]
    for curr := list.Head; curr != nil; curr = curr.Next {
        node := &node.Node[T]{Data: curr.Data, Next: head}
        head = node
    }
    result = GoList[T]{Head: head}
    return
}

// Return position and value of first element that func returns true.
// If every fun execution return false, return pos is -1.
func Search[T comparable](fun func(T) bool, list GoList[T]) (pos int, val T) {
    pos = -1
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            pos = i
            val = node.Data
            return
        }
        i++
    }
    return
}

// Return sequence of numbers that starts with from and contains the successive
// result of adding incr to the previous element, until to is reached or passed.
func Seq[T constraints.Numeric](from, to, incr T) (result GoList[T]) {
    result = GoList[T]{}
    for i := from; i <= to; i += incr {
        result.doAppendHead(i)
    }
    result.Reverse()
    return
}

// Split list into list1 and list2, list1 contains n first elements and list2
// contains the remaining elements.
func Split[T comparable](n int, list GoList[T]) (list1, list2 GoList[T]) {
    node := list.Head
    for i := 0; i < n; i++ {
        list1.doAppendHead(node.Data)
        node = node.Next
    }
    for node != nil {
        list2.doAppendHead(node.Data)
        node = node.Next
    }
    list1.Reverse()
    list2.Reverse()
    return
}

// Split list into list1 and list2, list1 contains elements which fun returns
// true, list2 contains elements which fun returns false.
func SplitWith[T comparable](fun func(T) bool, list GoList[T]) (list1, list2 GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list1.doAppendHead(node.Data)
        } else {
            list2.doAppendHead(node.Data)
        }
    }
    list1.Reverse()
    list2.Reverse()
    return
}

// Return sublist of input list as new list, starting at start and has maximum
// len elements.
func Sublist[T comparable](list GoList[T], start, len int) (result GoList[T]) {
    node := list.Head
    for i := 0; node != nil && i < start; i++ {
        node = node.Next
    }
    for j := 0; node != nil && j < len; j++ {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    result.Reverse()
    return
}

// Return a new list that is a copy of list1 which is for each element in list2,
// its first occurrence in list1 is deleted.
func Subtract[T comparable](list1, list2 GoList[T]) (result GoList[T]) {
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        result.doAppendHead(node1.Data)
    }
    result.Reverse()
    for node2 := list2.Head; node2 != nil; node2 = node2.Next {
        result.Delete(node2.Data)
    }
    return
}

// Returns true if list1 is a suffix of list2, otherwise false.
func Suffix[T comparable](list1, list2 GoList[T]) bool {
    reverse1 := Reverse(list1)
    reverse2 := Reverse(list2)
    return Prefix(reverse1, reverse2)
}

// Returns sum of elements in list.
func Sum[T constraints.Ordered](list GoList[T]) (sum T) {
    for node := list.Head; node != nil; node = node.Next {
        sum += node.Data
    }
    return
}

// Take elements in list while fun returns true, and return as new list.
func TakeWhile[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.doAppendHead(node.Data)
        } else {
            break
        }
    }
    result.Reverse()
    return
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

// Append all values into last of list.
func (list *GoList[T]) Append(values ...T) {
    list.Reverse()
    for _, value := range values {
        list.doAppendHead(value)
    }
    list.Reverse()
}

// Append all values into head of list.
func (list *GoList[T]) AppendHead(values ...T) {
    for i := len(values) - 1; i >= 0; i-- {
        list.doAppendHead(values[i])
    }
}

// Delete first node in list with value of input.
func (list *GoList[T]) Delete(value T) {
    if list.Head == nil {
        return
    }
    // Ignore if it's first node
    if list.Head.Data == value {
        list.Head = list.Head.Next
        return
    }
    node := list.Head
    for node.Next != nil {
        if node.Next.Data == value {
            node.Next = node.Next.Next
            return
        }
        node = node.Next
    }
}

// Drop the last element in list.
func (list *GoList[T]) DropLast() {
    node := list.Head
    for node != nil {
        if node.Next.Next == nil {
            node.Next = nil
            break
        }
        node = node.Next
    }
}

// Drop elements in list while fun returns true.
func (list *GoList[T]) DropWhile(fun func(T) bool) {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list.Head = node.Next
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
    if !fun(list.Head.Data) {
        list.Head = list.Head.Next
    }
    for node := list.Head; node.Next != nil; node = node.Next {
        if !fun(node.Next.Data) {
            node.Next = node.Next.Next
        }
    }
}

// Insert sep between each element in list.
func (list *GoList[T]) Join(sep T) {
    curr := list.Head
    for curr != nil {
        if curr.Next != nil {
            node := &node.Node[T]{Data: sep, Next: curr.Next}
            curr.Next = node
            curr = node.Next
        } else {
            curr = nil
        }
    }
}

// Applying function to every elements in list.
func (list *GoList[T]) Map(fun func(T) T) {
    for node := list.Head; node != nil; node = node.Next {
        node.Data = fun(node.Data)
    }
}

// Merge list2 into last of list1
func (list1 *GoList[T]) Merge(list2 GoList[T]) {
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        if node1.Next == nil {
            node1.Next = list2.Head
            return
        }
    }
}

// Return sublist from the nth element.
func (list *GoList[T]) NthTail(n int) {
    node := list.Head
    for i := 0; i < n; i++ {
        node = node.Next
    }
    list.Head = node
}

// Reverse the input list
func (list *GoList[T]) Reverse() {
    var prev *node.Node[T]
    node := list.Head
    for node != nil {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
    }
    list.Head = prev
}

// Return sublist that starting at start and has maximum len elements.
func (list *GoList[T]) Sublist(start, len int) {
    node := list.Head
    for i := 0; node != nil && i < start; i++ {
        node = node.Next
    }
    list.Head = node
    for j := 0; node != nil && node.Next != nil && j < len - 1; j++ {
        node = node.Next
    }
    if node != nil {
        node.Next = nil
    }
    return
}

// Delete elements in list1 that is its first occurrence to each element in list2.
func (list1 *GoList[T]) Subtract(list2 GoList[T]) {
    for node := list2.Head; node != nil; node = node.Next {
        list1.Delete(node.Data)
    }
}

// Take elements in list while fun returns true.
func (list *GoList[T]) TakeWhile(fun func(T) bool) {
    if !fun(list.Head.Data) {
        list.Head = nil
        return
    }
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Next.Data) {
            node.Next = nil
        }
    }
}

// Return a string representing singly linked list.
func (list GoList[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[ ")
    for node := list.Head; node != nil; node = node.Next {
        fmt.Fprintf(&builder, "%v", node.Data)
        if node.Next != nil {
            builder.WriteString(" -> ")
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

func (list *GoList[T]) doAppendHead(value T) {
    node := &node.Node[T]{Data: value, Next: list.Head}
    list.Head = node
}