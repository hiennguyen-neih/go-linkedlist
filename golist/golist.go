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
func New[T comparable](values ...T) GoList[T] {
    var list GoList[T]
    return *list.AppendHead(values...)
}

// Return true if fun returns true for all elements in list, otherwise return
// false.
func All[T comparable](list GoList[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Data) {
            return false
        }
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise
// return false.
func Any[T comparable](list GoList[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return true
        }
    }
    return false
}

// Return new list is append of input list and values.
func Append[T comparable](list GoList[T], values ...T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
    }
    for _, value := range values {
        result.doAppendHead(value)
    }
    return *result.Reverse()
}

// Return new list is input list that append values into head of it.
func AppendHead[T comparable](list GoList[T], values ...T) GoList[T] {
    var result GoList[T]
    for _, value := range values {
        result.doAppendHead(value)
    }
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
    }
    return *result.Reverse()
}

// Return a list that is concatenated of list1 and list2.
func Concat[T comparable](list1, list2 GoList[T]) GoList[T] {
    var result GoList[T]
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        result.doAppendHead(node1.Data)
    }
    for node2 := list2.Head; node2 != nil; node2 = node2.Next {
        result.doAppendHead(node2.Data)
    }
    return *result.Reverse()
}

// Delte first node in list with value of input and return as new list.
func Delete[T comparable](list GoList[T], value T) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
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
    return *result.Reverse()
}

// Delete node at the specific index in list and return as new list. Negative
// index indicate an offset from the end of list.
func DeleteAt[T comparable](list GoList[T], index int) GoList[T] {
    var result GoList[T]
    len := Len(list)

    if list.Head == nil {
        return result
    }

    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index < 0 || index >= len {
        for node := list.Head; node != nil; node = node.Next {
            result.doAppendHead(node.Data)
        }
    } else {
        i := 0
        for node := list.Head; node != nil; node = node.Next {
            if i == index {
                i++
                continue
            }
            i++
            result.doAppendHead(node.Data)
        }
    }

    return *result.Reverse()
}

// Drop the last element in list and return as new list.
func DropLast[T comparable](list GoList[T]) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if node.Next.Next == nil {
            result.doAppendHead(node.Data)
            return *result.Reverse()
        }
        result.doAppendHead(node.Data)
    }
    return *result.Reverse()
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T comparable](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
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
    return *result.Reverse()
}

// Return a list containing n copies of term element. If n is negative or equal
// 0, return empty list.
func Duplicate[T comparable](n int, elem T) GoList[T] {
    var result GoList[T]
    for i := 0; i < n; i++ {
        result.doAppendHead(elem)
    }
    return result
}

// Return new list contains elements in input that fun returns true.
func Filter[T comparable](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.doAppendHead(node.Data)
        }
    }
    return *result.Reverse()
}

// Calls fun on successive elements of list. fun must return (bool, value).
// The function returns a new list of elements for which fun returns
// a new value, where a value of true is synonymous with (true, value).
func FilterMap[T comparable](list GoList[T], fun func(T) (bool, T)) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; node != nil; node = node.Next {
        if keep, value := fun(node.Data); keep {
            result.doAppendHead(value)
        }
    }
    return *result.Reverse()
}

// Return position of first element in list that match with value. If there is
// no matching element, return -1.
func Find[T comparable](list GoList[T], value T) int {
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if node.Data == value {
            return i
        }
        i++
    }
    return -1
}

// Execute fun with input is elements in list from left to right and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldl[T1, T2 comparable](list GoList[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    for node := list.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Execute fun with input is elements in list from right to left and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldr[T1, T2 comparable](list GoList[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Applying function for each element in list.
func ForEach[T comparable](list GoList[T], fun func(T)) {
    for node := list.Head; node != nil; node = node.Next {
        fun(node.Data)
    }
}

// Return a list with val is inserted at specific index. Note that index is
// capped at list length. Negative index indicate an offset from the end of list.
func InsertAt[T comparable](list GoList[T], index int, val T) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index > len {
        panic("InsertAt, index is out of bound!")
    }

    var result GoList[T]
    if index == len {
        for node := list.Head; node != nil; node = node.Next {
            result.doAppendHead(node.Data)
        }
        result.doAppendHead(val)
    } else {
        i := 0
        for node := list.Head; node != nil; node = node.Next {
            if i == index {
                result.doAppendHead(val)
            }
            result.doAppendHead(node.Data)
            i++
        }
    }

    return *result.Reverse()
}

