package main

/*
PROBLEM

*/

import (
  "fmt"
  "math/big"
  "strconv"
  "strings"

  //"../myutil"
)

func main() {
  r := new(big.Int).SetInt64(1)
  two := new(big.Int).SetInt64(2)
  for i := 0; i < 1000; i++{
    r.Mul(r, two)
  }
	fmt.Println(r)
  digs := strings.Split(r.String(), "")
  sum := 0
  for _, d := range digs {
    num, _ := strconv.Atoi(d)
    sum += num
  }
	fmt.Println(sum)
}
