// Package constraints contains custom type contraints in go-linkedlist.
package constraints

type Ordered interface {
    Numeric | ~string
}

type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64
}