package main

import (
	"fmt"
	"math"
)

func sq(num int64) int64 {
  return int64(math.Ceil(math.Sqrt(float64(num))))
}

func is_prime(x int64) bool {
  lim := sq(x)
  for i := int64(2); i <= lim ; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func max_prime_factor(num int64) int64 {
  max := int64(1);
  lim := sq(num);
  for i := int64(2); i < lim; i++ {
    c2 := num / i
    if (c2 * i == num) {
      if is_prime(c2) {
        max = c2
        break;
      }
      if is_prime(i) {
        max = i;
      }
    }
  }
  return max
}


func main() {
  num :=  int64(600851475143)
	fmt.Println(max_prime_factor(num))
}
