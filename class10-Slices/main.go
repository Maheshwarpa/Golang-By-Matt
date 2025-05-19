package main

import (
	"fmt"
)

func main() {
	var s []int            // nil slice
	t := []int{}           // non-nil slice
	u := make([]int, 5)    // Slice with length 5 and all initialized to zero
	v := make([]int, 0, 5) // slice with capacity 5, data may or may not be initialzied to zero

	fmt.Printf("%T -  %[1]s - %s\n", s, s == nil)
	fmt.Println(t, t == nil)
	fmt.Println(u, u == nil)
	fmt.Println(v, v == nil)

	k := [3]int{1, 2, 3}
	l := k[0:1]
	m := l[0:2]
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Printf("%p %[1]v\n", &k)
	fmt.Printf("%p %[1]v\n", l)
	fmt.Printf("%p %[1]v\n", m)
	n := m[0:1:1]
	fmt.Printf("%p %[1]v\n", n)

	n = append(n, 5)
	fmt.Printf("%p", n)

}
