// Package golist2 contains functions and methods for doubly linked list in Go.
package golist2

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

// Struct of Go doubly linked list.
type GoList2[T any] struct {
    Head *node.Node2[T]    // First node of the list.
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new doubly linked list from input values.
func New[T any](values ...T) GoList2[T] {
    var list GoList2[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input slice into new doubly linked list.
func FromSlice[T any](values []T) GoList2[T] {
    var list GoList2[T]
    for _, val := range values {
        list.appendHead(val)
    }
    return *list.reverse()
}

// Convert input doubly linked list into new slice.
func ToSlice[T any](list GoList2[T]) []T {
    var result []T
    for node := list.Head; node != nil; node = node.Next {
        result = append(result, node.Data)
    }
    return result
}

// Returns true if fun returns true for all node data in list, otherwise
// returns false.
func All[T any](list GoList2[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if !fun(node.Data) {
            return false
        }
    }
    return true
}

// Returns true if fun returns true for at least 1 node data in list,
// otherwise returns false.
func Any[T any](list GoList2[T], fun func(T) bool) bool {
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return true
        }
    }
    return false
}

// Appends values into last of input list.
func Append[T any](list GoList2[T], values ...T) GoList2[T] {
    var result GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
    }
    for _, value := range values {
        result.appendHead(value)
    }
    return *result.reverse()
}

// Appends values into head of input list.
func AppendHead[T any](list GoList2[T], values ...T) GoList2[T] {
    var result GoList2[T]
    for _, value := range values {
        result.appendHead(value)
    }
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
    }
    return *result.reverse()
}

// Returns a list that is concatenated of all input lists.
func Concat[T any](lists ...GoList2[T]) GoList2[T] {
    var result GoList2[T]
    for _, list := range lists {
        for node := list.Head; node != nil; node = node.Next {
            result.appendHead(node.Data)
        }
    }
    return *result.reverse()
}

// Returns a copy of input list where the first node data that matching value
// is removed.
func Delete[T any](list GoList2[T], value T) GoList2[T] {
    var result GoList2[T]
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

// Deletes node at the specific index of list. If index is out of bound, the
// original list is returned. Negative index indicate an offset from the end
// of list.
func DeleteAt[T any](list GoList2[T], index int) GoList2[T] {
    var result GoList2[T]
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

// Drops the last node of input list. If input list is an empty list, returns
// an empty list.
func DropLast[T any](list GoList2[T]) GoList2[T] {
    var result GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        if node.Next.Next == nil {
            result.appendHead(node.Data)
            break
        }
        result.appendHead(node.Data)
    }
    return *result.reverse()
}

// Drops nodes from list while fun returns true.
func DropWhile[T any](list GoList2[T], fun func(T) bool) GoList2[T] {
    var result GoList2[T]
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

// Returns a list containing n copies of term elem. If n is negative or equal
// 0, return empty list.
func Duplicate[T any](n int, elem T) GoList2[T] {
    var result GoList2[T]
    for i := 0; i < n; i++ {
        result.appendHead(elem)
    }
    return result
}

// Returns true if all corresponding nodes in both list1 and list2 have the
// same value, otherwise return false.
func Equal[T any](list1, list2 GoList2[T]) bool {
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

// Returns a list contains node data from input list for which fun returns true.
func Filter[T any](list GoList2[T], fun func(T) bool) GoList2[T] {
    var result GoList2[T]
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

// Calls fun on successive nodes of list to update or remove nodes from list.
// Input fun must return (bool, value). The functions returns a list that nodes
// data are value in which fun returns (true, value).
func FilterMap[T any](list GoList2[T], fun func(T) (bool, T)) GoList2[T] {
    var result GoList2[T]
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

// Returns position of first node of list that match with value. If there is
// no matching node, returns -1.
func Find[T any](list GoList2[T], value T) int {
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if cmp.Equal(node.Data, value) {
            return i
        }
        i++
    }
    return -1
}

// Calls fun(data, acc) on successive nodes of list from left to right (from
// start of list to end of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldl[T1, T2 any](list GoList2[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    for node := list.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Calls fun(data, acc) on successive nodes of list from right to left (from
// end of list to start of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldr[T1, T2 any](list GoList2[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        acc0 = fun(node.Data, acc0)
    }
    return acc0
}

// Calls fun(data) for each node in list, ignoring the return value. This
// function is used for its side effects and the evaluation order is defined
// to be the same as the order of the nodes in the list.
func ForEach[T any](list GoList2[T], fun func(T)) {
    for node := list.Head; node != nil; node = node.Next {
        fun(node.Data)
    }
}

// Returns a list with val is inserted at specific index. index is capped at
// list length. Negative index indicate an offset from the end of list.
func InsertAt[T any](list GoList2[T], index int, val T) GoList2[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index > len {
        panic("InsertAt, index is out of bound!")
    }

    var result GoList2[T]
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

// Inserts sep between each node in list. This function has no effect on an
// empty list or a singleton list.
func Join[T any](list GoList2[T], sep T) GoList2[T] {
    var result GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(node.Data)
        if node.Next != nil {
            result.appendHead(sep)
        }
    }
    return *result.reverse()
}

// Returns the last node in list.
func Last[T any](list GoList2[T]) *node.Node2[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list.Head
    }
    node := list.Head
    for node.Next != nil {
        node = node.Next
    }
    return node
}

// Returns the length of list.
func Len[T any](list GoList2[T]) int {
    len := 0
    for node := list.Head; node != nil; node = node.Next {
        len += 1
    }
    return len
}

// Calls fun(data) to every nodes in list and returns a list contains returned
// values of that fun.
func Map[T any](list GoList2[T], fun func(T) T) GoList2[T] {
    var result GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        result.appendHead(fun(node.Data))
    }
    return *result.reverse()
}

