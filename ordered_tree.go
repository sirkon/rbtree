package rbtree

import "golang.org/x/exp/constraints"

// NewOrdered constructs a red-black tree with ordered elements.
func NewOrdered[T constraints.Ordered]() *TreeOrdered[T] {
	return &TreeOrdered[T]{
		t: New[treeKeyOrdered[T]](),
	}
}

// TreeOrdered a red-black tree with ordered values.
type TreeOrdered[T constraints.Ordered] struct {
	t *Tree[treeKeyOrdered[T]]
}

// Iter returns tree iterator.
func (t TreeOrdered[T]) Iter() OrderedIterator[T] {
	return OrderedIterator[T]{
		iter: t.t.Iter(),
	}
}

// Insert tries to insert a non-existing element. Returns
// true if the element didn't exist indeed.
func (t TreeOrdered[T]) Insert(n T) (status bool) {
	return t.t.Insert(treeKeyOrdered[T]{v: n})
}

// Delete tries to delete existing element. Returns true
// if the element existed.
func (t TreeOrdered[T]) Delete(n T) (status bool) {
	return t.t.Delete(treeKeyOrdered[T]{v: n})
}

// OrderedIterator iterator definition over red-black tree of
// ordereds.
type OrderedIterator[T constraints.Ordered] struct {
	iter *Iterator[treeKeyOrdered[T]]
}

// Next to implement TreeIterator
func (i OrderedIterator[T]) Next() bool {
	return i.iter.Next()
}

// Item to implement TreeIterator
func (i OrderedIterator[T]) Item() T {
	return i.iter.Item().v
}
