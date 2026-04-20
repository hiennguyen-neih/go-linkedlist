package golistc

import (
    "testing"
    "reflect"
)

func TestNew_ToSlice(t *testing.T) {
    list := New(1, 2, 3, 4)
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(New(...)): %v\nexpected: %v", result, expected)
    }
}

func TestFromSlice_ToSlice(t *testing.T) {
    list := FromSlice([]int{1, 2, 3, 4})
    expected := []int{1, 2, 3, 4}
    if result := ToSlice(list); !reflect.DeepEqual(result, expected) {
        t.Errorf("ToSlice(FromSlice(...)): %v\nexpected: %v", result, expected)
    }
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

func TestAll(t *testing.T) {
    list1 := New(1, 3, 5, 7)
    list2 := New(1, 3, 5, 8)
    if result1 := All(list1, func(n int) bool { return n%2 != 0 }); false {
        t.Errorf("All\nresult1: %v\nexpected: true", result1)
    }
    if result2 := All(list2, func(n int) bool { return n%2 != 0 }); false {
        t.Errorf("All\nresult2: %v\nexpected: false", result2)
    }
}

func TestAny(t *testing.T) {
    list1 := New(2, 4, 6, 7)
    list2 := New(2, 4, 6, 8)
    if result1 := Any(list1, func(n int) bool { return n%2 != 0}); false {
        t.Errorf("Any\nresult1: %v\nexpected: true", result1)
    }
    if result2 := Any(list2, func(n int) bool { return n%2 != 0}); false {
        t.Errorf("Any\nresult2: %v\nexpected: false", result2)
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
    list := New("c", "d")
    appended := AppendHead(list, "a", "b")
    expected := []string{"a", "b", "c", "d"}
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

func TestDelete_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 2, 4)
    deleted := Delete(list, 2)
    expected := []int{1, 3, 2, 4}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("Delete\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDelete_EmptyList(t *testing.T) {
    list := New[int]()
    deleted := Delete(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
}

func TestDeleteAt_NormalCase(t *testing.T) {
    list := New("a", "b", "c", "d")
    deleted := DeleteAt(list, -2)
    expected := []string{"a", "b", "d"}
    if result := ToSlice(deleted); !reflect.DeepEqual(result, expected) {
        t.Errorf("DeleteAt\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestDeleteAt_EmptyList(t *testing.T) {
    list := New[int]()
    deleted := DeleteAt(list, 0)
    if result := ToSlice(deleted); len(result) != 0 {
        t.Errorf("Delete\nresult: %v\nexpected: []", result)
    }
}

func TestDeleteAt_IndexOutOfBound(t *testing.T) {
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

func TestDropWhile_TakeWhile(t *testing.T) {
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

func TestFilter_NormalCase(t *testing.T) {
    list := New(1, 2, 3, 4, 5, 6)
    filtered := Filter(list, func(n int) bool { return n%2 == 0 })
    expected := []int{2, 4, 6}
    if result := ToSlice(filtered); !reflect.DeepEqual(result, expected) {
        t.Errorf("Filter\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestFilter_EmptyList(t *testing.T) {
    list := New[int]()
    filtered := Filter(list, func(n int) bool { return n%2 != 0 })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("Filter\nresult: %v\nexpected: []", result)
    }
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
}

func TestFilterMap_EmptyList(t *testing.T) {
    list := New[int]()
    filtered := FilterMap(list, func(n int) (bool, int) {
        return n % 2 == 0, n * 2
    })
    if result := ToSlice(filtered); len(result) != 0 {
        t.Errorf("FilterMap\nresult: %v\nexpected: []", result)
    }
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

func TestFinc_EmptyList(t *testing.T) {
    list := New[string]()
    result := Find(list, "A")
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
