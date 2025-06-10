package main

import (
	"fmt"
	"sort"
)

func main() {
	primes := []int{19, 5, 7, 11, 13, 17, 3}
	sort.Ints(primes)
	fmt.Println(primes)
}
