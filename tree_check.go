package rbtree

import (
	"fmt"
	"strings"
)

func (t *Tree[T]) check() error {
	if _, _, err := checkTreeValidity(t.root); err != nil {
		return err
	}

	return nil
}

func checkTreeValidity[T TreeKey[T]](n *node[T]) (blackDepth int, invalid []*node[T], err error) {
	if n == nil {
		return 1, nil, nil
	}

	if !n.red {
		blackDepth++
	}

	if n.isRed() && n.left.isRed() {
		return 0, []*node[T]{n.left}, fmt.Errorf("nodes '%v' -> '%v' are both red", n, n.left)
	}

	if n.isRed() && n.right.isRed() {
		return 0, []*node[T]{n.left}, fmt.Errorf("nodes '%v' -> '%v' are both red", n, n.right)
	}

	d1, inv, err := checkTreeValidity(n.left)
	if err != nil {
		return 0,
			append([]*node[T]{n}, inv...),
			fmt.Errorf("check left subtree of node %v: %w", n, err)
	}

	d2, inv, err := checkTreeValidity(n.right)
	if err != nil {
		return 0,
			append([]*node[T]{n}, inv...),
			fmt.Errorf("check right subtree of node %v: %w", n, err)
	}

	if d1 != d2 {
		return 0,
			[]*node[T]{n},
			fmt.Errorf("black depth mismatch between left and right subtrees of node %v: %d != %d", n, d1, d2)
	}

	if !n.isRed() {
		d1++
	}

	return d1, nil, nil
}

func leftPath[T TreeKey[T]](n *node[T]) []*node[T] {
	var res []*node[T]
	for n != nil {
		res = append(res, n)
		n = n.left
	}

	res = append(res, nil)

	return nil
}

func showPath[T TreeKey[T]](path []*node[T]) string {
	var buf strings.Builder
	for i, n := range path {
		if i > 0 {
			buf.WriteString("->")
		}
		buf.WriteString(n.String())
	}

	return buf.String()
}
