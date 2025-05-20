// Package golist contains functions and methods for singly linked list in Go.
package golist

import (
    "fmt"
    "strings"
)

/*
 *******************************************************************************
 * Define structs and interfaces
 *******************************************************************************
 */

// Node in singly linked List.
type Node[T comparable] struct {
    Data T
    Next *Node[T]
}

// Struct of Go singly linked list.
type GoList[T comparable] struct {
    Head *Node[T]
}

type ordered interface {
    number | ~string
}

type number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly linked list.
func ListOf[T comparable](values ...T) (list GoList[T]) {
    list.Append(values...)
    return
}

// Return true if fun returns true for all elements in list, otherwise return false.
func All[T comparable](fun func(T) bool, list GoList[T]) bool {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if !fun(curr.Data) {
            return false
        }
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise return false.
func Any[T comparable](fun func(T) bool, list GoList[T]) bool {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if fun(curr.Data) {
            return true
        }
    }
    return false
}

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) (result GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        result.Append(curr.Data)
    }
    result.Append(values...)
    return
}

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    curr := list.Head
    for curr != nil {
        if curr.Data != value {
            result.Append(curr.Data)
            curr = curr.Next
        } else {
            curr = curr.Next
            break
        }
    }
    for curr != nil {
        result.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Drop the last element in list and return as new list.
func DropLast[T comparable](list GoList[T]) (result GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if curr.Next.Next == nil {
            result.Append(curr.Data)
            return
        }
        result.Append(curr.Data)
    }
    return
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for curr != nil {
        if !fun(curr.Data) {
            break
        }
        curr = curr.Next
    }
    for curr != nil {
        result.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Returns a list containing n copies of term element.
func Duplicate[T comparable](n int, elem T) (result GoList[T]) {
    for i := 0; i < n; i++ {
        result.Append(elem)
    }
    return
}

// Return new list contains elements in input that fun returns true.
func Filter[T comparable](fun func(T) bool, list GoList[T]) (result GoList[T]) {
    if list.Head == nil {
        return
    }
    for curr := list.Head; curr != nil; curr = curr.Next {
        if fun(curr.Data) {
            result.Append(curr.Data)
        }
    }
    return
}

// Execute fun with input is elements in list from left to right and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldl[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    for curr := list.Head; curr != nil; curr = curr.Next {
        acc0 = fun(curr.Data, acc0)
    }
    return acc0
}

// Execute fun with input is elements in list from right to left and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldr[T1, T2 comparable](fun func(T1, T2) T2, acc0 T2, list GoList[T1]) T2 {
    reverse := Reverse(list)
    for curr := reverse.Head; curr != nil; curr = curr.Next {
        acc0 = fun(curr.Data, acc0)
    }
    return acc0
}

// Applying function for each element in list.
func ForEach[T comparable](fun func(T), list GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        fun(curr.Data)
    }
}

// Insert sep between each element in list and return as new list.
func Join[T comparable](sep T, list GoList[T]) (result GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        result.Append(curr.Data)
        if curr.Next != nil {
            result.Append(sep)
        }
    }
    return
}

// Return the last element in list.
func Last[T comparable](list GoList[T]) (last T) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if curr.Next == nil {
            last = curr.Data
            return
        }
    }
    return
}

// Return length of list.
func Length[T comparable](list GoList[T]) (len int) {
    len = 0
    for curr := list.Head; curr != nil; curr = curr.Next {
        len += 1
    }
    return
}

// Applying function to every elements in list and return as new list.
func Map[T comparable](fun func(T) T, list GoList[T]) (result GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        result.Append(fun(curr.Data))
    }
    return
}

// Return maximum element in list.
// This function only works with list of numbers or strings.
func Max[T ordered](list GoList[T]) (max T) {
    curr := list.Head
    max = curr.Data
    for curr != nil {
        if curr.Data > max {
            max = curr.Data
        }
        curr = curr.Next
    }
    return
}

// Return true if elem in list, otherwise return false.
func Member[T comparable](elem T, list GoList[T]) bool {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if curr.Data == elem {
            return true
        }
    }
    return false
}

// Return a list that is merged of list1 and list2
func Merge[T comparable](list1 GoList[T], list2 GoList[T]) (result GoList[T]) {
    curr := list1.Head
    for curr != nil {
        result.Append(curr.Data)
        curr = curr.Next
    }
    curr = list2.Head
    for curr != nil {
        result.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Return minimum element in list.
// This functions only works with list of numbers or strings.
func Min[T ordered](list GoList[T]) (min T) {
    curr := list.Head
    min = curr.Data
    for curr != nil {
        if curr.Data < min {
            min = curr.Data
        }
        curr = curr.Next
    }
    return
}

// Return the nth element in list. Note that list index count from 0.
func Nth[T comparable](n int, list GoList[T]) (elem T) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.Next
    }
    elem = curr.Data
    return
}

