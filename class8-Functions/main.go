package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")

	mp := map[int]int{1: 2, 2: 3, 3: 4, 4: 5}
	do(mp)
	fmt.Println("OG map is :", mp)

	dom(&mp)
	fmt.Println("DOG map is:", mp)
}

func do(m map[int]int) {
	defer fmt.Println("Do Map is: ", m)
	m[5] = 6
	m = make(map[int]int)
	m[4] = 4
}

func dom(m *map[int]int) {
	defer fmt.Println("dom is:", *m)
	(*m)[6] = 7
	*m = make(map[int]int)
	(*m)[1] = 0
}
