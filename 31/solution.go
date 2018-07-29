package main

/*
PROBLEM
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

import (
  //"fmt"
  //"math"
  //"strconv"
  //"strings"

  //"../myutil"
)

type key struct {
  m int
  n int
}

var m map[key]int
var denoms []int
var ctr int

func main() {
  ctr = 0
  m = make(map[key]int)
  denoms = []int{1, 2, 5, 10, 20, 50, 100, 200}
  println(memo(key{8,10000}))
  println(ctr)
  println(len(m))
}

func memo(k key) int {
  saved := m[k]
  if saved != 0 {
    //fmt.Printf("{%d, %d}: %d\n", k.m, k.n, saved)
    return saved
  } else {
    calc := count(k)
    m[k] = calc
    ctr += 1
    //fmt.Printf("new: {%d, %d}: %d\n", k.m, k.n, calc)
    return calc
  }
}

func count(k key) int {
  if k.n == 0 {
    return 1
  }
  if k.n < 0 {
    return 0
  }
  if k.m <= 0 && k.n >= 1 {
    return 0
  }
  return memo(key{k.m - 1, k.n}) + memo(key{k.m, k.n - denoms[k.m - 1]})
  //return memo(key{k.m, k.n - denoms[k.m - 1]}) + memo(key{k.m - 1, k.n})
}
