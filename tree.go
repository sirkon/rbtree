package rbtree

// New creates new empty red-black tree.
func New[T TreeKey[T]]() *Tree[T] {
	return &Tree[T]{}
}

// Tree definition of red-black tree.
//
// TODO arena for element allocation is definitely a viable
//      idea.
type Tree[T TreeKey[T]] struct {
	root *node[T]
	size int
}

// Iter returns reb-black tree iterator.
func (t *Tree[T]) Iter() *Iterator[T] {
	return &Iterator[T]{
		r: t.root,
	}
}

// Len returns tree length.
func (t *Tree[T]) Len() int {
	return t.size
}

func swapChild[T TreeKey[T]](parent, from, to *node[T]) bool {
	if parent == nil {
		return false
	}

	if parent.left == from {
		parent.left = to
	} else {
		parent.right = to
	}

	return true
}

func (t *Tree[T]) lToP(p *node[T]) {
	l := p.left

	l.parent, l.right, p.parent, p.left = p.parent, p, l, l.right

	if !swapChild(l.parent, p, l) {
		t.root = l
		l.red = false
	}
	if p.left != nil {
		p.left.parent = p
	}
}

func (t *Tree[T]) lrToP(p *node[T]) {
	l := p.left
	lr := p.left.right

	lr.parent, lr.left, lr.right, l.parent, l.right, p.parent, p.left =
		p.parent, l, p, lr, lr.left, lr, lr.right

	if !swapChild(lr.parent, p, lr) {
		t.root = lr
		lr.red = false
	}
	if p.left != nil {
		p.left.parent = p
	}
	if l.right != nil {
		l.right.parent = l
	}
}

func (t *Tree[T]) rToP(p *node[T]) {
	r := p.right

	r.parent, r.left, p.parent, p.right = p.parent, p, r, r.left

	if !swapChild(r.parent, p, r) {
		t.root = r
		r.red = false
	}
	if p.right != nil {
		p.right.parent = p
	}
}

func (t *Tree[T]) rlToP(p *node[T]) {
	r := p.right
	rl := p.right.left

	rl.parent, rl.left, rl.right, r.parent, r.left, p.parent, p.right =
		p.parent, p, r, rl, rl.right, rl, rl.left

	if !swapChild(rl.parent, p, rl) {
		t.root = rl
		rl.red = false
	}
	if p.right != nil {
		p.right.parent = p
	}
	if r.left != nil {
		r.left.parent = r
	}
}
