package main

/*
PROBLEM

Starting in the top left corner of a 2×2 grid, and only being able to move to
the right and down, there are exactly 6 routes to the bottom right corner.


How many such routes are there through a 20×20 grid?

*/

import "fmt"

func main() {
  dim := 20 + 1
  matrix := make([][]int, dim)
  for i := 0; i < dim; i++ {
    matrix[i] = make([]int, dim)
  }
  for i := 0; i < dim; i++ {
    matrix[0][i] = 1
    matrix[i][0] = 1
  }
  for i := 1; i < dim; i++ {
    row := matrix[i]
    for j := 1; j < dim; j++ {
      row[j] = matrix[i-1][j] + matrix[i][j-1]
    }
  }
	fmt.Println(matrix[dim-1][dim-1])
}
