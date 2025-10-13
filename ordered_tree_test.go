package rbtree_test

import (
	"fmt"

	"github.com/sirkon/rbtree"
)

func ExampleTreeOrdered_Iter() {
	tree := rbtree.NewOrdered[int]()
	tree.Insert(1)
	tree.Insert(15)
	tree.Insert(14)
	tree.Insert(2)

	for v := range tree.Iter() {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 14
	// 15
}
