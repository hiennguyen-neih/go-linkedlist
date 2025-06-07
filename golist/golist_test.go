package golist

import (
    "testing"
    "reflect"
)

func TestNewAndToSlice(t *testing.T) {
    list := New(1, 2, 3, 4)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(New(...)): %v\nexpected: %v", result, expected)
    }
}

func TestFromSliceAndToSlice(t *testing.T) {
    list := FromSlice([]int{1, 2, 3, 4})
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(FromSlice(...)): %v\nexpected: %v", result, expected)
    }
}

func TestAll(t *testing.T) {
    list1 := New(2, 4, 6, 8)
    list2 := New(2, 4, 6, 9)
    if result1 := All(list1, func(n int) bool {return n%2 == 0}); !result1 {
        t.Errorf("All\nresult1: %v\nexpected1: true", result1)
    }
    if result2 := All(list2, func(n int) bool {return n%2 == 0}); result2 {
        t.Errorf("All\nresult2: %v\nexpected2: false", result2)
    }
}

func TestAny(t *testing.T) {
    list1 := New(2, 4, 6, 8)
    list2 := New(2, 4, 5, 8)
    if result1 := Any(list1, func(n int) bool {return n%2 != 0}); result1 {
        t.Errorf("Any\nresult1: %v\nexpected: false", result1)
    }
    if result2 := Any(list2, func(n int) bool {return n%2 != 0}); !result2 {
        t.Errorf("Any\nresult2: %v\nexpected: true", result2)
    }
}

func TestAppend(t *testing.T) {
    list := New(1, 2)
    appended := Append(list, 3, 4)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(appended); !reflect.DeepEqual(result, expected) {
        t.Errorf("Append\nresul: %v\nexpected: %v", result, expected)
    }
}