// Return sublist from the nth element as new list.
func NthTail[T comparable](n int, list GoList[T]) (result GoList[T]) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.Next
    }
    for curr != nil {
        result.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Return true if list1 is prefix of list2, otherwise return false.
func Prefix[T comparable](list1 GoList[T], list2 GoList[T]) bool {
    curr1 := list1.Head
    curr2 := list2.Head
    for curr1 != nil {
        if curr2 == nil || curr1.Data != curr2.Data {
            return false
        }
        curr1 = curr1.Next
        curr2 = curr2.Next
    }
    return true
}

// Return result is reverse of input list
func Reverse[T comparable](list GoList[T]) (result GoList[T]) {
    var head *Node[T]
    for curr := list.Head; curr != nil; curr = curr.Next {
        node := &Node[T]{Data: curr.Data, Next: head}
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
    for curr := list.Head; curr != nil; curr = curr.Next {
        if fun(curr.Data) {
            pos = i
            val = curr.Data
            return
        }
        i++
    }
    return
}

// Return sequence of numbers that starts with from and contains the successive
// result of adding incr to the previous element, until to is reached or passed.
func Seq[T number](from, to, incr T) (result GoList[T]) {
    result = GoList[T]{}
    for i := from; i <= to; i += incr {
        result.Append(i)
    }
    return
}

// Split list into list1 and list2, list1 contains n first elements and list2
// contains the remaining elements.
func Split[T comparable](n int, list GoList[T]) (list1, list2 GoList[T]) {
    curr := list.Head
    for i := 0; i < n; i++ {
        list1.Append(curr.Data)
        curr = curr.Next
    }
    for curr != nil {
        list2.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Split list into list1 and list2, list1 contains elements which fun returns
// true, list2 contains elements which fun returns false.
func SplitWith[T comparable](fun func(T) bool, list GoList[T]) (list1, list2 GoList[T]) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if fun(curr.Data) {
            list1.Append(curr.Data)
        } else {
            list2.Append(curr.Data)
        }
    }
    return
}

// Return sublist of input list as new list, starting at start and has maximum
// len elements.
func Sublist[T comparable](list GoList[T], start, len int) (result GoList[T]) {
    curr := list.Head
    for i := 0; curr != nil && i < start; i++ {
        curr = curr.Next
    }
    for j := 0; curr != nil && j < len; j++ {
        result.Append(curr.Data)
        curr = curr.Next
    }
    return
}

// Return a new list that is a copy of list1 which is for each element in list2,
// its first occurrence in list1 is deleted.
func Subtract[T comparable](list1, list2 GoList[T]) (result GoList[T]) {
    for curr1 := list1.Head; curr1 != nil; curr1 = curr1.Next {
        result.Append(curr1.Data)
    }
    for curr2 := list2.Head; curr2 != nil; curr2 = curr2.Next {
        result.Delete(curr2.Data)
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
func Sum[T ordered](list GoList[T]) (sum T) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        sum += curr.Data
    }
    return
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

 // Append all values into last of list.
func (list *GoList[T]) Append(values ...T) {
    for _, value := range values {
        newNode := &Node[T]{Data: value}
        if list.Head == nil {
            list.Head = newNode
            continue
        }
        curr := list.Head
        for curr.Next != nil {
            curr = curr.Next
        }
        curr.Next = newNode
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
    curr := list.Head
    for curr.Next != nil {
        if curr.Next.Data == value {
            curr.Next = curr.Next.Next
            return
        }
        curr = curr.Next
    }
}

// Drop the last element in list.
func (list *GoList[T]) DropLast() {
    curr := list.Head
    for curr != nil {
        if curr.Next.Next == nil {
            curr.Next = nil
            break
        }
        curr = curr.Next
    }
}

// Drop elements in list while fun returns true.
func (list *GoList[T]) DropWhile(fun func(T) bool) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        if fun(curr.Data) {
            list.Head = curr.Next
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
    for curr := list.Head; curr.Next != nil; curr = curr.Next {
        if !fun(curr.Next.Data) {
            curr.Next = curr.Next.Next
        }
    }
}

// Insert sep between each element in list.
func (list *GoList[T]) Join(sep T) {
    curr := list.Head
    for curr != nil {
        if curr.Next != nil {
            node := &Node[T]{Data: sep, Next: curr.Next}
            curr.Next = node
            curr = node.Next
        } else {
            curr = nil
        }
    }
}

// Applying function to every elements in list.
func (list *GoList[T]) Map(fun func(T) T) {
    for curr := list.Head; curr != nil; curr = curr.Next {
        curr.Data = fun(curr.Data)
    }
}

// Merge list2 into last of list1
func (list1 *GoList[T]) Merge(list2 GoList[T]) {
    for curr1 := list1.Head; curr1 != nil; curr1 = curr1.Next {
        if curr1.Next == nil {
            curr1.Next = list2.Head
            return
        }
    }
}

// Return sublist from the nth element.
func (list *GoList[T]) NthTail(n int) {
    curr := list.Head
    for i := 0; i < n; i++ {
        curr = curr.Next
    }
    list.Head = curr
}

// Reverse the input list
func (list *GoList[T]) Reverse() {
    var prev *Node[T]
    curr := list.Head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    list.Head = prev
}

// Return sublist that starting at start and has maximum len elements.
func (list *GoList[T]) Sublist(start, len int) {
    curr := list.Head
    for i := 0; curr != nil && i < start; i++ {
        curr = curr.Next
    }
    list.Head = curr
    for j := 0; curr != nil && curr.Next != nil && j < len - 1; j++ {
        curr = curr.Next
    }
    if curr != nil {
        curr.Next = nil
    }
    return
}

// Delete elements in list1 that is its first occurrence to each element in list2.
func (list1 *GoList[T]) Subtract(list2 GoList[T]) {
    for curr := list2.Head; curr != nil; curr = curr.Next {
        list1.Delete(curr.Data)
    }
}

// Return a string representing singly linked list.
func (list GoList[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[ ")
    for curr := list.Head; curr != nil; curr = curr.Next {
        fmt.Fprintf(&builder, "%v", curr.Data)
        if curr.Next != nil {
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

