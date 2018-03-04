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
  ctr := 1
  num := 2
  max := 10001
  for ctr < max {
    num++
    if is_prime(num) {
      ctr++
    }
  }
  fmt.Println(num)
}
