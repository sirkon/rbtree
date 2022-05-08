package rbtree

// Min gets the least tree element.
func (t *Tree[T]) Min() (val T, exists bool) {
	if t.root == nil {
		return val, false
	}

	cur := t.root
	for cur.left != nil {
		cur = cur.left
	}

	return cur.value, true
}
