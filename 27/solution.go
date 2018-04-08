package main

/*
Problem 27
==========


   Euler discovered the remarkable quadratic formula:

                                  n² + n + 41

   It turns out that the formula will produce 40 primes for the consecutive
   values n = 0 to 39. However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41
   is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly
   divisible by 41.

   The incredible formula  n² − 79n + 1601 was discovered, which produces 80
   primes for the consecutive values n = 0 to 79. The product of the
   coefficients, −79 and 1601, is −126479.

   Considering quadratics of the form:

     n² + an + b, where |a| < 1000 and |b| < 1000

     where |n| is the modulus/absolute value of n
     e.g. |11| = 11 and |−4| = 4

   Find the product of the coefficients, a and b, for the quadratic
   expression that produces the maximum number of primes for consecutive
   values of n, starting with n = 0.


   Answer: 69d9e3218fd7abb6ff453ea96505183d
   1 + a + b = p
   a = p - b - 1
   p - b - 1 <= 1000
   p - b - 1 >= -1000
   p <= 1000 + b + 1
   p >= b - 1000 - 1

*/

import (
  "fmt"
  //"math"
  //"strconv"
  //"strings"

  "../myutil"
)

type pair struct {
  a int
  b int
}

func primes_up_to(primes []int, l int) []int {
  bs := make([]int, 0)
  for _, v := range primes {
    if v <= l {
      bs = append(bs, v)
    }
  }
  return bs
}

func print_primes(a, b int) {
  res := make([]int, 0)
  for n := 0; myutil.IsPrime(n*n + n*a + b); n++ {
    res = append(res, n*n + n*a + b)
  }
  fmt.Println(res)
}

func main() {
  primes := make([]int, 0)
  for i := 2; i < 2001; i++ {
    if myutil.IsPrime(i) {
      primes = append(primes, i)
    }
  }
  bs := primes_up_to(primes, 1000)


  res := make(map[pair]int)

  for _, b := range bs {
    ps := primes_up_to(primes, 1001 + b)
    for _, p := range ps {
      a := p - b - 1
      if a > 1000 || a < -1000 {
        continue
      }
      c := 2
      for n := 2; myutil.IsPrime(n*n + n*a + b); n++ { c++ }
      res[pair{a, b}] = c
    }
  }
  fmt.Println(res[pair{1, 41}])

  product := 0
  var max_pair pair
  max_c := 0

  for k, v := range res {
    if v > max_c {
      product = k.a * k.b
      max_pair = k
      max_c = v
    }
  }

  fmt.Println(max_c)
  fmt.Println(product)
  print_primes(max_pair.a, max_pair.b)
}
