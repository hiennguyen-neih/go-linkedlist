package golist2

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
    expected := "[ 0.1 <-> 0.2 <-> 0.3 <-> 0.4 <-> 0.5 ]"
    if result := list.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestGoListString_String(t *testing.T) {
    list := New("A", "B", "C", "D", "E", "F")
    expected := `[ "A" <-> "B" <-> "C" <-> "D" <-> "E" <-> "F" ]`
    if result := list.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}