// Insert sep between each element in list and return as new list.
func Join[T comparable](list GoList[T], sep T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(node.Data)
        if node.Next != nil {
            result.doAppendHead(sep)
        }
    }
    return *result.Reverse()
}

// Return data of last node in list.
func Last[T comparable](list GoList[T]) T {
    if list.Head == nil || list.Head.Next == nil {
        return list.Head.Data
    }
    node := list.Head
    for node.Next != nil {
        node = node.Next
    }
    return node.Data
}

// Return length of list.
func Len[T comparable](list GoList[T]) int {
    len := 0
    for node := list.Head; node != nil; node = node.Next {
        len += 1
    }
    return len
}

// Applying function to every elements in list and return as new list.
func Map[T comparable](list GoList[T], fun func(T) T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.doAppendHead(fun(node.Data))
    }
    return *result.Reverse()
}

// Combines the operations of Map function and Foldl function into one pass.
func MapFoldl[T1, T2 comparable](list GoList[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList[T1], T2) {
    var value T1
    var result GoList[T1]
    for node := list.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.doAppendHead(value)
    }
    return *result.Reverse(), acc0
}

// Combines the operations of Map function and Foldr function into one pass.
func MapFoldr[T1, T2 comparable](list GoList[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList[T1], T2) {
    var value T1
    var result GoList[T1]
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.doAppendHead(value)
    }
    return *result.Reverse(), acc0
}

// Return maximum element in list.
// This function only works with list of numbers or strings.
func Max[T constraints.Ordered](list GoList[T]) T {
    node := list.Head
    max := node.Data
    for node != nil {
        if node.Data > max {
            max = node.Data
        }
        node = node.Next
    }
    return max
}

// Return true if elem in list, otherwise return false.
func Member[T comparable](list GoList[T], elem T) bool {
    for node := list.Head; node != nil; node = node.Next {
        if node.Data == elem {
            return true
        }
    }
    return false
}

// Return a sorted list forming by merging list1 and list2.
func Merge[T constraints.Ordered](list1, list2 GoList[T]) GoList[T] {
    result := Concat(list1, list2)
    return Sort(result)
}

// Return minimum element in list.
// This functions only works with list of numbers or strings.
func Min[T constraints.Ordered](list GoList[T]) T {
    node := list.Head
    min := node.Data
    for node != nil {
        if node.Data < min {
            min = node.Data
        }
        node = node.Next
    }
    return min
}

// Return data of node in list at specific index. Note that index is capped at
// list length. Negative index indicate an offset from the end of list.
func Nth[T comparable](list GoList[T], index int) T {
    len := Len(list)

    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index < 0 || index >= len {
        panic("Nth, index is out of bound!")
    }

    node := list.Head
    for i := 0; i < index; i++ {
        node = node.Next
    }
    return node.Data
}

// Return sublist from node in list at specific index as a new list. Note that
// index is capped at list length. Negative index indicate an offset from
// the end of list.
func NthTail[T comparable](list GoList[T], index int) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index >= len {
        panic("NthTail, index is out of bound!")
    }

    var result GoList[T]
    node := list.Head
    for i := 0; i < index; i++ {
        node = node.Next
    }
    for node != nil {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    return *result.Reverse()
}

// Split list into list1 and list2, where list1 contains elements which fun
// returns true and list2 contains elements which fun returns false.
func Partition[T comparable](list GoList[T], fun func(T) bool) (GoList[T], GoList[T]) {
    var list1 GoList[T]
    var list2 GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list1.doAppendHead(node.Data)
        } else {
            list2.doAppendHead(node.Data)
        }
    }
    return *list1.Reverse(), *list2.Reverse()
}

// Return true if list1 is prefix of list2, otherwise return false.
func Prefix[T comparable](list1, list2 GoList[T]) bool {
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

// Return a list that node at specific index is replaced with val. If index is
// out of bound, the original list is returned. Negative index indicate an
// offset from the end of list.
func ReplaceAt[T comparable](list GoList[T], index int, val T) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            result.doAppendHead(val)
        } else {
            result.doAppendHead(node.Data)
        }
        i++
    }

    return *result.Reverse()
}

// Return result is reverse of input list
func Reverse[T comparable](list GoList[T]) GoList[T] {
    var head *node.Node[T]
    for curr := list.Head; curr != nil; curr = curr.Next {
        node := &node.Node[T]{Data: curr.Data, Next: head}
        head = node
    }
    return GoList[T]{Head: head}
}

