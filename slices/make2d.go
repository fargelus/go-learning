package main

import "fmt"

func matrixSize(vector []int, rows int) int {
    size := len(vector) / rows
    if len(vector) % rows != 0 {
        size++
    }
    
    return size
}

func make2D(vector []int, rows int) (int, [][]int) {
    matrix := make([][]int, 0, matrixSize(vector, rows))
    row := make([]int, rows)
    for i, elem := range vector {
        row[i % rows] = elem
        if i % rows == rows - 1 || len(vector) - 1 == i {
            matrix = append(matrix, row)
            row = make([]int, rows)
        }
    }

    return rows, matrix
}

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
      11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

    fmt.Println(make2D(slice, 3))
    fmt.Println(make2D(slice, 4))
    fmt.Println(make2D(slice, 5))
    fmt.Println(make2D(slice, 6))
    fmt.Println(make2D(slice, 20))
}
