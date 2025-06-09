// Package golist contains functions and methods for singly linked list in Go.
package golist

import (
    "fmt"
    "strings"
    "github.com/google/go-cmp/cmp"
    "github.com/hiennguyen-neih/go-linkedlist/node"
    "github.com/hiennguyen-neih/go-linkedlist/constraints"
)

/*
 *******************************************************************************
 * Define structs and interfaces
 *******************************************************************************
 */

// Struct of Go singly linked list.
type GoList[T any] struct {
    Head *node.Node[T]
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly linked list.
func New[T any](values ...T) GoList[T] {
    var list GoList[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input slice into linked list.
func FromSlice[T any](values []T) GoList[T] {
    var list GoList[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input linked list into slice.
func ToSlice[T any](list GoList[T]) []T {
    var result []T
    for node := list.Head; node != nil; node = node.Next {
        result = append(result, node.Data)
    }
    return result
}

// Return true if fun returns true for all elements in list, otherwise return
// false.
func All[T any](list GoList[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Data) {
            return false
        }
    }
    return true
}

// Return true if fun returns true for at least 1 element in list, otherwise
// return false.
func Any[T any](list GoList[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return true
        }
    }
    return false
}

// Return new list is append of input list and values.
func Append[T any](list GoList[T], values ...T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
    }
    for _, value := range values {
        result.appendHead(value)
    }
    return *result.reverse()
}

// Return new list is input list that append values into head of it.
func AppendHead[T any](list GoList[T], values ...T) GoList[T] {
    var result GoList[T]
    for _, value := range values {
        result.appendHead(value)
    }
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
    }
    return *result.reverse()
}

// Return a list that is concatenated of all input lists.
func Concat[T any](lists ...GoList[T]) GoList[T] {
    var result GoList[T]
    for _, list := range lists {
        for node := list.Head; node != nil; node = node.Next {
            result.appendHead(node.Data)
        }
    }
    return *result.reverse()
}

// Delte first node in list with value of input and return as new list.
func Delete[T any](list GoList[T], value T) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
    }
    node := list.Head
    for node != nil {
        if !cmp.Equal(node.Data, value) {
            result.appendHead(node.Data)
            node = node.Next
        } else {
            node = node.Next
            break
        }
    }
    for node != nil {
        result.appendHead(node.Data)
        node = node.Next
    }
    return *result.reverse()
}

// Delete node at the specific index in list and return as new list. If index
// is out of bound, the original list is returned. Negative index indicate an
// offset from the end of list.
func DeleteAt[T any](list GoList[T], index int) GoList[T] {
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
            result.appendHead(node.Data)
        }
    } else {
        i := 0
        for node := list.Head; node != nil; node = node.Next {
            if i == index {
                i++
                continue
            }
            i++
            result.appendHead(node.Data)
        }
    }

    return *result.reverse()
}

// Drop the last element in list and return as new list.
func DropLast[T any](list GoList[T]) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if node.Next.Next == nil {
            result.appendHead(node.Data)
            break
        }
        result.appendHead(node.Data)
    }
    return *result.reverse()
}

// Drop elements in list while fun returns true, and return as new list.
func DropWhile[T any](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
    node := list.Head
    for node != nil {
        if !fun(node.Data) {
            break
        }
        node = node.Next
    }
    for node != nil {
        result.appendHead(node.Data)
        node = node.Next
    }
    return *result.reverse()
}

// Return a list containing n copies of term element. If n is negative or equal
// 0, return empty list.
func Duplicate[T any](n int, elem T) GoList[T] {
    var result GoList[T]
    for i := 0; i < n; i++ {
        result.appendHead(elem)
    }
    return result
}

// Return true if all corresponding nodes in both list1 and list2 have the same
// value, otherwise return false.
func Equal[T any](list1, list2 GoList[T]) bool {
    node2 := list2.Head
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        if node2 == nil || !cmp.Equal(node1.Data, node2.Data) {
            return false
        }
        node2 = node2.Next
    }
    if node2 != nil {
        return false
    }
    return true
}

// Return new list contains elements in input that fun returns true.
func Filter[T any](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.appendHead(node.Data)
        }
    }
    return *result.reverse()
}

// Calls fun on successive elements of list. fun must return (bool, value).
// The function returns a new list of elements for which fun returns
// a new value, where a value of true is synonymous with (true, value).
func FilterMap[T any](list GoList[T], fun func(T) (bool, T)) GoList[T] {
    var result GoList[T]
    if list.Head == nil {
        return result
    }
    for node := list.Head; node != nil; node = node.Next {
        if keep, value := fun(node.Data); keep {
            result.appendHead(value)
        }
    }
    return *result.reverse()
}

// Return position of first element in list that match with value. If there is
// no matching element, return -1.
func Find[T any](list GoList[T], value T) int {
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if cmp.Equal(node.Data, value) {
            return i
        }
        i++
    }
    return -1
}

