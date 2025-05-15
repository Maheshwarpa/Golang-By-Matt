package main

import (
	"fmt"
)

func main() {
	fmt.Println("class-05")

	var arr = [...]int{1, 2, 3}
	sl := []int{0, 0, 0}
	fmt.Printf("%T %T\n", arr, sl)
	y := Do(arr, sl)
	fmt.Println(arr, sl, y)

	var m map[string]int
	mp := make(map[string]int)
	mp["king"] = 1
	m = mp
	fmt.Println(m, len(m))

	readfile()

}

func Do(a [3]int, b []int) []int {
	a[0] = 4
	b[0] = 3
	b = append(b, 3, 4, 5)

	c := make([]int, 9)

	c[1] = 42
	copy(c, b)
	fmt.Println(b)
	return c
}
