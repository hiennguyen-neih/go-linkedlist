// Package golist contains functions and methods for singly circular linked
// list in Go.
package golistc

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

// Struct of Go singly circular linked list.
type GoListC[T any] struct {
    Head *node.Node[T] // First node of the list.
    Tail *node.Node[T] // Last node of the list.
}

/*
 *******************************************************************************
 * Exported functions
 *******************************************************************************
 */

// Create new singly circular linked list from input values.
func New[T any](values ...T) GoListC[T] {
    var list GoListC[T]
    for _, val := range values {
        list.append(val)
    }
    return list
}

// Convert input slice into new singly circular linked list.
func FromSlice[T any](values []T) GoListC[T] {
    var list GoListC[T]
    for _, val := range values {
        list.append(val)
    }
    return list
}

// Convert input singly circular linked list into new slice.
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
// false. If input list is an empty list, returns true.
func All[T any](list GoListC[T], fun func(T) bool) bool {
    node := list.Head
    if node == nil {
        return true
    }
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
// returns false. If input list is an empty list, returns false.
func Any[T any](list GoListC[T], fun func(T) bool) bool {
    node := list.Head
    if node == nil {
        return false
    }
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

// Appends values into last of input list preserving values order.
func Append[T any](list GoListC[T], values ...T) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node != nil {
        for {
            result.append(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }
    for _, value := range values {
        result.append(value)
    }
    return result
}

// Appends values into head of input list preserving values order.
func AppendHead[T any](list GoListC[T], values ...T) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    for _, value := range values {
        result.append(value)
    }
    if node != nil {
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

// Returns a list that is concatenated of all input lists.
func Concat[T any](lists ...GoListC[T]) GoListC[T] {
    var result GoListC[T]
    for _, list := range lists {
        if list.Head == nil {
            continue
        }
        for node := list.Head; ; {
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
    node := list.Head
    if node == nil {
        return result
    }
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
    if len == 0 {
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
// a new empty list.
func DropLast[T any](list GoListC[T]) GoListC[T] {
    var result GoListC[T]
    node := list.Head
    if node == nil || node == node.Next {
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
    if node1 == nil || node2 == nil {
        return node1 == nil && node2 == nil
    }
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

// Calls fun(data, acc) on successive nodes of list from Head to Tail (from
// start of list to end of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldl[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    node := list.Head
    if node == nil {
        return acc0
    }
    for {
        acc0 = fun(node.Data, acc0)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return acc0
}

// Calls fun(data, acc) on successive nodes of list from Tail to Head (from
// end of list to start of list), starting with acc0. Input fun must return a
// new accumulator, which is passed to the next call. The function returns the
// final value of the accumulator. Input acc0 is returned if the list is empty.
func Foldr[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) T2) T2 {
    reverse := Reverse(list)
    node := reverse.Head
    if node == nil {
        return acc0
    }
    for {
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
    node := list.Head
    if node != nil {
        for {
            fun(node.Data)
            node = node.Next
            if node == list.Head {
                break
            }
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
    if len == 0 && index == 0 {
        result.append(val)
    } else if index == len {
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
    node := list.Head
    if node == nil {
        return result
    }
    for {
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

// Combines the operations of Map function and Foldl function into one pass.
func MapFoldl[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoListC[T1], T2) {
    var value T1
    var result GoListC[T1]
    if list.Head == nil {
        return result, acc0
    }
    for node := list.Head; ; {
        value, acc0 = fun(node.Data, acc0)
        result.append(value)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result, acc0
}

// Combines the operations of Map function and Foldr function into one pass.
func MapFoldr[T1, T2 any](list GoListC[T1], acc0 T2, fun func(T1, T2) (T1, T2)) (GoListC[T1], T2) {
    var value T1
    var result GoListC[T1]
    if list.Head == nil {
        return result, acc0
    }
    reverse := Reverse(list)
    for node := reverse.Head; ; {
        value, acc0 = fun(node.Data, acc0)
        result.appendHead(value)
        node = node.Next
        if node == reverse.Head {
            break
        }
    }
    return result, acc0
}

// Returns the first node in list that compares greater than or equal to all
// other nodes of list. This function only works with constraint Ordered list.
// If input list is an empty list, returns nil.
func Max[T constraints.Ordered](list GoListC[T]) *node.Node[T] {
    node := list.Head
    max := node
    if node != nil {
        for {
            if node.Data > max.Data {
                max = node
            }
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }
    return max
}

// Returns true if elem matches some node data of list, otherwise returns false.
func Member[T any](list GoListC[T], elem T) bool {
    node := list.Head
    if node == nil {
        return false
    }
    for {
        if cmp.Equal(node.Data, elem) {
            return true
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return false
}

// Returns a sorted list forming by merging all input lists. This function only
// works with constraint Ordered lists.
func Merge[T constraints.Ordered](lists ...GoListC[T]) GoListC[T] {
    result := Concat(lists...)
    return Sort(result)
}

// Returns the first node in list that compares less than or equal to all
// other nodes of list. This function only works with constraint Ordered list.
// If input list is an empty list, returns nil.
func Min[T constraints.Ordered](list GoListC[T]) *node.Node[T] {
    node := list.Head
    min := node
    if node != nil {
        for {
            if node.Data < min.Data {
                min = node
            }
            node = node.Next
            if node == list.Head {
                break
            }
        }
    }
    return min
}

// Returns node in list at specific index. index is capped at list length.
// Negative index indicate an offset from the end of list.
func Nth[T any](list GoListC[T], index int) *node.Node[T] {
    len := Len(list)
    if len == 0 {
        panic("Nth, list is empty!")
    }

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
func NthTail[T any](list GoListC[T], index int) GoListC[T] {
    len := Len(list)
    if len == 0 {
        panic("NthTail, list is empty!")
    }

    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index < 0 || index >= len {
        panic("NthTail, index is out of bound!")
    }

    var result GoListC[T]
    node := list.Head
    for i := 0; i < index; i++ {
        node = node.Next
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

// Partitions input list into list1 and list2, where list1 contains nodes
// which fun returns true and list2 contains nodes which fun returns false.
func Partition[T any](list GoListC[T], fun func(T) bool) (GoListC[T], GoListC[T]) {
    var list1 GoListC[T]
    var list2 GoListC[T]
    node := list.Head
    if node == nil {
        return list1, list2
    }
    for {
        if fun(node.Data) {
            list1.append(node.Data)
        } else {
            list2.append(node.Data)
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return list1, list2
}

// Returns true if list1 is a prefix of list2, otherwise returns false.
// A prefix of a list is the first part of the list, starting from the
// beginning and stopping at any point.
func Prefix[T any](list1, list2 GoListC[T]) bool {
    node1 := list1.Head
    node2 := list2.Head
    for ; node1 != nil; {
        if node2 == nil || !cmp.Equal(node1.Data, node2.Data) {
            return false
        }
        node1 = node1.Next
        node2 = node2.Next
        if node1 == list1.Head {
            break
        }
    }
    return true
}

// Returns a list that node at specific index is replaced with val. If index
// is out of bound, the original list is returned. Negative index indicate an
// offset from the end of list.
func ReplaceAt[T any](list GoListC[T], index int, val T) GoListC[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index < 0 || index >= len {
        panic("ReplaceAt, index is out of bound!")
    }

    var result GoListC[T]
    i := 0
    for node := list.Head; ; {
        if i == index {
            result.append(val)
        } else {
            result.append(node.Data)
        }
        i++
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
    node := list.Head
    if node == nil {
        return result
    }
    for {
        result.appendHead(node.Data)
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return result
}

// Returns position and first node in list that fun returns true. If list is
// empty or every fun execution returns false, returns position is -1.
func Search[T any](list GoListC[T], fun func(T) bool) (int, *node.Node[T]) {
    node := list.Head
    if node == nil {
        return -1, nil
    }
    for i := 0; ; i++ {
        if fun(node.Data) {
            return i, node
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }
    return -1, nil
}

// Returns a list containing the sorted nodes data of input list. This function
// only works with constraint Ordered list.
func Sort[T constraints.Ordered](list GoListC[T]) GoListC[T] {
    return quickSort(list)
}

// Returns true if list1 is a suffix of list2, otherwise returns false.
// A suffix of a list if the last part of the list, starting from any position
// and going all the way to the end.
func Suffix[T any](list1, list2 GoListC[T]) bool {
    len1 := Len(list1)
    len2 := Len(list2)
    if len1 > len2 {
        return false
    } else if len1 == 0 {
        return true
    }

    node1 := list1.Head
    node2 := Nth(list2, len2-len1)
    for {
        if !cmp.Equal(node1.Data, node2.Data) {
            return false
        }
        node1 = node1.Next
        node2 = node2.Next
        if node1 == list1.Head {
            break
        }
    }
    return true
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

// Returns a sorted list formed by merging all input lists, while removing
// duplicates. This function only works with constraint Ordered lists.
func UMerge[T constraints.Ordered](lists ...GoListC[T]) GoListC[T] {
    result := Concat(lists...)
    return uniqueQuickSort(result)
}

// Returns a sorted list of the nodes data of list, keeping only the first
// occurrence of nodes that compare equal and removing duplicates. This
// function only works with constraint Ordered list.
func USort[T constraints.Ordered](list GoListC[T]) GoListC[T] {
    return uniqueQuickSort(list)
}

// Returns a list that node at specific index is updated with returns value of
// fun. If index is out of bound, the original list is returned. Negative index
// indicate an offset from the end of list.
func UpdateAt[T any](list GoListC[T], index int, fun func(T) T) GoListC[T] {
    len := Len(list)
    if index < 0 {
        index = len + index // same as len - abs(index)
    }

    if index < 0 || index >= len {
        panic("UpdateAt, index is out of bound!")
    }

    var result GoListC[T]
    i := 0
    for node := list.Head; ; {
        if i == index {
            result.append(fun(node.Data))
        } else {
            result.append(node.Data)
        }
        i++
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

// Returns a string representing the singly circular linked list.
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

// Do concatenate all input lists.
func (list *GoListC[T]) concat(lists ...GoListC[T]) *GoListC[T] {
    for _, other := range lists {
        if other.Head == nil {
            continue
        }
        if list.Head == nil {
            list.Head = other.Head
            list.Tail = other.Tail
            continue
        }
        list.Tail.Next = other.Head
        other.Tail.Next = list.Head
        list.Tail = other.Tail
    }
    return list
}

// Do quick sort input list.
func quickSort[T constraints.Ordered](list GoListC[T]) GoListC[T] {
    if list.Head == nil || list.Head.Next == list.Head {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoListC[T]

    // Partitioning
    for node := list.Head; ; {
        switch {
        case node.Data < pivot:
            less.append(node.Data)
        case node.Data == pivot:
            equal.append(node.Data)
        case node.Data > pivot:
            greater.append(node.Data)
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }

    // Recursive sort
    sortedLess := quickSort(less)
    sortedGreater := quickSort(greater)

    // Concatenates 3 lists: sortedLess + equal + sortedGreater
    return *sortedLess.concat(equal, sortedGreater)
}

// Do quick sort input list and remove duplicate nodes.
func uniqueQuickSort[T constraints.Ordered](list GoListC[T]) GoListC[T] {
    if list.Head == nil || list.Head.Next == list.Head {
        return list
    }

    pivot := list.Head.Data
    var less, equal, greater GoListC[T]
    seen := make(map[T]bool) // store already seen node data into map

    // Partitioning and remove seen node
    for node := list.Head; ; {
        if !seen[node.Data] {
            seen[node.Data] = true

            switch {
            case node.Data < pivot:
                less.append(node.Data)
            case node.Data == pivot:
                equal.append(node.Data)
            case node.Data > pivot:
                greater.append(node.Data)
            }
        }
        node = node.Next
        if node == list.Head {
            break
        }
    }

    // Recursive sort
    sortedLess := uniqueQuickSort(less)
    sortedGreater := uniqueQuickSort(greater)

    // Concatenate: sortedLess + equal + sortedGreater
    return *sortedLess.concat(equal, sortedGreater)
}
