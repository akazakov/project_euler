package main

/*
PROBLEM
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

import (
  "fmt"
  //"math"
  //"strconv"
  //"strings"
  //"../myutil"
)

var alldigits uint

func getDigits(num uint) (a , b uint) {
  b = num % 10
  a = (num - b) / 10
  return
}

func divisableBy(num uint, div uint) bool {
  return num % div == 0
}

func cancel(num, denom uint) (uint, uint) {
  num_a, num_b := getDigits(num)
  denom_a, denom_b := getDigits(denom)
  if (num_a == denom_a) {
    return num_b, denom_b
  }

  if (num_a == denom_b) {
    return num_b, denom_a
  }

  if (num_b == denom_a) {
    return num_a, denom_b
  }

  if (num_b == denom_b) {
    return num_a, denom_a
  }
  return 0, 0
}

func normalize(num, denom uint) (uint, uint) {
  var n, d uint
  var c bool
  for n, d, c = normalize_inner(num, denom); c; n, d, c = normalize_inner(n, d) {}
  return n, d
}

func normalize_inner(num, denom uint) (uint, uint, bool) {
  var i uint
  for i = 49; i > 1; i-- {
    if divisableBy(num, i) && divisableBy(denom, i) {
      return num / i, denom / i, true
    }
  }
  return num, denom, false
}

func main() {
  var n_product, d_product , num, denom uint = 1, 1, 0, 0
  for denom = 10; denom < 100; denom++ {
    if denom % 10 == 0 {
      continue
    }
    for num = 10; num < denom; num++ {
      if num % 10 == 0 {
        continue
      }
      n_num, n_denom := normalize(num, denom)
      funny_num, funny_denom := cancel(num, denom)
      if funny_denom == 0 {
        continue
      }
      n_funny_num, n_funny_denom := normalize(funny_num, funny_denom)

      if n_num == n_funny_num && n_denom == n_funny_denom {
        fmt.Printf("%d/%d %d/%d | %d/%d %d/%d\n",
          num, denom, n_num, n_denom,
          funny_num, funny_denom, n_funny_num, n_funny_denom)
        n_product *= n_num
        d_product *= n_denom
      }
    }
  }
  x, y := normalize(n_product, d_product)
  fmt.Printf("%v %v", x, y)
}
