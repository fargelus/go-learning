package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println("-------")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!!")
	default:
		fmt.Println("It's weekday:(")
	}
	fmt.Println("-------")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	fmt.Println("-------")

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("It's logic")
		case int:
			fmt.Println("It's integer")
		case string:
			fmt.Println("It's string")
		default:
			fmt.Println("I don't know what it is")
		}
	}
	whatAmI(true)
	whatAmI("123")
	whatAmI(1.2)
}
