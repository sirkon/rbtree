package rbtree

import (
	"bytes"
	"fmt"
)

// show show tree in a dot format.
func (t *Tree[T]) show() []byte {
	var buf bytes.Buffer

	buf.WriteString("graph rbtree {\n")
	it := t.Iter()
	for it.Next() {
		color := "black"
		if it.n.isRed() {
			color = "red"
		}

		_, _ = fmt.Fprintf(&buf, " %v [color=%s];\n", it.Item(), color)
	}

	it = t.Iter()
	for it.Next() {
		n := it.n
		if n.parent != nil {
			_, _ = fmt.Fprintf(&buf, " %v -- %v;\n", n.parent.value, it.Item())
		}
	}

	buf.WriteString("}\n")

	return buf.Bytes()
}
