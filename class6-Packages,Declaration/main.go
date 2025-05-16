package main

import (
	"fmt"
)

type nm int

var (
	x, y int
	k    float64
)

func main() {
	fmt.Println("Packages")
	var nb int
	var x nm
	x = 5
	fmt.Printf("%T [1]%v", x)
	fmt.Println(nb)

outer:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if j == 2 {
				break outer
			}
		}
	}

	// Switch

	switch k {
	case 1.0, 2.0, 3.0, 4.0:
		fmt.Println("1-4")
	case 0.0:
		fmt.Println("Zero")
	default:
		fmt.Println("Default")
	}

	//switch with true

	switch {
	case y == nb+1:
		fmt.Println("New One")
	}

	// Map and its iteration
	mp := make(map[string]int)

	mp["str"] = 1
	mp["ktr"] = 2

	for k := range mp {
		fmt.Println(k)
	}

	for k, v := range mp {
		fmt.Println(k, " : ", v)
	}

}
