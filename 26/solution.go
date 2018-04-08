package main

/*
PROBLEM 26
   A unit fraction contains 1 in the numerator. The decimal representation of
   the unit fractions with denominators 2 to 10 are given:

     1/2  =  0.5
     1/3  =  0.(3)
     1/4  =  0.25
     1/5  =  0.2
     1/6  =  0.1(6)
     1/7  =  0.(142857)
     1/8  =  0.125
     1/9  =  0.(1)
     1/10 =  0.1

   Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can
   be seen that 1/7 has a 6-digit recurring cycle.

   Find the value of d < 1000 for which ^1/[d] contains the longest recurring
   cycle in its decimal fraction part.


   Answer: 6aab1270668d8cac7cef2566a1c5f569
*/

import (
  "fmt"
)

func find_repeat(i int) int {
  seen := make([]bool, i)
  r := 10 % i
  for !seen[r] {
    seen[r] = true
    r = r*10 % i
  }
  return r
}

func seq_len(i int) int {
  r_r := find_repeat(i)
  if r_r > 0 {
    r := 1
    // find the start of the cycle
    for r != r_r {
      r = r*10 % i
    }
    count := 1
    r = r*10 % i
    // cycle once
    for r != r_r {
      r = r*10 % i
      count += 1
    }
    return count
  } else {
    return 0
  }
}

func main() {
  d := 0
  max := 0
  for i := 2; i < 1000; i ++ {
    repeat_len := seq_len(i)
    fmt.Printf("%d %d\n", i, repeat_len)
    if repeat_len > max {
      d = i
      max = repeat_len
    }
  }
	fmt.Println(d)
  find_repeat(7)
}
