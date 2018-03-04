package main

import "fmt"

func even_fibs_sum() int64 {
  n1 := int64(1)
  n2 := int64(2)
  var sum int64 = 2
  for n2 < 4000000 {
    next := n1 + n2
    n1 = n2
    n2 = next
    if n2 & 1 == 0 {
      sum += n2
    }
  }
  return sum
}

func main() {
	fmt.Println(even_fibs_sum())
}
