package main

/*
Problem 30
==========


   Surprisingly there are only three numbers that can be written as the sum
   of fourth powers of their digits:

     1634 = 1^4 + 6^4 + 3^4 + 4^4
     8208 = 8^4 + 2^4 + 0^4 + 8^4
     9474 = 9^4 + 4^4 + 7^4 + 4^4

   As 1 = 1^4 is not a sum it is not included.

   The sum of these numbers is 1634 + 8208 + 9474 = 19316.

   Find the sum of all the numbers that can be written as the sum of fifth
   powers of their digits.


   Answer: 27a1779a8a8c323a307ac8a70bc4489d
*/

import (
  "fmt"
  //"math"
  "strconv"
  //"strings"

  //"../myutil"
)

func main() {
  // there can only be max 6 digits
  digits := make([]int, 10)
  for i := 0 ; i < 10; i++ {
    digits[i] = i*i*i*i*i
  }
  lim := 6*9*9*9*9*9
  res := make([]int, 0)
  ans := 0
  for i := 2; i < lim; i++ {
    str := strconv.Itoa(i)
    sum := 0
    for _, c := range str {
      d := int(c - '0')
      sum += digits[d]
    }
    if sum == i {
      res = append(res, sum)
      ans += sum
    }
  }
  fmt.Print(ans)
}
