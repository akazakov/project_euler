package myutil

import (
  "math"
)

func Sq(num int) int {
  return int(math.Ceil(math.Sqrt(float64(num))))
}

func Divisors(n int) []int {
  factors := make([]int, 0)
  limt := Sq(n)
  for i := 1; i < limt; i++ {
    if n % i == 0 {
      factors = append(factors, i, n / i)
    }
  }
  return factors
}