func TestAppendHead(t *testing.T) {
    list := New(3, 4)
    appended := AppendHead(list, 1, 2)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(appended); !reflect.DeepEqual(result, expected) {
        t.Errorf("AppendHead\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestConcat(t *testing.T) {
    list1 := New(1, 2, 3)
    list2 := New(4, 5, 6)
    list3 := New(7, 8, 9)
    concatenated := Concat(list1, list2, list3)
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if result := ToSlice(concatenated); !reflect.DeepEqual(result, expected) {
        t.Errorf("Concat\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDeleteNormalCase(t *testing.T) {
    list := New(1, 2, 3, 2, 4)
    deleted := Delete(list, 2)
    expected := []int{1, 3, 2, 4}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("Delete\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDeleteEmptyList(t *testing.T) {
    list := New[int]()
    deleted := Delete(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
}

func TestDeleteAtNormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")
    deleted := DeleteAt(list, -2)
    expected := []string{"a", "b", "d"}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("DeleteAt\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDeleteAtEmptyList(t *testing.T) {
    list := New[int]()
    deleted := DeleteAt(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
}

func TestDeleteAtOutOfBound(t *testing.T) {
    list := New("a", "b", "c", "d")
    deleted := DeleteAt(list, 4)
    expected := []string{"a", "b", "c", "d"}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("DeleteAt\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDropLast(t *testing.T) {
    list := New("a", "b", "c", "d")
    droped := DropLast(list)
    expected := []string{"a", "b", "c"}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected) {
        t.Errorf("DropLast\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDropWhileTakeWhile(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 2)
    droped := DropWhile(list, func(n int) bool { return n < 4 })
    taken := TakeWhile(list, func(n int) bool { return n < 4 })
    expected1 := []int{4, 5, 2}
    expected2 := []int{1, 2, 3}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected1) {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected1)
    }
    if result := ToSlice(taken); !reflect.DeepEqual(result, expected2) {
        t.Errorf("TakeWhile\nresult: %v\nexpected: %v", result, expected2)
    }
}

func TestDuplicate(t *testing.T) {
    duplicate := Duplicate(4, 0)
    expected := []int{0, 0, 0, 0}
    if result := ToSlice(duplicate); !reflect.DeepEqual(result, expected) {
        t.Errorf("Duplicate\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFilterNormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := Filter(list, func(n int) bool { return n%2 == 0 })
    expected := []int{2, 4, 6}
    if result := ToSlice(filtered); !reflect.DeepEqual(result, expected) {
        t.Errorf("Filter\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFilterEmptyList(t *testing.T) {
    list := New[int]()
    filtered := Filter(list, func(n int) bool { return n%2 != 0 })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("Filter\nresult: %v\nexpected: []", result)
    }
}

func TestFilterMapNormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := FilterMap(list, func(n int) (bool, int) {
        return n % 2 != 0, n * n
    })
    expected := []int{1, 9, 25}
    if result := ToSlice(filtered); !reflect.DeepEqual(result, expected) {
        t.Errorf("FilterMap\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFilterMapEmptyList(t *testing.T) {
    list := New[int]()
    filtered := FilterMap(list, func(n int) (bool, int) {
        return n % 2 == 0, n * 2
    })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("FilterMap\nresult: %v\nexpected: []", result)
    }
}

func TestFindFound(t *testing.T) {
    list := New(1, 2, 3, 4)
    result := Find(list, 3)
    expected := 2
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFindNotFound(t *testing.T) {
    list := New(1, 2, 3, 4)
    result := Find(list, 5)
    expected := -1
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldl(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    result := Foldl(list, 0, func(n, s int) int { return n + s })
    expected := 15
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldr(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    result := Foldr(list, 1, func(n, s int) int { return n * s })
    expected := 120
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestForEach(t *testing.T) {
    list := New(1, 2, 3, 4, 5)

    var result []int
    ForEach(list, func(val int) {
        result = append(result, val)
    })

    expected := []int{1, 2, 3, 4, 5}
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("ForEach\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestInsertAtNormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")

    inserted1 := InsertAt(list, -2, "X")
    expected1 := []string{"a", "b", "X", "c", "d"}
    if result1 := ToSlice(inserted1); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("InsertAt\nresult: %v\nexpected: %v", result1, expected1)
    }

    inserted2 := InsertAt(list, 4, "X")
    expected2 := []string{"a", "b", "c", "d", "X"}
    if result2 := ToSlice(inserted2); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("InsertAt\nresult: %v\nexpected: %v", result2, expected2)
    }
}

func TestInsertAtIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("InsertAt\nExpect panic")
        } else if r != "InsertAt, index is out of bound!" {
            t.Errorf("InsertAt\nWrong panic message")
        }
    }()
    InsertAt(New(1, 2, 3, 4), 10, 0)
}

func TestJoin(t *testing.T) {
    list := New("a", "b", "c", "d")
    joined := Join(list, "X")
    expected := []string{"a", "X", "b", "X", "c", "X", "d"}
    if result := ToSlice(joined); !reflect.DeepEqual(result, expected) {
        t.Errorf("Join\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestLast(t *testing.T) {
    last1 := Last(New(1, 2, 3, 4))
    expected1 := 4
    if last1.Data != expected1 {
        t.Errorf("Last\nresult: %v\nexpected: %v", last1, expected1)
    }

    last2 := Last(New[int]())
    if last2 != nil {
        t.Errorf("Last\nresult: %v\nexpected: nil", last2)
    }
}

func TestMap(t *testing.T) {
    list := New(1, 2, 3, 4)
    mapped := Map(list, func(n int) int { return n * n })
    expected := []int{1, 4, 9, 16}
    if result := ToSlice(mapped); !reflect.DeepEqual(result, expected) {
        t.Errorf("Map\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestMapFoldlMapFoldr(t *testing.T) {
    list := New(1, 2, 3, 4)
    mapped1, sum := MapFoldl(list, 0, func(n, s int) (int, int) {
        return n * 2, s + n
    })
    mapped2, fac := MapFoldr(list, 1, func(n, f int) (int, int) {
        return n * 2, f * n
    })
    expectedL := []int{2, 4, 6, 8}
    expectedS := 10
    expectedF := 24
    if result1 := ToSlice(mapped1); !reflect.DeepEqual(result1, expectedL) || sum != expectedS {
        t.Errorf("MapFoldl\nresult: %v - %v\nexpected: %v - %v", sum, result1, expectedS, expectedL)
    }
    if result2 := ToSlice(mapped2); !reflect.DeepEqual(result2, expectedL) || fac != expectedF {
        t.Errorf("MapFoldr\nresult: %v - %v\nexpected: %v - %v", fac, result2, expectedF, expectedL)
    }
}

func TestMaxMin(t *testing.T) {
    list := New("d", "b", "e", "a", "c")
    max := Max(list)
    min := Min(list)
    expectedMax := "e"
    expectedMin := "a"

    if max.Data != expectedMax {
        t.Errorf("Max\nresult: %v\nexpected: %v", max, expectedMax)
    }
    if min.Data != expectedMin {
        t.Errorf("Min\nresult: %v\nexpected: %v", min, expectedMin)
    }
}

func TestMember(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    member1 := Member(list, 4)
    member2 := Member(list, 6)
    if !member1 {
        t.Errorf("Member\nresult: %v\nexpected: true", member1)
    }
    if member2 {
        t.Errorf("Member\nresult: %v\nexpected: false", member2)
    }
}

func TestMergeUMerge(t *testing.T) {
    list1 := New(2, 8, 6)
    list2 := New(1, 3, 3)
    list3 := New(8, 4, 5)
    merged := Merge(list1, list2, list3)
    umerged := UMerge(list1, list2, list3)

    expected1 := []int{1, 2, 3, 3, 4, 5, 6, 8, 8}
    expected2 := []int{1, 2, 3, 4, 5, 6, 8}

    if result1 := ToSlice(merged); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("Merge\nresult: %v\nexpected: %v", result1, expected1)
    }
    if result2 := ToSlice(umerged); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("UMerge\nresult: %v\nexpected: %v", result2, expected2)
    }
}

func TestNthNormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    nth1 := Nth(list, 2)
    nth2 := Nth(list, -2)

    expected1 := 3
    expected2 := 4

    if nth1.Data != expected1 {
        t.Errorf("Nth\nresult: %v\nexpected: %v", nth1, expected1)
    }
    if nth2.Data != expected2 {
        t.Errorf("Nth\nresult: %v\nexpected: %v", nth2, expected2)
    }
}

func TestNthIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Nth\nExpect panic")
        } else if r != "Nth, index is out of bound!" {
            t.Errorf("Nth\nWrong panic message")
        }
    }()
    Nth(New(1, 2, 3, 4), 10)
}

func TestNthTailNormalCase(t *testing.T) {
    list := New("a", "b", "c", "d", "e")
    tail := NthTail(list, -3)
    expected := []string{"c", "d", "e"}

    if result := ToSlice(tail); !reflect.DeepEqual(result, expected) {
        t.Errorf("NthTail\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestNthTailIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("NthTail\nExpect panic")
        } else if r != "NthTail, index is out of bound!" {
            t.Errorf("NthTail\nWrong panic message")
        }
    }()
    NthTail(New(1, 2, 3, 4), -10)
}

func TestPartition(t *testing.T) {
    input := New(1, 2, 3, 4, 5, 6)
    list1, list2 := Partition(input, func(n int) bool { return n % 2 != 0 })
    expected1 := []int{1, 3, 5}
    expected2 := []int{2, 4, 6}

    if result1 := ToSlice(list1); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("Partition\nresult: %v\nexpected: %v", result1, expected1)
    }
    if result2 := ToSlice(list2); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("Partition\nresult: %v\nexpected: %v", result2, expected2)
    }
}

func TestPreffixSuffix(t *testing.T) {
    list1 := New("a", "b")
    list2 := New("e", "f")
    list3 := New("a", "b", "c", "d", "e", "f")

    if result := Prefix(list1, list3); !result {
        t.Errorf("Prefix\nresult: %v\nexpected: true", result)
    }
    if result := Prefix(list2, list3); result {
        t.Errorf("Prefix\nresult: %v\nexpected: false", result)
    }

    if result := Suffix(list2, list3); !result {
        t.Errorf("Suffix\nresult: %v\nexpected: true", result)
    }
    if result := Suffix(list1, list3); result {
        t.Errorf("Suffix\nresult: %v\nexpected: false", result)
    }
}

func TestReplaceAtUpdateAt(t *testing.T) {
    list := New(1, 2, 3, 4)
    replaced := ReplaceAt(list, -2, 0)
    updated := UpdateAt(list, -3, func(n int) int { return n * n })

    expected1 := []int{1, 2, 0, 4}
    expected2 := []int{1, 4, 3, 4}

    if result := ToSlice(replaced); !reflect.DeepEqual(result, expected1) {
        t.Errorf("ReplaceAt\nresult: %v\nexpected: %v", result, expected1)
    }
    if result := ToSlice(updated); !reflect.DeepEqual(result, expected2) {
        t.Errorf("UpdateAt\nresult: %v\nexpected: %v", result, expected2)
    }
}

func TestSearch(t *testing.T) {
    list := New(1, 2, 3, 4)
    index1, node1 := Search(list, func(n int) bool { return n % 2 == 0 })
    index2, _ := Search(list, func(n int) bool { return n > 4 })

    if index1 != 1 || node1.Data != 2 {
        t.Errorf("Search\nresult: %v - %v\nexpected: 1 - 2", index1, node1)
    }
    if index2 != -1 {
        t.Errorf("Search\nresult: %v\nexpected: -1", index2)
    }
}

func TestSeq(t *testing.T) {
    list := Seq(1, 10, 2)
    expected := []int{1, 3, 5, 7, 9}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("Seq\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestSplitNormalCase(t *testing.T) {
    list1, list2 := Split(New("a", "b", "c", "d", "e"), -3)
    expected1 := []string{"a", "b"}
    expected2 := []string{"c", "d", "e"}

    if result := ToSlice(list1); !reflect.DeepEqual(result, expected1) {
        t.Errorf("Split\nresult: %v\nexpected: %v", result, expected1)
    }
    if result := ToSlice(list2); !reflect.DeepEqual(result, expected2) {
        t.Errorf("Split\nresult: %v\nexpected: %v", result, expected2)
    }
}

func TestSplitIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Split\nExpect panic")
        } else if r != "Split, n is out of bound!" {
            t.Errorf("Split\nWrong panic message")
        }
    }()
    Split(New(1, 2, 3, 4, 5), 6)
}

func TestSplitWith(t *testing.T) {
    list := New(1, 2, 3, 4, 1, 3)
    list1, list2 := SplitWith(list, func(n int) bool { return n < 4 })
    expected1 := []int{1, 2, 3}
    expected2 := []int{4, 1, 3}
    if result := ToSlice(list1); !reflect.DeepEqual(result, expected1) {
        t.Errorf("SplitWith\nresult: %v\nexpected: %v", result, expected1)
    }
    if result := ToSlice(list2); !reflect.DeepEqual(result, expected2) {
        t.Errorf("SplitWith\nresult: %v\nexpected: %v", result, expected2)
    }
}

func TestSublistNormalCase(t *testing.T) {
    list := New("a", "b", "c", "d", "e", "f")
    sublist := Sublist(list, 2, 3)
    expected := []string{"c", "d", "e"}
    if result := ToSlice(sublist); !reflect.DeepEqual(result, expected) {
        t.Errorf("Sublist\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestSublistNegativeLen(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Sublist\nExpect panic")
        } else if r != "Sublist, input len must not be negative!" {
            t.Errorf("Sublist\nWrong panic message")
        }
    }()
    Sublist(New(1, 2, 3, 4), 2, -2)
}

func TestSublistStartOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Sublist\nExpect panic")
        } else if r != "Sublist, start is out of bound!" {
            t.Errorf("Sublist\nWrong panic message")
        }
    }()
    Sublist(New(1, 2, 3, 4), -5, 2)
}

func TestSubtract(t *testing.T) {
    list1 := New("a","b","c","b","a","b")
    list2 := New("b","a","b")
    subtract := Subtract(list1, list2)
    expected := []string{"c", "a", "b"}
    if result := ToSlice(subtract); !reflect.DeepEqual(result, expected) {
        t.Errorf("Subtract\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestSum(t *testing.T) {
    list := New("a", "b", "c", "d", "e", "f")
    sum := Sum(list)
    expected := "abcdef"
    if sum != expected {
        t.Errorf("Sum\nresult: %v\nexpected: %v", sum, expected)
    }
}

func TestUSort(t *testing.T) {
    list := New(2, 5, 1, 2, 7, 3, 9, 4, 8, 6, 4)
    sorted := USort(list)
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if result := ToSlice(sorted); !reflect.DeepEqual(result, expected) {
        t.Errorf("USort\nresult: %v\nexpected: %v", result, expected)
    }
}
