package main

/*
Problem 22
==========


   Using [1]names.txt, a 46K text file containing over five-thousand first
   names, begin by sorting it into alphabetical order. Then working out the
   alphabetical value for each name, multiply this value by its alphabetical
   position in the list to obtain a name score.

   For example, when the list is sorted into alphabetical order, COLIN, which
   is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So,
   COLIN would obtain a score of 938 × 53 = 49714.

   What is the total of all the name scores in the file?


   Visible links
   1. names.txt
   Answer: f2c9c91cb025746f781fa4db8be3983f

*/

import (
  "fmt"
  "strings"
  "sort"
  "../myutil"
)

func main() {
  qNames, _ := myutil.ReadCSVFile("names.txt")
  names := make([]string, len(qNames))
  for i, name := range qNames {
    names[i] = strings.Trim(name, "\"")
  }
  sort.Strings(names)
  sum := 0
  for i, name := range names {
    nameValue := 0
    for _, c := range name {
      nameValue += int(c - 'A') + 1
    }
    nameValue *= i + 1
    sum += nameValue
  }
  fmt.Print(sum)
}
