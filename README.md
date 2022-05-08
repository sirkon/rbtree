# RBTree

Generic implementation of red-black trees.

## Installation

```shell
go get github.com/sirkon/rbtree
```

## Usage example.

```go
package main

import (
	"fmt"
	"github.com/sirkon/rbtree"
)

func main() {
	t := rbtree.NewOrdered[int]()
	t.Insert(10)
	t.Insert(9)
	t.Insert(5)
	t.Insert(7)
	t.Insert(1)
	t.Delete(7)

	it := t.Iter()
	for it.Next() {
		fmt.Println(it.Item())
	}

	// Output
	// 1
	// 5
	// 9
	// 10
}
```

## Notice.

Remember, this is not `map[K, V]` analogue, no explicit values bounded to their keys support here.
Yet it is possible to emulate them:

```go
package main

import (
	"fmt"
	"github.com/sirkon/rbtree"
)

type KeyValue struct {
	Key string
	Val string
}

// Cmp to implement rbtree.TreeKey
func (kv KeyValue) Cmp(elem KeyValue) int {
	switch {
	case kv.Key < elem.Key:
		return -1
	case kv.Key == elem.Key:
		return 0
	default:
		return 1
	}
}

func main() {
	t := rbtree.New[KeyValue]()
	t.Insert(KeyValue{Key: "1", Val: "a"})
	t.Insert(KeyValue{Key: "2", Val: "b"})

	it := t.Iter()
	for it.Next() {
		fmt.Println(it.Item().Key, "-", it.Item().Val)
	}

	// Output:
	// 1 - a
	// 2 - b
}
```