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