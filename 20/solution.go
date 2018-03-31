package main

/*
PROBLEM
 n! means n × (n − 1) × ... × 3 × 2 × 1

   For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
   and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 =
   27.

   Find the sum of the digits in the number 100!


   Answer: 443cb001c138b2561a0d90720d6ce111
*/

import (
  "fmt"
  //"math"
  //"strconv"
  //"strings"

  //"../myutil"
)


func multiply_by_n(digits []int, m int) []int {
  result := make([]int, len(digits))
  mem := 0
  for i := 0; i < len(digits); i++ {
    r := digits[i] * m + mem
    result[i] = r % 10
    mem = r / 10
  }
  return result
}

func print_digits(digits []int) {
  i := len(digits) - 1
  for ; i > 0 && digits[i] == 0; i-- {}
  for i >= 0 {
    fmt.Printf("%d", digits[i])
    i--
  }
  fmt.Println()
}

func sum(digits []int) int {
  s := 0
  for _, v := range digits {
    s += v
  }
  return s
}

func main() {
  max := 100
  digits := make([]int, 1000)
  digits[0] = 1
  for i := 1; i <= max; i++ {
    digits = multiply_by_n(digits, i)
  }
  print_digits(digits)
  fmt.Print(sum(digits))
}
