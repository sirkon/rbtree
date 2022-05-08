package rbtree

// Delete tries to delete existing value.
// Returns true when and only when the element exist.
func (t *Tree[T]) Delete(value T) (status bool) {
	n := lookupForValue(t.root, value)
	if n == nil {
		return false
	}

	p := n.parent
	defer func() {
		n.parent = nil
		n.left = nil
		n.right = nil
		n = nil
		p = nil

		if status {
			t.size--
		}
	}()

	isRight := n.isRight()

	if n.left != nil && n.right != nil {
		c := n.right
		for c.left != nil {
			c = c.left
		}

		n.value = c.value
		n = c
		p = c.parent
		isRight = n.isRight()
	}

	switch {
	case n.left == nil && n.right == nil:

		if !swapChild(p, n, nil) {
			t.root = nil
			return
		}

		if n.red {
			// Удаление красной вершины не влияет на целостность структуры
			// красно-чёрных деревьев.
			return
		}

	case n.left == nil:
		if !n.red {
			n.right.red = false
		}

		n.right.parent = p
		if !swapChild(p, n, n.right) {
			t.root = n.right
			if n.right != nil {
				n.right.red = false
			}
		}

		return

	case n.right == nil:
		if !n.red {
			n.left.red = false
		}

		n.left.parent = p
		if !swapChild(p, n, n.left) {
			t.root = n.left
			if n.left != nil {
				n.left.red = false
			}
		}

		return
	}

	if p == nil {
		// удалённая вершина была корнем дерева то баланс корректен т.к.
		// структура поддерева не затронута
		return
	}

	t.deleteFix(p, isRight)

	return true
}

func (t *Tree[T]) deleteFix(p *node[T], isRight bool) {
start:

	if p == nil {
		return
	}

	if isRight {
		l := p.left

		switch {
		case p.red:
			switch {
			case l.right.isRed():
				t.lrToP(p)
				p.red = false
				return

			case l.left.isRed():
				t.lToP(p)
				return

			default:
				l.red = true
				p.red = false
				return
			}

		default:
			switch {
			case l.isRed():
				lr := l.right
				switch {
				case lr.left.isRed():
					t.lrToP(p)
					l.right.red = false
					return

				case lr.right.isRed():
					t.lrToP(p)
					p = l
					isRight = true
					goto start

				default:
					t.lToP(p)
					p.red = false
					l.red = false
					p.left.red = true
					return
				}

			default:
				switch {
				case l.right.isRed():
					l.right.red = false
					t.lrToP(p)
					return
				case l.left.isRed():
					l.left.red = false
					t.lToP(p)
					return
				default:
					t.lToP(p)
					p.red = true
					if p.parent == nil {
						// дошли до конца, дальнейшая перебалансировка не нужна
						return
					}
					isRight = l.isRight()
					p = l.parent
					goto start
				}
			}
		}
	}

	r := p.right
	switch {
	case p.red:
		switch {
		case r.left.isRed():
			t.rlToP(p)
			p.red = false
			return

		case r.right.isRed():
			t.rToP(p)
			return

		default:
			r.red = true
			p.red = false
			return
		}

	default:
		switch {
		case r.isRed():
			rl := r.left

			switch {
			case rl.right.isRed():
				t.rlToP(p)
				r.left.red = false
				return

			case rl.left.isRed():
				t.rlToP(p)
				p = r
				isRight = false
				goto start

			default:
				t.rToP(p)
				p.red = false
				r.red = false
				p.right.red = true
				return
			}

		default:
			switch {
			case r.left.isRed():
				r.left.red = false
				t.rlToP(p)
				return
			case r.right.isRed():
				r.right.red = false
				t.rToP(p)
				return
			default:
				t.rToP(p)
				p.red = true
				if p.parent == nil {
					// дошли до конца, дальнейшая перебалансировка не нужна
					return
				}
				isRight = r.isRight()
				p = r.parent
				goto start
			}
		}
	}
}
