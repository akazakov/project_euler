package main

/*
PROBLEM
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is thought
that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/

import (
  "fmt"
  "../myutil"
)

func collatz_seq_len(n int) int {
  i := n
  l := 1
  for i > 1 {
    if i % 2 == 0 {
      i = i / 2
    } else  {
      i = 3*i + 1
    }
    l++
  }
  return l
}

func main() {
  m := 0
  n := 0
  for i := 0; i < 1000000; i++ {
    s := collatz_seq_len(i)
    m = myutil.Max(m, s)
    if m == s {
      n = i
    }
  }
  fmt.Println(m)
  fmt.Println(n)
}