// Execute fun with input is elements in list from left to right and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldl[T1, T2 any](list GoList[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    for node := list.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Execute fun with input is elements in list from right to left and acc0,
// fun return new acc and it's used as input for next execution.
// Return the acc of the last execution.
func Foldr[T1, T2 any](list GoList[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Applying function for each element in list.
func ForEach[T any](list GoList[T], fun func(T)) {
    for node := list.Head; node != nil; node = node.Next {
        fun(node.Data)
    }
}

// Return a list with val is inserted at specific index. Note that index is
// capped at list length. Negative index indicate an offset from the end of list.
func InsertAt[T any](list GoList[T], index int, val T) GoList[T] {
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
            result.appendHead(node.Data)
        }
        result.appendHead(val)
    } else {
        i := 0
        for node := list.Head; node != nil; node = node.Next {
            if i == index {
                result.appendHead(val)
            }
            result.appendHead(node.Data)
            i++
        }
    }

    return *result.reverse()
}

// Insert sep between each element in list and return as new list.
func Join[T any](list GoList[T], sep T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
        if node.Next != nil {
            result.appendHead(sep)
        }
    }
    return *result.reverse()
}

// Return last node in list.
func Last[T any](list GoList[T]) *node.Node[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list.Head
    }
    node := list.Head
    for node.Next != nil {
        node = node.Next
    }
    return node
}

// Return length of list.
func Len[T any](list GoList[T]) int {
    len := 0
    for node := list.Head; node != nil; node = node.Next {
        len += 1
    }
    return len
}

// Applying function to every elements in list and return as new list.
func Map[T any](list GoList[T], fun func(T) T) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(fun(node.Data))
    }
    return *result.reverse()
}

// Combines the operations of Map function and Foldl function into one pass.
func MapFoldl[T1, T2 any](list GoList[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList[T1], T2) {
    var value T1
    var result GoList[T1]
    for node := list.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.appendHead(value)
    }
    return *result.reverse(), acc0
}

// Combines the operations of Map function and Foldr function into one pass.
func MapFoldr[T1, T2 any](list GoList[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList[T1], T2) {
    var value T1
    var result GoList[T1]
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.appendHead(value)
    }
    return result, acc0
}

// Return node with maximum value in list. This function only works with
// constraint Ordered list.
func Max[T constraints.Ordered](list GoList[T]) *node.Node[T] {
    node := list.Head
    max := node
    for node != nil {
        if node.Data > max.Data {
            max = node
        }
        node = node.Next
    }
    return max
}

// Return true if elem in list, otherwise return false.
func Member[T any](list GoList[T], elem T) bool {
    for node := list.Head; node != nil; node = node.Next {
        if cmp.Equal(node.Data, elem) {
            return true
        }
    }
    return false
}

// Return a sorted list forming by merging all input lists.
func Merge[T constraints.Ordered](lists ...GoList[T]) GoList[T] {
    result := Concat(lists...)
    return Sort(result)
}

// Return node with minimum value in list. This function only works with
// constraint Ordered list.
func Min[T constraints.Ordered](list GoList[T]) *node.Node[T] {
    node := list.Head
    min := node
    for node != nil {
        if node.Data < min.Data {
            min = node
        }
        node = node.Next
    }
    return min
}

// Return node in list at specific index. Note that index is capped at list
// length. Negative index indicate an offset from the end of list.
func Nth[T any](list GoList[T], index int) *node.Node[T] {
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
    return node
}

// Return sublist from node in list at specific index as a new list. Note that
// index is capped at list length. Negative index indicate an offset from
// the end of list.
func NthTail[T any](list GoList[T], index int) GoList[T] {
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
        result.appendHead(node.Data)
        node = node.Next
    }
    return *result.reverse()
}

// Split list into list1 and list2, where list1 contains elements which fun
// returns true and list2 contains elements which fun returns false.
func Partition[T any](list GoList[T], fun func(T) bool) (GoList[T], GoList[T]) {
    var list1 GoList[T]
    var list2 GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list1.appendHead(node.Data)
        } else {
            list2.appendHead(node.Data)
        }
    }
    return *list1.reverse(), *list2.reverse()
}

// Return true if list1 is prefix of list2, otherwise return false.
func Prefix[T any](list1, list2 GoList[T]) bool {
    node1 := list1.Head
    node2 := list2.Head
    for node1 != nil {
        if node2 == nil || !cmp.Equal(node1.Data, node2.Data) {
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
func ReplaceAt[T any](list GoList[T], index int, val T) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            result.appendHead(val)
        } else {
            result.appendHead(node.Data)
        }
        i++
    }

    return *result.reverse()
}

// Return result is reverse of input list
func Reverse[T any](list GoList[T]) GoList[T] {
    var head *node.Node[T]
    for curr := list.Head; curr != nil; curr = curr.Next {
        node := &node.Node[T]{Data: curr.Data, Next: head}
        head = node
    }
    return GoList[T]{Head: head}
}

