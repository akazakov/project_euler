package main

/*
PROBLEM
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

*/

import (
  "fmt"
  //"math"
  //"strconv"
  //"strings"

  "../myutil"
)

var alldigits uint

func checkMask(maskA, maskB, maskC uint) bool {
  if (maskA & maskB == 0) && (maskA & maskC == 0) && (maskB & maskC == 0) {
    return (maskA | maskB | maskC) == alldigits
  }
  return false
}


func main() {
  alldigits = 0
  m := make(map[uint]bool)
  var sum uint = 0
  var i,j,jlim,product uint
  for i = 1; i <= 9; i++ {
    alldigits |= 1 << i
  }
  for i = 1; i < 99; i ++ {
    maskA, isUnique := myutil.GetDigitsMask(i)
    if !isUnique {
      continue
    }
    jlim = 10000 / i
    for j = 1; j < jlim; j++ {
      maskB, isUnique := myutil.GetDigitsMask(j)
      if !isUnique {
        continue
      }
      product = i * j
      maskC, isUnique := myutil.GetDigitsMask(product)
      if !isUnique {
        continue
      }
      if checkMask(maskA, maskB, maskC) {
        _, exist := m[product]
        if !exist {
          sum += product
          m[product] = true
        }
      }
    }
  }
  println(sum)
	fmt.Printf("%t\n", checkPandigital(123, 456, 789))
}