// Combines the operations of Map function and Foldl function into one pass.
func MapFoldl[T1, T2 any](list GoList2[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList2[T1], T2) {
    var value T1
    var result GoList2[T1]
    for node := list.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.appendHead(value)
    }
    return *result.reverse(), acc0
}

// Combines the operations of Map function and Foldr function into one pass.
func MapFoldr[T1, T2 any](list GoList2[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoList2[T1], T2) {
    var value T1
    var result GoList2[T1]
    reverse := Reverse(list)
    for node := reverse.Head; node != nil; node = node.Next {
        value, acc0 = fun(node.Data, acc0)
        result.appendHead(value)
    }
    return result, acc0
}

// Returns the first node in list that compares greater than or equal to all
// other nodes of list. This function only works with constraint Ordered list.
func Max[T constraints.Ordered](list GoList2[T]) *node.Node2[T] {
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

// Returns true if elem matches some node data of list, otherwise retusn false.
func Member[T any](list GoList2[T], elem T) bool {
    for node := list.Head; node != nil; node = node.Next {
        if cmp.Equal(node.Data, elem) {
            return true
        }
    }
    return false
}

// Returns a sorted list forming by merging all input lists. This function only
// works with constraint Ordered lists.
func Merge[T constraints.Ordered](lists ...GoList2[T]) GoList2[T] {
    result := Concat(lists...)
    return Sort(result)
}

// Returns the first node in list that compares less than or equal to all
// other nodes of list. This function only works with constraint Ordered list.
func Min[T constraints.Ordered](list GoList2[T]) *node.Node2[T] {
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

// Returns node in list at specific index. index is capped at list length.
// Negative index indicate an offset from the end of list.
func Nth[T any](list GoList2[T], index int) *node.Node2[T] {
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

// Returns sublist from node in list at specific index. index is capped at
// list length. Negative index indicate an offset from the end of list.
func NthTail[T any](list GoList2[T], index int) GoList2[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }
    if index < 0 || index >= len {
        panic("NthTail, index is out of bound!")
    }

    var result GoList2[T]
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

// Partitions input list into list1 and list2, where list1 contains nodes
// which fun returns true and list2 contains nodes which fun returns false.
func Partition[T any](list GoList2[T], fun func(T) bool) (GoList2[T], GoList2[T]) {
    var list1 GoList2[T]
    var list2 GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            list1.appendHead(node.Data)
        } else {
            list2.appendHead(node.Data)
        }
    }
    return *list1.reverse(), *list2.reverse()
}

// Returns true if list1 is a prefix of list2, otherwise returns false.
// A prefix of a list is the first part of the list, starting from the
// beginning and stopping at any point.
func Prefix[T any](list1, list2 GoList2[T]) bool {
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

// Returns a list that node at specific index is replaced with val. If index
// is out of bound, the original list is returned. Negative index indicate an
// offset from the end of list.
func ReplaceAt[T any](list GoList2[T], index int, val T) GoList2[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList2[T]
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

// Returns a list containing the nodes of input list in reverse order.
func Reverse[T any](list GoList2[T]) GoList2[T] {
    var head *node.Node2[T]
    for curr := list.Head; curr != nil; curr = curr.Next {
        node := &node.Node2[T]{Data: curr.Data, Next: head}
        head = node
    }
    return GoList2[T]{Head: head}
}

// Returns position and first node in list that fun returns true. If every fun
// execution returns false, returns position is -1.
func Search[T any](list GoList2[T], fun func(T) bool) (int, *node.Node2[T]) {
    var zero *node.Node2[T]
    i := 0
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            return i, node
        }
        i++
    }
    return -1, zero
}

// Returns sequence of numbers that starts with from and contains the
// successive results of adding incr to the previous node data, until to is
// reached or passed (in later case, to is not an node data of the sequence).
func Seq[T constraints.Numeric](from, to, incr T) GoList2[T] {
    var result GoList2[T]
    for i := from; i <= to; i += incr {
        result.appendHead(i)
    }
    return *result.reverse()
}

// Returns a list containing the sorted nodes data of input list. This function
// only works with constraint Ordered list.
func Sort[T constraints.Ordered](list GoList2[T]) GoList2[T] {
    return quickSort(list)
}

// Split input list into list1 and list2, list1 contains n first nodes and
// list2 contains the remaining nodes. n is capped at list length. Negative
// n indicate an offset from the end of list.
func Split[T any](list GoList2[T], n int) (GoList2[T], GoList2[T]) {
    len := Len(list)
    if n < 0 {
        n = len + n // same as len - abs(n)
    }
    if n < 0 || n >= len {
        panic("Split, n is out of bound!")
    }

    var list1 GoList2[T]
    var list2 GoList2[T]
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
func SplitWith[T any](list GoList2[T], fun func(T) bool) (GoList2[T], GoList2[T]) {
    var list1 GoList2[T]
    var list2 GoList2[T]
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

// Returns sublist of input list, starting at start and has maximum len nodes.
// start is capped at list length. Negative start indicate an offset from the
// end of list. len must be a non-negative integer. It is not an error for
// start + len to exceed the length of list.
func Sublist[T any](list GoList2[T], start, len int) GoList2[T] {
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

    var result GoList2[T]
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

// Returns a new list that is a copy of list1 which is for each node data in
// list2, its first occurrence in list1 is deleted.
func Subtract[T any](list1, list2 GoList2[T]) GoList2[T] {
    var result GoList2[T]
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

// Returns true if list1 is a suffix of list2, otherwise returns false.
// A suffix of a list if the last part of the list, starting from any position
// and going all the way to the end.
func Suffix[T any](list1, list2 GoList2[T]) bool {
    reverse1 := Reverse(list1)
    reverse2 := Reverse(list2)
    return Prefix(reverse1, reverse2)
}

// Returns sum of all nodes data in list. This function only works with
// constraint Ordered list.
func Sum[T constraints.Ordered](list GoList2[T]) T {
    var sum T
    for node := list.Head; node != nil; node = node.Next {
        sum += node.Data
    }
    return sum
}

// Takes nodes data in list while fun returns true, returning the longest
// prefix in which all nodes data satisfy the predicate.
func TakeWhile[T any](list GoList2[T], fun func(T) bool) GoList2[T] {
    var result GoList2[T]
    for node := list.Head; node != nil; node = node.Next {
        if fun(node.Data) {
            result.appendHead(node.Data)
        } else {
            break
        }
    }
    return *result.reverse()
}

// Returns a sorted list formed by merging all input lists, while removing
// duplicates. This function only works with constraint Ordered lists.
func UMerge[T constraints.Ordered](lists ...GoList2[T]) GoList2[T] {
    result := Concat(lists...)
    return uniqueQuickSort(result)
}

// Returns a sorted list of the nodes data of list, keeping only the first
// occurrence of nodes that compare equal and removing duplicates. This
// function only works with constraint Ordered list.
func USort[T constraints.Ordered](list GoList2[T]) GoList2[T] {
    return uniqueQuickSort(list)
}

// Returns a list that node at specific index is updated with returns value of
// fun. If index is out of bound, the original list is returned. Negative index
// indicate an offset from the end of list.
func UpdateAt[T any](list GoList2[T], index int, fun func(T) T) GoList2[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    var result GoList2[T]
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

// Returns a string representing the doubly linked list.
func (list GoList2[T]) String() string {
    var builder strings.Builder
    builder.WriteString("[")
    for node := list.Head; node != nil; node = node.Next {
        var data any = node.Data
        if str, ok := data.(string); ok {
            fmt.Fprintf(&builder, "%q", str)
        } else {
            fmt.Fprintf(&builder, "%v", node.Data)
        }
        if node.Next != nil {
            builder.WriteString("<->")
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
func (list *GoList2[T]) appendHead(value T) *GoList2[T] {
    node := &node.Node2[T]{Data: value, Next: list.Head}
    if list.Head != nil {
        list.Head.Prev = node
    }
    list.Head = node
    return list
}

// Do reverse the list.
func (list *GoList2[T]) reverse() *GoList2[T] {
    var prev *node.Node2[T]
    for node := list.Head; node != nil; node = node.Prev {
        node.Prev, node.Next = node.Next, node.Prev
        prev = node
    }
    list.Head = prev
    return list
}

// Do quick sort input list.
func quickSort[T constraints.Ordered](list GoList2[T]) GoList2[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoList2[T]

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

// Do quick sort input list and remove duplicate nodes.
func uniqueQuickSort[T constraints.Ordered](list GoList2[T]) GoList2[T] {
    if list.Head == nil || list.Head.Next == nil {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoList2[T]
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
