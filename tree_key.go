package rbtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// TreeKey tree keys must satisfy this interface.
type TreeKey[T treeElementReferencer[T]] interface {
	Cmp(elem T) int
}

type treeElementReferencer[T any] interface {
	Cmp(elem T) int
}

// treeKeyOrdered TreeKey implementation for ordered types.
type treeKeyOrdered[T constraints.Ordered] struct {
	v T
}

// Cmp to implement TreeKey.
func (k treeKeyOrdered[T]) Cmp(elem treeKeyOrdered[T]) int {
	if k.v < elem.v {
		return -1
	}

	if k.v > elem.v {
		return 1
	}

	return 0
}

func (k treeKeyOrdered[T]) String() string {
	return fmt.Sprintf("%v", k.v)
}
