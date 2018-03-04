package main

/*
PROBLEM
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a**2 + b**2 = c**2
For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/

import "fmt"

func main() {
  for i := 1; i < 998; i++ {
    for j := 1; j < 1000 - i - 1; j++ {
      k := 1000 - i - j
      if i*i + j*j == k*k {
        fmt.Println(i, j, k)
        fmt.Println(i*j*k)
      }
    }
  }
}
