package rbtree

// Add adds new or replaces existing element.
// Returns true when and only when the element
// didn't exist before.
func (t *Tree[T]) Add(v T) (fresh bool) {
	if t.root == nil {
		t.root = &node[T]{
			value: v,
		}
		t.size = 1
		return true
	}

	p, isRight, alreadyExist := lookupForFreeValueParent(t.root, v)
	if alreadyExist {
		if isRight {
			p.right.value = v
		} else {
			p.left.value = v
		}
		return false
	}

	n := &node[T]{
		value:  v,
		parent: p,
		red:    true,
	}
	if isRight {
		p.right = n
	} else {
		p.left = n
	}

	t.size++
	return t.rebalanceInserted(p, n)
}

// Insert tries to insert a new value.
// Returns true when and only when the tree
// actually had an element.
func (t *Tree[T]) Insert(v T) (existBefore bool) {
	if t.root == nil {
		t.root = &node[T]{
			value: v,
		}
		t.size = 1
		return true
	}

	p, isRight, alreadyExist := lookupForFreeValueParent(t.root, v)
	if alreadyExist {
		return false
	}

	n := &node[T]{
		value:  v,
		parent: p,
		red:    true,
	}
	if isRight {
		p.right = n
	} else {
		p.left = n
	}

	t.size++
	return t.rebalanceInserted(p, n)
}

// InsertReturn tries to insert an element.
// Returns inserted element if it didn't exist.
// Returns existing element otherwise.
func (t *Tree[T]) InsertReturn(v T) T {
	if t.root == nil {
		t.root = &node[T]{
			value: v,
		}
		t.size = 1
		return v
	}

	p, isRight, alreadyExist := lookupForFreeValueParent(t.root, v)
	if alreadyExist {
		return p.value
	}

	n := &node[T]{
		value:  v,
		parent: p,
		red:    true,
	}
	if isRight {
		p.right = n
	} else {
		p.left = n
	}

	t.size++
	t.rebalanceInserted(p, n)
	return v
}

func (t *Tree[T]) rebalanceInserted(p *node[T], n *node[T]) bool {
	for p.red {
		g, u := p.olderRelatives()
		if g == nil {
			p.red = false
			return true
		}

		if u.isRed() {
			if u != nil {
				u.red = false
			}
			p.red = false
			g.red = true

			p = g.parent
			n = g
			if p == nil {
				g.red = false
				return true
			}
			continue
		}

		switch {
		case !p.isRight() && !n.isRight():
			p.red = false
			g.red = true
			t.lToP(g)

		case !p.isRight() && n.isRight():
			n.red = false
			g.red = true
			t.lrToP(g)

		case p.isRight() && !n.isRight():
			n.red = false
			g.red = true
			t.rlToP(g)

		case p.isRight() && n.isRight():
			p.red = false
			g.red = true
			t.rToP(g)
		}

		return true
	}

	return true
}
