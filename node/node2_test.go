package node

import (
    "testing"
)

func TestNode2String_Int(t *testing.T) {
    node := Node2[int]{Data: 8}
    expected := "8"
    if result := node.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestNode2String_String(t *testing.T) {
    node := Node2[string]{Data: "X"}
    expected := `"X"`
    if result := node.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}
