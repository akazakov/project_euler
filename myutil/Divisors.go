package myutil

import (
  "math"
)

func Sq(num int) int {
  return int(math.Floor(math.Sqrt(float64(num))))
}

func Divisors(n int) []int {
  return append(ProperDivisors(n), n)
}

func ProperDivisors(n int) []int {
  factors := make([]int, 1)
  factors[0] = 1
  limt := Sq(n)
  for i := 2; i <= limt; i++ {
    if n % i == 0 {
      factors = append(factors, i)
      if r := n / i; r != i {
        factors = append(factors, r)
      }
    }
  }
  return factors
}

func SumOfProperDivisors(n int) int {
  divisors := ProperDivisors(n)
  sum := 0
  for _, v := range divisors {
    sum += v
  }
  return sum
}

func IsPrime(x int) bool {
  lim := Sq(x)
  if x == 2 { return true }
  if x < 2 { return false }
  for i := 2; i <= lim; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}
