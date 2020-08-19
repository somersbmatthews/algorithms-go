package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	makePointers()
}

func makePointers() {
	data := 10
	ptr := &data
	fmt.Println("Value stored at variable var is", data)
	fmt.Println("Value stored at variable var is", *ptr)
	fmt.Println("The address of variable var is", &data)
	fmt.Println("The address of variable var is", ptr)
}
