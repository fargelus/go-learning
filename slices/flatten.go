package main

import "fmt"

func flatten(matrix [][]int) (vector []int) {
    for _, row := range matrix {
        for _, el := range row {
            vector = append(vector, el)
        }
    }

    return vector
}

func main() {
    matrix := [][]int{{1, 2, 3, 4},
      {5, 6, 7, 8},
      {9, 10, 11},
      {12, 13, 14, 15},
      {16, 17, 18, 19, 20}}
    fmt.Println(flatten(matrix))
}
