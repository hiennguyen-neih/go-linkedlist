package golistc

import (
    "testing"
    "reflect"
)

/*
 *******************************************************************************
 * Test cases
 *******************************************************************************
 */

func TestNew_ToSlice(t *testing.T) {
    list := New(1, 2, 3, 4)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(New(...)): %v\nexpected: %v", result, expected)
    }
    assertCircular(t, list)
}

func TestFromSlice_ToSlice(t *testing.T) {
    list := FromSlice([]int{1, 2, 3, 4})
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(FromSlice(...)): %v\nexpected: %v", result, expected)
    }
    assertCircular(t, list)
}

func TestGoListString_Float(t *testing.T) {
    list := New(0.1, 0.2, 0.3, 0.4, 0.5)
    expected := "[0.1=>0.2=>0.3=>0.4=>0.5=>]"
    if result := list.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestGoListString_String(t *testing.T) {
    list := New("A", "B", "C", "D", "E", "F")
    expected := `["A"=>"B"=>"C"=>"D"=>"E"=>"F"=>]`
    if result := list.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestGoListString_Empty(t *testing.T) {
    list := New[string]()
    expected := `[]`
    if result := list.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestAll_NormalCase(t *testing.T) {
    list1 := New(1, 3, 5, 7)
    list2 := New(1, 3, 5, 8)
    if result1 := All(list1, func(n int) bool { return n%2 != 0 }); false {
        t.Errorf("All\nresult1: %v\nexpected: true", result1)
    }
    if result2 := All(list2, func(n int) bool { return n%2 != 0 }); false {
        t.Errorf("All\nresult2: %v\nexpected: false", result2)
    }
}

func TestAll_EmptyList(t *testing.T) {
    list := New[int]()
    if result := All(list, func(n int) bool { return n%2 != 0 }); false {
        t.Errorf("All\nresult: %v\nexpected: true", result)
    }
}

func TestAny_NormalCase(t *testing.T) {
    list1 := New(2, 4, 6, 7)
    list2 := New(2, 4, 6, 8)
    if result1 := Any(list1, func(n int) bool { return n%2 != 0}); false {
        t.Errorf("Any\nresult1: %v\nexpected: true", result1)
    }
    if result2 := Any(list2, func(n int) bool { return n%2 != 0}); false {
        t.Errorf("Any\nresult2: %v\nexpected: false", result2)
    }
}

func TestAny_EmptyList(t *testing.T) {
    list := New[int]()
    if result := Any(list, func(n int) bool { return n%2 != 0}); false {
        t.Errorf("Any\nresult: %v\nexpected: false", result)
    }
}

func TestAppend(t *testing.T) {
    list := New(1, 2)
    appended := Append(list, 3, 4)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(appended); !reflect.DeepEqual(result, expected) {
        t.Errorf("Append\nresul: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, appended)
}

func TestAppendHead(t *testing.T) {
    list := New("c", "d")
    appended := AppendHead(list, "a", "b")
    expected := []string{"a", "b", "c", "d"}
    if result := ToSlice(appended); !reflect.DeepEqual(result, expected) {
        t.Errorf("AppendHead\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, appended)
}

func TestConcat_NormalCase(t *testing.T) {
    list1 := New(1, 2, 3)
    list2 := New(4, 5, 6)
    list3 := New(7, 8, 9)
    concatenated := Concat(list1, list2, list3)
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if result := ToSlice(concatenated); !reflect.DeepEqual(result, expected) {
        t.Errorf("Concat\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, concatenated)
}

func TestConcat_EmptyList(t *testing.T) {
    list1 := New[int]()
    list2 := New[int]()
    list3 := New[int]()
    concatenated := Concat(list1, list2, list3)
    if result := ToSlice(concatenated); len(result) != 0 {
        t.Errorf("Concat\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, concatenated)
}

func TestConcat_MixList(t *testing.T) {
    list1 := New(1, 2, 3)
    list2 := New[int]()
    list3 := New(7, 8, 9)
    concatenated := Concat(list1, list2, list3)
    expected := []int{1, 2, 3, 7, 8, 9}
    if result := ToSlice(concatenated); !reflect.DeepEqual(result, expected) {
        t.Errorf("Concat\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, concatenated)
}

func TestDelete_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 2, 4)
    deleted := Delete(list, 2)
    expected := []int{1, 3, 2, 4}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("Delete\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, deleted)
}

func TestDelete_Singleton(t *testing.T) {
    list := New(1)
    deleted := Delete(list, 1)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, deleted)
}

func TestDelete_EmptyList(t *testing.T) {
    list := New[int]()
    deleted := Delete(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, deleted)
}

func TestDelete_NotFound(t *testing.T) {
    list := New(1, 2, 3, 2, 4)
    deleted := Delete(list, 5)
    expected := []int{1, 2, 3, 2, 4}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("Delete\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, deleted)
}

func TestDeleteAt_NormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")
    deleted := DeleteAt(list, -2)
    expected := []string{"a", "b", "d"}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("DeleteAt\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, deleted)
}

func TestDeleteAt_EmptyList(t *testing.T) {
    list := New[int]()
    deleted := DeleteAt(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("DeleteAt\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, deleted)
}

func TestDeleteAt_IndexOutOfBound(t *testing.T) {
    list := New("a", "b", "c", "d")
    expected := []string{"a", "b", "c", "d"}

    deleted1 := DeleteAt(list, 4)
    if result1 := ToSlice(deleted1); !reflect.DeepEqual(result1, expected) {
        t.Errorf("DeleteAt\nresult1: %v\nexpected: %v", result1, expected)
    }
    assertCircular(t, deleted1)

    deleted2 := DeleteAt(list, -5)
    if result2 := ToSlice(deleted2); !reflect.DeepEqual(result2, expected) {
        t.Errorf("DeleteAt\nresult2: %v\nexpected: %v", result2, expected)
    }
    assertCircular(t, deleted2)
}

func TestDropLast_NormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")
    droped := DropLast(list)
    expected := []string{"a", "b", "c"}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected) {
        t.Errorf("DropLast\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, droped)
}

func TestDropLast_Singleton(t *testing.T) {
    list := New("a")
    droped := DropLast(list)
    if result := ToSlice(droped); len(result) != 0 {
        t.Errorf("DropLast\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, droped)
}

func TestDropLast_EmptyList(t *testing.T) {
    list := New[int]()
    droped := DropLast(list)
    if result := ToSlice(droped); len(result) != 0 {
        t.Errorf("DropLast\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, droped)
}

func TestDropWhile_TakeWhile(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 2)

    droped := DropWhile(list, func(n int) bool { return n < 4 })
    expected1 := []int{4, 5, 2}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected1) {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected1)
    }
    assertCircular(t, droped)

    taken := TakeWhile(list, func(n int) bool { return n < 4 })
    expected2 := []int{1, 2, 3}
    if result := ToSlice(taken); !reflect.DeepEqual(result, expected2) {
        t.Errorf("TakeWhile\nresult: %v\nexpected: %v", result, expected2)
    }
    assertCircular(t, taken)
}

func TestDropWhile_TakeWhile_EmptyList(t *testing.T) {
    list := New[string]()

    droped := DropWhile(list, func(n string) bool { return n == "a" })
    expected1 := []string{}
    if result := ToSlice(droped); len(result) != 0 {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected1)
    }
    assertCircular(t, droped)

    taken := TakeWhile(list, func(n string) bool { return n == "a" })
    expected2 := []string{}
    if result := ToSlice(taken); len(result) != 0 {
        t.Errorf("TakeWhile\nresult: %v\nexpected: %v", result, expected2)
    }
    assertCircular(t, taken)
}

func TestDropWhile_TakeWhile_AllOfList(t *testing.T) {
    list := New[string]("a", "a", "a", "a")

    droped := DropWhile(list, func(n string) bool { return n == "a" })
    expected1 := []string{}
    if result := ToSlice(droped); len(result) != 0 {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected1)
    }
    assertCircular(t, droped)

    taken := TakeWhile(list, func(n string) bool { return n == "a" })
    expected2 := []string{"a", "a", "a", "a"}
    if result := ToSlice(taken); !reflect.DeepEqual(result, expected2) {
        t.Errorf("TakeWhile\nresult: %v\nexpected: %v", result, expected2)
    }
    assertCircular(t, taken)
}

func TestDropWhile_TakeWhile_Boundary(t *testing.T) {
    list := New[string]("a", "a", "a", "a")

    droped := DropWhile(list, func(n string) bool { return n == "b" })
    expected1 := []string{"a", "a", "a", "a"}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected1) {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected1)
    }
    assertCircular(t, droped)

    taken := TakeWhile(list, func(n string) bool { return n == "b" })
    expected2 := []string{}
    if result := ToSlice(taken); len(result) != 0 {
        t.Errorf("TakeWhile\nresult: %v\nexpected: %v", result, expected2)
    }
    assertCircular(t, taken)
}

func TestDuplicate_NormalCase(t *testing.T) {
    duplicate := Duplicate(4, 0)
    expected := []int{0, 0, 0, 0}
    if result := ToSlice(duplicate); !reflect.DeepEqual(result, expected) {
        t.Errorf("Duplicate\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, duplicate)
}

func TestDuplicate_EmptyList(t *testing.T) {
    duplicate1 := Duplicate(0, "X")
    if result := ToSlice(duplicate1); len(result) != 0 {
        t.Errorf("Duplicate\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, duplicate1)

    duplicate2 := Duplicate(-1, "X")
    if result := ToSlice(duplicate2); len(result) != 0 {
        t.Errorf("Duplicate\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, duplicate2)
}

func TestEqual_ReturnTrue(t *testing.T) {
    list1 := New(1, 2, 3, 4)
    list2 := New(1, 2, 3, 4)
    if !Equal(list1, list2) {
        t.Errorf("Equal\nExpected true but got false")
    }
}

func TestEqual_SameLenReturnFalse(t *testing.T) {
    list1 := New(1, 2, 3, 4)
    list2 := New(1, 2, 4, 3)
    if Equal(list1, list2) {
        t.Errorf("Equal\nExpected false but got true")
    }
}

func TestEqual_List1Longer(t *testing.T) {
    list1 := New(1, 2, 3, 4, 5)
    list2 := New(1, 2, 3, 4)
    if Equal(list1, list2) {
        t.Errorf("Equal\nExpected false but got true")
    }
}

func TestEqual_List2Longer(t *testing.T) {
    list1 := New(1, 2, 3, 4)
    list2 := New(1, 2, 3, 4, 5)
    if Equal(list1, list2) {
        t.Errorf("Equal\nExpected false but got true")
    }
}

func TestEqual_List1Empty(t *testing.T) {
    list1 := New[int]()
    list2 := New(1, 2, 3, 4)
    if Equal(list1, list2) {
        t.Errorf("Equal\nExpected false but got true")
    }
}

func TestEqual_List2Empty(t *testing.T) {
    list1 := New(1, 2, 3, 4)
    list2 := New[int]()
    if Equal(list1, list2) {
        t.Errorf("Equal\nExpected false but got true")
    }
}

func TestEqual_BothEmpty(t *testing.T) {
    list1 := New[string]()
    list2 := New[string]()
    if !Equal(list1, list2) {
        t.Errorf("Equal\nExpected true but got false")
    }
}

func TestFilter_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := Filter(list, func(n int) bool { return n%2 == 0 })
    expected := []int{2, 4, 6}
    if result := ToSlice(filtered); !reflect.DeepEqual(result, expected) {
        t.Errorf("Filter\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, filtered)
}

func TestFilter_NoMatch(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := Filter(list, func(n int) bool { return n > 10 })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("Filter\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, filtered)
}

func TestFilter_EmptyList(t *testing.T) {
    list := New[int]()
    filtered := Filter(list, func(n int) bool { return n%2 != 0 })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("Filter\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, filtered)
}

func TestFilterMap_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := FilterMap(list, func(n int) (bool, int) {
        return n % 2 != 0, n * n
    })
    expected := []int{1, 9, 25}
    if result := ToSlice(filtered); !reflect.DeepEqual(result, expected) {
        t.Errorf("FilterMap\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, filtered)
}

func TestFilterMap_RemoveAll(t *testing.T) {
    list := New(1,2,3)
    filtered := FilterMap(list, func(n int)(bool,int){
        return false, n
    })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("FilterMap\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, filtered)
}

func TestFilterMap_EmptyList(t *testing.T) {
    list := New[int]()
    filtered := FilterMap(list, func(n int) (bool, int) {
        return n % 2 == 0, n * 2
    })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("FilterMap\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, filtered)
}

func TestFind_Found(t *testing.T) {
    list := New(1, 2, 3, 4)
    result := Find(list, 3)
    expected := 2
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFind_NotFound(t *testing.T) {
    list := New(1, 2, 3, 4)
    result := Find(list, 5)
    expected := -1
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFind_EmptyList(t *testing.T) {
    list := New[string]()
    result := Find(list, "A")
    expected := -1
    if result != expected {
        t.Errorf("Find\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldl_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    result := Foldl(list, 0, func(n, s int) int { return n + s })
    expected := 15
    if result != expected {
        t.Errorf("Foldl\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldl_EmptyList(t *testing.T) {
    list := New[int]()
    result := Foldl(list, 0, func(n, s int) int { return n + s })
    expected := 0
    if result != expected {
        t.Errorf("Foldl\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldr_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    result := Foldr(list, 1, func(n, s int) int { return n * s })
    expected := 120
    if result != expected {
        t.Errorf("Foldr\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFoldr_EmptyList(t *testing.T) {
    list := New[int]()
    result := Foldr(list, 1, func(n, s int) int { return n * s })
    expected := 1
    if result != expected {
        t.Errorf("Foldr\nresult: %v\nexpected: %v", result, expected)
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

func TestInsertAt_NormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")

    inserted1 := InsertAt(list, -2, "X")
    expected1 := []string{"a", "b", "X", "c", "d"}
    if result1 := ToSlice(inserted1); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("InsertAt\nresult: %v\nexpected: %v", result1, expected1)
    }
    assertCircular(t, inserted1)

    inserted2 := InsertAt(list, 4, "X")
    expected2 := []string{"a", "b", "c", "d", "X"}
    if result2 := ToSlice(inserted2); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("InsertAt\nresult: %v\nexpected: %v", result2, expected2)
    }
    assertCircular(t, inserted2)
}

func TestInsertAt_PositiveIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("InsertAt\nExpect panic")
        } else if r != "InsertAt, index is out of bound!" {
            t.Errorf("InsertAt\nWrong panic message")
        }
    }()
    InsertAt(New(1, 2, 3, 4), 10, 0)
}

func TestInsertAt_NegativeIndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("InsertAt\nExpect panic")
        } else if r != "InsertAt, index is out of bound!" {
            t.Errorf("InsertAt\nWrong panic message")
        }
    }()
    InsertAt(New(1, 2, 3, 4), -5, 0)
}

func TestInsertAt_EmptyList(t *testing.T) {
    list := New[string]()

    inserted1 := InsertAt(list, 0, "X")
    expected1 := []string{"X"}
    if result1 := ToSlice(inserted1); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("InsertAt\nresult: %v\nexpected: %v", result1, expected1)
    }
    assertCircular(t, inserted1)
}

func TestInsertAt_EmptyList_IndexOutOfBound(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("InsertAt\nExpect panic")
        } else if r != "InsertAt, index is out of bound!" {
            t.Errorf("InsertAt\nWrong panic message")
        }
    }()
    InsertAt(New[int](), 10, 0)
}

func TestJoin_NormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")
    joined := Join(list, "X")
    expected := []string{"a", "X", "b", "X", "c", "X", "d"}
    if result := ToSlice(joined); !reflect.DeepEqual(result, expected) {
        t.Errorf("Join\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, joined)
}

func TestJoin_Singleton(t *testing.T) {
    list := New("a")
    joined := Join(list, "X")
    expected := []string{"a"}
    if result := ToSlice(joined); !reflect.DeepEqual(result, expected) {
        t.Errorf("Join\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, joined)
}

func TestJoin_EmptyList(t *testing.T) {
    list := New[string]()
    joined := Join(list, "X")
    if result := ToSlice(joined); len(result) != 0 {
        t.Errorf("Join\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, joined)
}

func TestMap(t *testing.T) {
    list := New(1, 2, 3, 4)
    mapped := Map(list, func(n int) int { return n * n })
    expected := []int{1, 4, 9, 16}
    if result := ToSlice(mapped); !reflect.DeepEqual(result, expected) {
        t.Errorf("Map\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, mapped)
}

func TestMap_EmptyList(t *testing.T) {
    list := New[int]()
    mapped := Map(list, func(n int) int { return n * n })
    expected := []int{}
    if result := ToSlice(mapped); len(result) != 0 {
        t.Errorf("Map\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, mapped)
}

func TestMapFoldl_MapFoldr_NormalCase(t *testing.T) {
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
    assertCircular(t, mapped1)
    assertCircular(t, mapped2)
}

func TestMapFoldl_MapFoldr_EmptyList(t *testing.T) {
    list := New[int]()
    mapped1, sum := MapFoldl(list, 0, func(n, s int) (int, int) {
        return n * 2, s + n
    })
    mapped2, fac := MapFoldr(list, 1, func(n, f int) (int, int) {
        return n * 2, f * n
    })
    expectedL := []int{}
    expectedS := 0
    expectedF := 1
    if result1 := ToSlice(mapped1); len(result1) != 0 || sum != expectedS {
        t.Errorf("MapFoldl\nresult: %v - %v\nexpected: %v - %v", sum, result1, expectedS, expectedL)
    }
    if result2 := ToSlice(mapped2); len(result2) != 0 || fac != expectedF {
        t.Errorf("MapFoldr\nresult: %v - %v\nexpected: %v - %v", fac, result2, expectedF, expectedL)
    }
    assertCircular(t, mapped1)
    assertCircular(t, mapped2)
}

func TestMax_Min_NormalCase(t *testing.T) {
    list := New("d", "b", "e", "a", "c")

    max := Max(list)
    expectedMax := "e"
    if max.Data != expectedMax {
        t.Errorf("Max\nresult: %v\nexpected: %v", max, expectedMax)
    }

    min := Min(list)
    expectedMin := "a"
    if min.Data != expectedMin {
        t.Errorf("Min\nresult: %v\nexpected: %v", min, expectedMin)
    }
}

func TestMax_Min_EmptyList(t *testing.T) {
    list := New[int]()

    max := Max(list)
    if max != nil {
        t.Errorf("Max\nresult: %v\nexpected: nil", max)
    }

    min := Min(list)
    if min != nil {
        t.Errorf("Min\nresult: %v\nexpected: nil", min)
    }
}

func TestMember_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5)

    member1 := Member(list, 4)
    if !member1 {
        t.Errorf("Member\nresult: %v\nexpected: true", member1)
    }

    member2 := Member(list, 6)
    if member2 {
        t.Errorf("Member\nresult: %v\nexpected: false", member2)
    }
}

func TestMember_EmptyList(t *testing.T) {
    list := New[int]()
    member := Member(list, 4)
    if member {
        t.Errorf("Member\nresult: %v\nexpected: false", member)
    }
}

func TestMerge_UMerge_NormalCase(t *testing.T) {
    list1 := New(2, 8, 6)
    list2 := New(1, 3, 3)
    list3 := New(8, 4, 5)

    merged := Merge(list1, list2, list3)
    expected1 := []int{1, 2, 3, 3, 4, 5, 6, 8, 8}
    if result1 := ToSlice(merged); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("Merge\nresult: %v\nexpected: %v", result1, expected1)
    }
    assertCircular(t, merged)

    umerged := UMerge(list1, list2, list3)
    expected2 := []int{1, 2, 3, 4, 5, 6, 8}
    if result2 := ToSlice(umerged); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("UMerge\nresult: %v\nexpected: %v", result2, expected2)
    }
    assertCircular(t, umerged)
}

func TestMerge_UMerge_EmptyList(t *testing.T) {
    list1 := New[int]()
    list2 := New[int]()
    list3 := New[int]()

    merged := Concat(list1, list2, list3)
    if result := ToSlice(merged); len(result) != 0 {
        t.Errorf("Merge\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, merged)

    umerged := Concat(list1, list2, list3)
    if result := ToSlice(umerged); len(result) != 0 {
        t.Errorf("UMerge\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, umerged)
}

func TestMerge_UMerge_MixList(t *testing.T) {
    list1 := New(2, 8, 6)
    list2 := New[int]()
    list3 := New(8, 4, 5)

    merged := Merge(list1, list2, list3)
    expected1 := []int{2, 4, 5, 6, 8, 8}
    if result1 := ToSlice(merged); !reflect.DeepEqual(result1, expected1) {
        t.Errorf("Merge\nresult: %v\nexpected: %v", result1, expected1)
    }
    assertCircular(t, merged)

    umerged := UMerge(list1, list2, list3)
    expected2 := []int{2, 4, 5, 6, 8}
    if result2 := ToSlice(umerged); !reflect.DeepEqual(result2, expected2) {
        t.Errorf("UMerge\nresult: %v\nexpected: %v", result2, expected2)
    }
    assertCircular(t, umerged)
}

func TestReverse_NormalCase(t *testing.T) {
    list := New(1,2,3,4)
    reversed := Reverse(list)
    expected := []int{4,3,2,1}
    if result := ToSlice(reversed); !reflect.DeepEqual(result, expected) {
        t.Errorf("Reverse\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, reversed)
}

func TestReverse_Singleton(t *testing.T) {
    list := New(1)
    reversed := Reverse(list)
    expected := []int{1}
    if result := ToSlice(reversed); !reflect.DeepEqual(result, expected) {
        t.Errorf("Reverse\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, reversed)
}

func TestReverse_EmptyList(t *testing.T) {
    list := New[int]()
    reversed := Reverse(list)
    expected := []int{}
    if result := ToSlice(reversed); len(result) != 0 {
        t.Errorf("Reverse\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, reversed)
}

func TestSort_NormalCase(t *testing.T) {
    list := New(2, 5, 1, 2, 7, 3, 9, 4, 8, 6, 4)
    sorted := Sort(list)
    expected := []int{1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 9}
    if result := ToSlice(sorted); !reflect.DeepEqual(result, expected) {
        t.Errorf("Sort\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, sorted)
}

func TestSort_EmptyList(t *testing.T) {
    sorted := Sort(New[int]())
    if result := ToSlice(sorted); len(result) != 0 {
        t.Errorf("Sort\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, sorted)
}

func TestUSort_NormalCase(t *testing.T) {
    list := New(2, 5, 1, 2, 7, 3, 9, 4, 8, 6, 4)
    sorted := USort(list)
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if result := ToSlice(sorted); !reflect.DeepEqual(result, expected) {
        t.Errorf("USort\nresult: %v\nexpected: %v", result, expected)
    }
    assertCircular(t, sorted)
}

func TestUSort_EmptyList(t *testing.T) {
    sorted := USort(New[int]())
    if result := ToSlice(sorted); len(result) != 0 {
        t.Errorf("USort\nresult: %v\nexpected: []", result)
    }
    assertCircular(t, sorted)
}

/*
 *******************************************************************************
 * Internal functions and methods
 *******************************************************************************
 */

func assertCircular[T any](t *testing.T, list GoListC[T]) {
    t.Helper()

    if list.Head == nil {
        if list.Tail != nil {
            t.Fatal("empty list: Tail should be nil")
        }
        return
    }

    if list.Tail == nil {
        t.Fatal("non-empty list: Tail is nil")
    }

    node := list.Head
    for {
        if node.Next == list.Head {
            if node != list.Tail {
                t.Fatal("Tail does not reference actual last node")
            }
            break
        }
        node = node.Next
        if node == nil {
            t.Fatal("found nil link in circular list")
        }
    }
}
