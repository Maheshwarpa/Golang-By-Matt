package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main")
	s := "Maheshwar"
	mp := make(map[string]string)
	mp["name"] = "Maheshwar"
	mp["age"] = "15"
	k := 15

	fmt.Printf("%d\n", k)
	fmt.Printf("%q\n", k)
	fmt.Printf("%T\n", k)
	fmt.Printf("%v\n", k)
	fmt.Printf("%#v\n", k)
	fmt.Println("String")
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Println("Maps")
	fmt.Printf("%s\n", mp)
	fmt.Printf("%q\n", mp)
	fmt.Printf("%T\n", mp)
	fmt.Printf("%v\n", mp)
	fmt.Printf("%#v\n", mp)

}
