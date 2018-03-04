package main
/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
  "fmt"
  "container/list"
)

func digits(num int) *list.List {
  l := list.New()
  i := num
  for i > 0 {
    r := i % 10
    l.PushBack(r)
    i = (i - r) / 10
  }
  return l
}

func is_palindrome(l *list.List) bool {
  c := l.Len() / 2
  s := l.Front()
  e := l.Back()
  for i := 0; i < c; i++ {
    if s.Value != e.Value {
      return false
    }
    s = s.Next()
    e = e.Prev()
  }
  return true
}

func largest_palindrome() int {
  max := 1
  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {
      num := i * j
      if num > max && is_palindrome(digits(num)) {
        max = num
      }
    }
  }
  return max
}

func main() {
	fmt.Println(largest_palindrome())
}