// Return position and first node in list that fun returns true. If every fun
// execution returns false, this function will returns position is -1.
func Search[T any](list GoList[T], fun func(T) bool) (int, *node.Node[T]) {
    var zero *node.Node[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return i, node
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
        result.appendHead(i)
    }
    return *result.reverse()
}

// Sort input list and returns as new list.
func Sort[T constraints.Ordered](list GoList[T]) GoList[T] {
    return quickSort(list)
}

// Split list into list1 and list2, list1 contains n first elements and list2
// contains the remaining elements. Note that n is capped at list length.
// Negative n indicate an offset from the end of list.
func Split[T any](list GoList[T], n int) (GoList[T], GoList[T]) {
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
        list1.appendHead(node.Data)
        node = node.Next
    }
    for node != nil {
        list2.appendHead(node.Data)
        node = node.Next
    }

    return *list1.reverse(), *list2.reverse()
}

// Split input list into list1 and list2, where list1 behave as
// TakeWhile(fun, list) and list2 behave as DropWhile(fun, list).
func SplitWith[T any](list GoList[T], fun func(T) bool) (GoList[T], GoList[T]) {
    var list1 GoList[T]
    var list2 GoList[T]
    node := list.Head
    for node != nil {
        if fun(node.Data) {
            list1.appendHead(node.Data)
        } else {
            break
        }
        node = node.Next
    }
    for node != nil {
        list2.appendHead(node.Data)
        node = node.Next
    }
    return *list1.reverse(), *list2.reverse()
}

// Return sublist of input list as new list, starting at start and has maximum
// len elements. Note that start is capped at list length. Negative start
// indicate an offset from the end of list.
func Sublist[T any](list GoList[T], start, len int) GoList[T] {
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
        result.appendHead(node.Data)
        node = node.Next
    }
    return *result.reverse()
}

// Return a new list that is a copy of list1 which is for each element in list2,
// its first occurrence in list1 is deleted.
func Subtract[T any](list1, list2 GoList[T]) GoList[T] {
    var result GoList[T]
    if list1.Head == nil {
        return result
    }
    for node1 := list1.Head; node1 != nil; node1 = node1.Next {
        result.appendHead(node1.Data)
    }
    result.reverse()
    for node2 := list2.Head; node2 != nil; node2 = node2.Next {
        if cmp.Equal(result.Head.Data, node2.Data) {
            result.Head = result.Head.Next
            continue
        }
        for node3 := result.Head; node3.Next != nil; node3 = node3.Next {
            if cmp.Equal(node3.Next.Data, node2.Data) {
                node3.Next = node3.Next.Next
                break
            }
        }
    }
    return result
}

// Returns true if list1 is a suffix of list2, otherwise false.
func Suffix[T any](list1, list2 GoList[T]) bool {
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
func TakeWhile[T any](list GoList[T], fun func(T) bool) GoList[T] {
    var result GoList[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.appendHead(node.Data)
        } else {
            break
        }
    }
    return *result.reverse()
}

// Return a unique sorted list forming by merging all input lists and removing
// all duplicated elements.
func UMerge[T constraints.Ordered](lists ...GoList[T]) GoList[T] {
    result := Concat(lists...)
    return uniqueQuickSort(result)
}

// Sort input list and remove all duplicated elements, then returns as new list.
func USort[T constraints.Ordered](list GoList[T]) GoList[T] {
    return uniqueQuickSort(list)
}

// Return a list that node at specific index is updated with return value of
// fun. If index is out of bound, the original list is returned. Negative index
// indicate an offset from the end of list.
func UpdateAt[T any](list GoList[T], index int, fun func(T) T) GoList[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if i == index {
            result.appendHead(fun(node.Data))
        } else {
            result.appendHead(node.Data)
        }
        i++
    }

    return *result.reverse()
}

/*
 *******************************************************************************
 * Exported methods
 *******************************************************************************
 */

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
func (list *GoList[T]) appendHead(value T) *GoList[T] {
    node := &node.Node[T]{Data: value, Next: list.Head}
    list.Head = node
    return list
}

// Do reverse the list.
func (list *GoList[T]) reverse() *GoList[T] {
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
            less.appendHead(node.Data)
        case node.Data == pivot:
            equal.appendHead(node.Data)
        case node.Data > pivot:
            greater.appendHead(node.Data)
        }
    }

    // Recursive sort
    sortedLess := quickSort(less)
    sortedGreater := quickSort(greater)

    // Concatenates 3 lists: sortedLess + equal + sortedGreater
    result := Concat(sortedLess, equal, sortedGreater)

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
            less.appendHead(node.Data)
        case node.Data == pivot:
            equal.appendHead(node.Data)
        case node.Data > pivot:
            greater.appendHead(node.Data)
        }
    }

    // Recursive sort
    sortedLess := uniqueQuickSort(less)
    sortedGreater := uniqueQuickSort(greater)

    // Concatenate: sortedLess + equal + sortedGreater
    result := Concat(sortedLess, equal, sortedGreater)

    return result
}