// Return position and value of first element in list that fun returns true. If
// every fun execution returns false, this function will returns -1 and zero
// value of T.
func Search[T comparable](list GoList[T], fun func(T) bool) (int, T) {
    var zero T
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return i, node.Data
        }
        i++
    }
    return -1, zero
}

// Return sequence of numbers that starts with from and contains the successive
// result of adding incr to the previous element, until to is reached or passed.
func Seq[T constraints.Numeric](from, to, incr T) GoList[T] {
    var result GoList[T]
    for i := from; i <= to; i += incr {
        result.doAppendHead(i)
    }
    return *result.Reverse()
}

// Sort input list and returns as new list.
func Sort[T constraints.Ordered](list GoList[T]) GoList[T] {
    return quickSort(list)
}

// Split list into list1 and list2, list1 contains n first elements and list2
// contains the remaining elements. Note that n is capped at list length.
// Negative n indicate an offset from the end of list.
func Split[T comparable](list GoList[T], n int) (GoList[T], GoList[T]) {
    len := Len(list)
    if n < 0 {
        n = len + n // same as len - abs(n)
    }
    if n < 0 || n >= len {
        panic("Split, n is out of bound!")
    }

    var list1 GoList[T]
    var list2 GoList[T]
    node := list.Head
    for i := 0; i < n; i++ {
        list1.doAppendHead(node.Data)
        node = node.Next
    }
    for node != nil {
        list2.doAppendHead(node.Data)
        node = node.Next
    }

    return *list1.Reverse(), *list2.Reverse()
}

// Split input list into list1 and list2, where list1 behave as
// TakeWhile(fun, list) and list2 behave as DropWhile(fun, list).
func SplitWith[T comparable](list GoList[T], fun func(T) bool) (GoList[T], GoList[T]) {
    var list1 GoList[T]
    var list2 GoList[T]
    node := list.Head
    for node != nil {
        if fun(node.Data) {
            list1.doAppendHead(node.Data)
        } else {
            break
        }
        node = node.Next
    }
    for node != nil {
        list2.doAppendHead(node.Data)
        node = node.Next
    }
    return *list1.Reverse(), *list2.Reverse()
}

// Return sublist of input list as new list, starting at start and has maximum
// len elements. Note that start is capped at list length. Negative start
// indicate an offset from the end of list.
func Sublist[T comparable](list GoList[T], start, len int) GoList[T] {
    if len < 0 {
        panic("Sublist, input len must not be negative!")
    }

    listLen := Len(list)
    if start < 0 {
        start = listLen + start // same as len - abs(start)
    }
    if start < 0 || start >= listLen {
        panic("Sublist, start is out of bound!")
    }

    var result GoList[T]
    node := list.Head
    for i := 0; node != nil && i < start; i++ {
        node = node.Next
    }
    for j := 0; node != nil && j < len; j++ {
        result.doAppendHead(node.Data)
        node = node.Next
    }
    return *result.Reverse()
}

// Return a new list that is a copy of list1 which is for each element in list2,
// its first occurrence in list1 is deleted.
func Subtract[T comparable](list1, list2 GoList[T]) GoList[T] {
    var result GoList[T]
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        result.doAppendHead(node1.Data)
    }
    result.Reverse()
    for node2 := list2.Head; node2 != nil; node2 = node2.Next {
        result.Delete(node2.Data)
    }
    return result
}

// Returns true if list1 is a suffix of list2, otherwise false.
func Suffix[T comparable](list1, list2 GoList[T]) bool {
    reverse1 := Reverse(list1)
    reverse2 := Reverse(list2)
    return Prefix(reverse1, reverse2)
}

// Returns sum of elements in list.
func Sum[T constraints.Ordered](list GoList[T]) T {
    var sum T
    for node := list.Head; node != nil; node = node.Next {
        sum += node.Data
    }
    return sum
}

// Take elements in list while fun returns true, and return as new list.
func TakeWhile[T comparable](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.doAppendHead(node.Data)
        } else {
            break
        }
    }
    return *result.Reverse()
}

// Return a unique sorted list forming by merging list1 and list2, then remove
// all duplicated elements.
func UMerge[T constraints.Ordered](list1, list2 GoList[T]) GoList[T] {
    result := Concat(list1, list2)
    return uniqueQuickSort(result)
}

// Sort input list and remove all duplicated elements, then returns as new list.
func USort[T constraints.Ordered](list GoList[T]) GoList[T] {
    return uniqueQuickSort(list)
}

