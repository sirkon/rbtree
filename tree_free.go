package rbtree

// Free "frees" the tree to decrease GC pressure.
func (t *Tree[T]) Free() {
	n := t.root

	for n != nil {
		// look for a node not having any children
	loop:
		for {
			switch {
			case n.left != nil:
				n = n.left
			case n.right != nil:
				n = n.right
			default:
				break loop
			}
		}

		// clean the node value with its 0
		var v T
		n.value = v

		x := n.parent
		if x != nil {
			if n.isRight() {
				x.right = nil
			} else {
				x.left = nil
			}
			n.parent = nil
		}

		n = x
	}

	t.root = nil
}
