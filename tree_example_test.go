package rbtree_test

import (
	"fmt"
	"strings"

	"github.com/sirkon/rbtree"
)

func ExampleTree_Iter() {
	t := rbtree.New[*cmpType]()

	t.Insert(&cmpType{
		name:  "true",
		value: true,
	})
	t.Insert(&cmpType{
		name:  "2",
		value: uint64(2),
	})
	t.Insert(&cmpType{
		name:  "1",
		value: "one",
	})

	for v := range t.Iter() {
		fmt.Println(v.name, v.value)
	}

	// Output:
	// 1 one
	// 2 2
	// true true
}

type cmpType struct {
	name  string
	value any
}

func (c *cmpType) Cmp(v *cmpType) int {
	return strings.Compare(c.name, v.name)
}
