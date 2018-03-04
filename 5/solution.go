package main

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?*/

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
  num := 1
  max := 20
  primes := make([]int, 0)
  for i := 2; i <= max; i++ {
    if is_prime(i) {
      fmt.Println(i)
      num *= i
      primes = append(primes, i)
    }
  }
  fmt.Println("Start num ", num)

  for i := 2; i <= max; i++ {
    if num % i != 0 {
      fmt.Println("Non div: ", i)
      remainder := i
      for _, value := range primes {
        if remainder % value == 0 {
          remainder = remainder / value
        }
      }
      num *= remainder
      primes = append(primes, remainder)
    }
  }
  fmt.Println(num)
}
