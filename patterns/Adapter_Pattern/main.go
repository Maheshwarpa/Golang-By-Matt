package main

import (
	ap "adapter/Apple"
	ac "adapter/Client"
	ad "adapter/android"
	"fmt"
)

func main() {
	fmt.Println("Starting!...")
	apple := &ap.Apple{}
	android := &ad.Android{}
	client := &ac.Client{}
	client.ChargeMobile(android)
	client.ChargeMobile(apple)
}
