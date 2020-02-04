package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Input two numbers here")

	var first, second int
	fmt.Scanln(&first, &second)
	fmt.Println(sum(first, second))
}
