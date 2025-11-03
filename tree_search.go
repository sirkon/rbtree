package rbtree

func (t *Tree[T]) Search(v T) T {
	n := t.root
	for {
		if n == nil {
			var t T
			return t
		}

		cmp := n.value.Cmp(v)
		if cmp < 0 {
			n = n.right
			continue
		}

		if cmp > 1 {
			n = n.left
			continue
		}

		return n.value
	}
}
