package node

import (
    "testing"
)

func TestNodeString_Int(t *testing.T) {
    node := Node[int]{Data: 8}
    expected := "8"
    if result := node.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}

func TestNodeString_String(t *testing.T) {
    node := Node[string]{Data: "X"}
    expected := `"X"`
    if result := node.String(); result != expected {
        t.Errorf("String\nresult: %v\nexpected: %v", result, expected)
    }
}