// Return a list that node at specific index is updated with return value of
// fun. If index is out of bound, the original list is returned. Negative index
// indicate an offset from the end of list.
func UpdateAt[T comparable](list GoList[T], index int, fun func(T) T) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            result.doAppendHead(fun(node.Data))
        } else {
            result.doAppendHead(node.Data)
        }
        i++
    }

    return *result.Reverse()
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

// Append all values into last of list.
func (list *GoList[T]) Append(values ...T) *GoList[T] {
    list.Reverse()
    for _, value := range values {
        list.doAppendHead(value)
    }
    list.Reverse()
    return list
}

// Append all values into head of list.
func (list *GoList[T]) AppendHead(values ...T) *GoList[T] {
    for i := len(values) - 1; i >= 0; i-- {
        list.doAppendHead(values[i])
    }
    return list
}

// Concatenates list2 into last of list1. This method won't remove list2, so
// after this method executed, changes made to list2 might affect list1 and
// changes made to list1 might affect list2 as well.
func (list1 *GoList[T]) Concat(list2 GoList[T]) *GoList[T] {
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        if node1.Next == nil {
            node1.Next = list2.Head
            return list1
        }
    }
    return list1
}

// Delete first node in list with value of input.
func (list *GoList[T]) Delete(value T) *GoList[T] {
    if list.Head == nil {
        return list
    }
    // Ignore if it's first node
    if list.Head.Data == value {
        list.Head = list.Head.Next
        return list
    }
    node := list.Head
    for node.Next != nil {
        if node.Next.Data == value {
            node.Next = node.Next.Next
            return list
        }
        node = node.Next
    }
    return list
}

// Delete node at specific index in list. Note that index is capped at list
// length. Negative index indicate offset from the end of list.
func (list *GoList[T]) DeleteAt(index int) *GoList[T] {
    len := Len(*list)

    if list.Head == nil {
        return list
    }

    if index == 0 {
        list.Head = list.Head.Next
        return list
    } else if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index > 0 && index < len {
        i := 1
        for node := list.Head; node.Next != nil; node = node.Next {
            if i == index {
                node.Next = node.Next.Next
                return list
            }
            i++
        }
    }

    return list
}

// Drop the last element in list.
func (list *GoList[T]) DropLast() *GoList[T] {
    node := list.Head
    for node != nil {
        if node.Next.Next == nil {
            node.Next = nil
            break
        }
        node = node.Next
    }
    return list
}

// Drop elements in list while fun returns true.
func (list *GoList[T]) DropWhile(fun func(T) bool) *GoList[T] {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list.Head = node.Next
        } else {
            break
        }
    }
    return list
}

// Only keep elements in list that fun return true.
func (list *GoList[T]) Filter(fun func(T) bool) *GoList[T] {
    if list.Head == nil {
        return list
    }
    if !fun(list.Head.Data) {
        list.Head = list.Head.Next
    }
    for node := list.Head; node.Next != nil; node = node.Next {
        if !fun(node.Next.Data) {
            node.Next = node.Next.Next
        }
    }
    return list
}

// Calls fun on successive elements of list. fun must return (bool, value).
// The function returns the list of elements for which fun returns
// a new value, where a value of true is synonymous with (true, value).
func (list *GoList[T]) FilterMap(fun func(T) (bool, T)) *GoList[T] {
    dummy := &node.Node[T]{Next: list.Head}
    prev := dummy

    for node := list.Head; node != nil; node = node.Next {
        if keep, value := fun(node.Data); keep {
            node.Data = value
            prev = node
        } else {
            prev.Next = node.Next
        }
    }

    list.Head = dummy.Next
    return list
}

// Insert val into list at specific index. Note that index is capped at list
// length. Negative index indicate an offset from the end of list.
func (list *GoList[T]) InsertAt(index int, val T) *GoList[T] {
    len := Len(*list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index > len {
        panic("InsertAt, index is out of bound!")
    }

    insertNode := &node.Node[T]{Data: val}
    if index == 0 {
        insertNode.Next = list.Head
        list.Head = insertNode
    } else {
        i := 1
        for node := list.Head; node != nil; node = node.Next {
            if i == index {
                insertNode.Next = node.Next
                node.Next = insertNode
                break
            }
            i++
        }
    }

    return list
}

// Insert sep between each element in list.
func (list *GoList[T]) Join(sep T) *GoList[T] {
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
    return list
}

// Applying function to every elements in list.
func (list *GoList[T]) Map(fun func(T) T) *GoList[T] {
    for node := list.Head; node != nil; node = node.Next {
        node.Data = fun(node.Data)
    }
    return list
}

// Return sublist from node in list at specific index. Note that index is capped
// at list length. Negative index indicate an offset from the end of list.
func (list *GoList[T]) NthTail(index int) *GoList[T] {
    len := Len(*list)
    if index < 0 {
        index = len + index // same as len - abs(n)
    }
    if index < 0 || index >= len {
        panic("NthTail, index is out of bound!")
    }

    node := list.Head
    for i := 0; i < index; i++ {
        node = node.Next
    }
    list.Head = node
    return list
}

// Replace a node at specific index in list with val. If index is out of bound,
// the original list is returned. Negative index indicate an offset from the
// end of list.
func (list *GoList[T]) ReplaceAt(index int, val T) *GoList[T] {
    len := Len(*list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            node.Data = val
        }
        i++
    }

    return list
}

