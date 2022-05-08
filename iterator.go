package rbtree

// TreeIterator an abstraction for iterators over red-black tree.
type TreeIterator[T any] interface {
	Next() bool
	Item() T
}

// Iterator an iterator over red-black tree nodes.
// Beware, the tree must not be changed during the iteration.
type Iterator[T TreeKey[T]] struct {
	n *node[T]
	r *node[T]

	justRaised bool
}

// Next returns true if some nodes left.
func (i *Iterator[T]) Next() bool {
	if i.n == nil && i.r == nil {
		return false
	}

	if i.n == nil {
		i.n = i.r
		for i.n.left != nil {
			i.n = i.n.left
		}

		return true
	}

	if i.n.right != nil {
		i.n = i.n.right
		for i.n.left != nil {
			i.n = i.n.left
		}
		return true
	}

	for i.n.isRight() {
		i.n = i.n.parent
	}

	i.n = i.n.parent

	if i.n == nil {
		i.n = nil
		i.r = nil
		return false
	}

	return true
}

// Item returns node value.
func (i *Iterator[T]) Item() T {
	return i.n.value
}
