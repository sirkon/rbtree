package rbtree

import (
	"bytes"
	"fmt"
)

// show shows the tree in a dot format.
func (t *Tree[T]) show() []byte {
	var buf bytes.Buffer

	buf.WriteString("graph rbtree {\n")
	it := &Iterator[T]{
		r: t.root,
	}
	for it.Next() {
		color := "black"
		if it.n.isRed() {
			color = "red"
		}

		_, _ = fmt.Fprintf(&buf, " %v [color=%s];\n", it.Item(), color)
	}

	it = &Iterator[T]{
		r: t.root,
	}
	for it.Next() {
		n := it.n
		if n.parent != nil {
			_, _ = fmt.Fprintf(&buf, " %v -- %v;\n", n.parent.value, it.Item())
		}
	}

	buf.WriteString("}\n")

	return buf.Bytes()
}
