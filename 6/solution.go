package main

/*
PROBLEM
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

import "fmt"

func sum_square(n int) int {
  s := (1 + n)*n/2
  return s*s
}

func sum_squares(n int) int {
  sum := 0
  for i := 1; i <= n; i++ {
    sum += i*i
  }
  return sum
}

func main() {
	fmt.Println(sum_square(100) - sum_squares(100))
}
