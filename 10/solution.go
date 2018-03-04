package main

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

import (
	"fmt"
	"math"
)

func sq(num int) int {
  return int(math.Ceil(math.Sqrt(float64(num))))
}

func is_prime(x int) bool {
  lim := sq(x)
  if x == 2 {return true}
  for i := 2; i <= lim; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  sum := 0
  lim := 2000000
  for i := 2; i < lim; i++ {
    if is_prime(i) {
      sum += i
    }
  }
  fmt.Println(sum)
}
