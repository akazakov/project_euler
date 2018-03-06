package main

/*
PROBLEM

*/

import (
  "fmt"
  //"math"
  //"strconv"
  //"strings"

  //"../myutil"
)

func make_str(num int) []string {

  numerals := []string {
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fivteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen" }

  Tens := []string {
    "twenty",
    "thirty",
    "forty",
    "fifty",
    "sixty",
    "seventy",
    "eighty",
    "ninety" }

  const Hundred = "hundred"
  const And = "and"

  singles := num % 10
  hundreds := num - num % 100
  tens := num - hundreds - singles
  res := make([]string, 0)
  if hundreds != 0 {
    idx := (hundreds / 100) - 1
    res = append(res, numerals[idx], Hundred)
    if tens != 0 || singles != 0 {
      res = append(res, And)
    }
  }
  if tens < 20 && tens >= 10 {
    r := tens + singles - 1
    res = append(res, numerals[r])
  } else {
    if tens > 0 {
      res = append(res, Tens[(tens / 10) - 2])
    }
    if singles > 0 {
      res = append(res, numerals[singles - 1])
    }
  }
  return res
}

func s_len(s []string) int {
  sum := 0
  for _, v := range s {
    sum += len(v)
  }
  return sum
}

func main() {
  lim := 999
  sum := len("onethousand")
  for i := 1; i <= lim; i++ {
    nums_strs := make_str(i)
    fmt.Println(nums_strs)
    for _, v := range nums_strs {
      sum += len(v)
    }
  }
  fmt.Println(sum)
  fmt.Println(s_len(make_str(115)))
}
