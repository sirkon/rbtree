package rbtree

import "fmt"

type node[T TreeKey[T]] struct {
	value T

	parent *node[T]
	left   *node[T]
	right  *node[T]
	red    bool
}

func (n *node[T]) String() string {
	if n == nil {
		return "nil(black)"
	}

	colorName := "black"
	if n.red {
		colorName = "red"
	}

	res := fmt.Sprintf("%v(%s)", n.value, colorName)
	return res
}

func (n *node[T]) isRight() bool {
	if n.parent == nil {
		return false
	}

	return n.parent.right == n
}

func (n *node[T]) isRed() bool {
	if n == nil {
		return false
	}

	return n.red
}

// olderRelatives retrieve parent (p) and brother (b) of given node.
// The brother is either left child of the parent if the element
// is the right one or, the otherwise, the right child in case if
// the current node is the left one.
func (n *node[T]) olderRelatives() (p *node[T], b *node[T]) {
	p = n.parent

	if p == nil {
		return nil, nil
	}

	if p.right == n {
		return p, p.left
	}

	return p, p.right
}

func lookupForFreeValueParent[T TreeKey[T]](n *node[T], value T) (_ *node[T], isRight bool, existing bool) {
	for {
		cmp := value.Cmp(n.value)
		switch {
		case cmp < 0:
			if n.left == nil {
				return n, false, false
			}
			n = n.left

		case cmp > 0:
			if n.right == nil {
				return n, true, false
			}
			n = n.right

		default:
			return n, false, true
		}
	}
}

func lookupForValue[T TreeKey[T]](n *node[T], value T) *node[T] {
	for {
		cmp := value.Cmp(n.value)
		switch {
		case cmp < 0:
			if n.left == nil {
				return nil
			}
			n = n.left

		case cmp > 0:
			if n.right == nil {
				return nil
			}
			n = n.right

		default:
			return n
		}
	}
}
