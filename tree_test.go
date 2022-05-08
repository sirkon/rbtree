package rbtree

import (
	"math/rand"
	"testing"
)

func FuzzInsert(f *testing.F) {
	rand.Seed(0)
	for i := 0; i < 1000; i++ {
		f.Add(i*10 + rand.Intn(10))
	}
	f.Fuzz(func(t *testing.T, l int) {
		data := make([]int, l)
		for i := range data {
			data[i] = i
		}
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})

		r := NewOrdered[int]()
		for _, v := range data {
			r.Insert(v)
		}

		if err := r.t.check(); err != nil {
			t.Errorf("check red black tree after %d insertions: %s", l, err)
		}
	})
}

func FuzzDelete(f *testing.F) {
	rand.Seed(0)
	for i := 0; i < 1000; i++ {
		d := rand.Intn(10)
		if i == 0 {
			for d == 0 {
				d = rand.Intn(10)
			}
		}
		f.Add(i*10 + rand.Intn(10))
	}

	f.Fuzz(func(t *testing.T, l int) {
		data := make([]int, l)
		for i := range data {
			data[i] = i
		}
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})

		r := NewOrdered[int]()
		for _, v := range data {
			r.Insert(v)
		}

		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})

		for i := 0; i < l; i += 2 {
			r.Delete(data[i])
		}

		if err := r.t.check(); err != nil {
			t.Errorf("check red black tree after deletions: %v", err)
		}
	})
}
