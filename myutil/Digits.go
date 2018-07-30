package myutil
import (
  "container/list"
)

func GetDigits(num int) *list.List {
  l := list.New()
  i := num
  for i > 0 {
    r := i % 10
    l.PushBack(r)
    i = (i - r) / 10
  }
  return l
}

func GetDigitsMask(num uint) (mask uint, isunique bool) {
  isunique = true
  mask = 0
  i := num
  for i > 0 {
    r := i % 10
    var d uint = 1 << r
    isunique = ((mask & d) == 0)
    if !isunique {
      return
    }
    mask |= d
    i = (i - r) / 10
  }
  return mask, true
}
