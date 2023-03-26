package main

import "fmt"

func uniqInts(slice []int) (uniq []int) {
    dict := make(map[int]int)
    for _, el := range slice {
        dict[el]++
        if dict[el] == 1 {
            uniq = append(uniq, el)
        }
    }

    return uniq
}

func main() {
    slice := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
    fmt.Println(uniqInts(slice))
}