// Reverse the input list
func (list *GoList[T]) Reverse() *GoList[T] {
    var prev *node.Node[T]
    node := list.Head
    for node != nil {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
    }
    list.Head = prev
    return list
}

// Return sublist that starting at start and has maximum len elements. Note
// that start is capped at list length. Negative start indicate an offset from
// the end of list.
func (list *GoList[T]) Sublist(start, len int) *GoList[T] {
    if len == 0 {
        list.Head = nil
        return list
    } else if len < 0 {
        panic("Sublist, input len must not be negative!")
    }

    listLen := Len(*list)
    if start < 0 {
        start = listLen + start // same as len - abs(start)
    }
    if start < 0 || start >= listLen {
        panic("Sublist, start is out of bound!")
    }

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

    return list
}

// Delete elements in list1 that is its first occurrence to each element in list2.
func (list1 *GoList[T]) Subtract(list2 GoList[T]) *GoList[T] {
    for node := list2.Head; node != nil; node = node.Next {
        list1.Delete(node.Data)
    }
    return list1
}

// Take elements in list while fun returns true.
func (list *GoList[T]) TakeWhile(fun func(T) bool) *GoList[T] {
    if !fun(list.Head.Data) {
        list.Head = nil
        return list
    }
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Next.Data) {
            node.Next = nil
        }
    }
    return list
}

// Update a node at specific index in list with return value of fun. If index is
// out of bound, the original list is returned. Negative index indicate an
// offset from the end of list.
func (list *GoList[T]) UpdateAt(index int, fun func(T) T) *GoList[T] {
    len := Len(*list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            node.Data = fun(node.Data)
        }
        i++
    }

    return list
}

// Return a string representing singly linked list.
func (list GoList[T]) String() string {
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

// Do append value into head of list.
func (list *GoList[T]) doAppendHead(value T) *GoList[T] {
    node := &node.Node[T]{Data: value, Next: list.Head}
    list.Head = node
    return list
}

// Do quick sort input list and returns as new list.
func quickSort[T constraints.Ordered](list GoList[T]) GoList[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoList[T]

    // Partitioning
    for node := list.Head; node != nil; node = node.Next {
        switch {
        case node.Data < pivot:
            less.doAppendHead(node.Data)
        case node.Data == pivot:
            equal.doAppendHead(node.Data)
        case node.Data > pivot:
            greater.doAppendHead(node.Data)
        }
    }

    // Recursive sort
    sortedLess := quickSort(less)
    sortedGreater := quickSort(greater)

    // Concatenates 3 lists: sortedLess + equal + sortedGreater
    result := Concat(sortedLess, equal)
    result = Concat(result, sortedGreater)

    return result
}

// Do quick sort input list and remove duplicate nodes and returns as new list.
func uniqueQuickSort[T constraints.Ordered](list GoList[T]) GoList[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoList[T]
    seen := make(map[T]bool) // store already seen node data into map

    // Partitioning and remove seen node
    for node := list.Head; node != nil; node = node.Next {
        if seen[node.Data] {
            continue
        }
        seen[node.Data] = true

        switch {
        case node.Data < pivot:
            less.doAppendHead(node.Data)
        case node.Data == pivot:
            equal.doAppendHead(node.Data)
        case node.Data > pivot:
            greater.doAppendHead(node.Data)
        }
    }

    // Recursive sort
    sortedLess := uniqueQuickSort(less)
    sortedGreater := uniqueQuickSort(greater)

    // Concatenate: sortedLess + equal + sortedGreater
    result := Concat(sortedLess, equal)
    result = Concat(result, sortedGreater)

    return result
}
