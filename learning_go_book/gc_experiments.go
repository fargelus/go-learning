package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	start := time.Now()

	persons := make([]Person, 10_000_000, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		persons[i] = Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       22,
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Completion time: %v\n", elapsed)
	fmt.Printf("Completion time: %d msec\n", elapsed.Milliseconds())
	fmt.Printf("Completion time: %.2f sec\n", elapsed.Seconds())
}
