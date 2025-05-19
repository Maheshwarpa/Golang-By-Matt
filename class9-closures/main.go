package main

import (
	"fmt"
)

func main() {
	fmt.Println("Closures")
	k := kid()
	fmt.Println(k())

	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p", i, &i)
		}
		do(v)
	}

	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Println("I is :", i)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}

}

func do(m func()) {
	m()
}

func kid() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		return j
	}
}
