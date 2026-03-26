package main

import "fmt"

func UpdateSlice(slice []string, addition string) {
	lastPos := len(slice)
	slice[lastPos-1] = addition
	fmt.Println("UpdateSlice:", slice)
}

func GrowSlice(slice []string, addition string) {
	slice = append(slice, addition)
	fmt.Println("GrowSlice:", slice)
}

func main() {
	games := []string{"Life Is Strange 2", "The Darkness 2"}
	UpdateSlice(games, "GTA V")
	GrowSlice(games, "GTA V")
	fmt.Println("After update", games)
}
