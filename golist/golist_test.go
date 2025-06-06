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
        t.Errorf("Delete\nresult: %v", result)
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
        t.Errorf("Delete\nresult: %v", result)
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

func TestDropWhile(t *testing.T) {
    list := New(1, 2, 3, 4, 5)
    droped := DropWhile(list, func(n int) bool { return n < 4 })
    expected := []int{4, 5}
    if result := ToSlice(droped); !reflect.DeepEqual(result, expected) {
        t.Errorf("DropWhile\nresult: %v\nexpected: %v", result, expected)
    }
}
