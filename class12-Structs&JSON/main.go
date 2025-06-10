package main

import (
	"fmt"
	//"iter"
	//"golang.org/x/exp/iter"
	"container/list"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	fmt.Println("Structs & JSON")
	q := list.New()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	fmt.Println(q)
	q.PushBack(1)
	fmt.Println(q.Back().Prev().Value)
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	front := q.Front()
	q.Remove(front)

	fmt.Println(q)

	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("\\\\")

	for k := q.Back(); k != q.Front(); k = k.Prev() {
		fmt.Println(k.Value)
	}
	fmt.Println("adding front: ", q.Front().Value)

	tree := treemap.NewWithIntComparator()

	tree.Put(3, "C")
	tree.Put(1, "A")
	tree.Put(2, "B")

	tree.Each(func(k, v interface{}) {
		fmt.Printf("%v: %v\n", k, v)
	})

	kl := []int{1, 5, 6, 8, 2, 7, 3, 9, 26}

	set := treeset.NewWithIntComparator()

	for _, v := range kl {
		set.Add(v)
	}

	// duplicate, will not be added again
	fmt.Println("cbakjcbakjcbkaj")
	fmt.Println(set.Values()...)
	set.Each(func(_ int, value interface{}) {
		fmt.Println(value)
	})

}
