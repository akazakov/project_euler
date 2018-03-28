package main

/*
PROBLEM
 You are given the following information, but you may prefer to do some
   research for yourself.

     • 1 Jan 1900 was a Monday.
     • Thirty days has September,
       April, June and November.
       All the rest have thirty-one,
       Saving February alone,
       Which has twenty-eight, rain or shine.
       And on leap years, twenty-nine.
     • A leap year occurs on any year evenly divisible by 4, but not on a
       century unless it is divisible by 400.

   How many Sundays fell on the first of the month during the twentieth
   century (1 Jan 1901 to 31 Dec 2000)?


   Answer: a4a042cf4fd6bfb47701cbc8a1653ada
*/

import (
  "fmt"
)

func feb_days(year int) int {
  if year % 4 == 0 {
    if year % 100 == 0 {
      if year % 400 == 0 {
        return 29
      } else {
        return 28
      }
    } else {
      return 29
    }
  }
  return 28
}

func num_days(year, month int) int {
  switch month {
    case 1: return 31
    case 2: return feb_days(year)
    case 3: return 31
    case 4: return 30
    case 5: return 31
    case 6: return 30
    case 7: return 31
    case 8: return 31
    case 9: return 30
    case 10: return 31
    case 11: return 30
    case 12: return 31
  }
  return 31
}

func main() {
  // Init 1901 to Tue
  day := 1
  sundays := 0
  for year := 1901; year <= 2000; year++ {
    for month := 1; month <= 12; month++ {
      day = (day + num_days(year, month)) % 7
      if day == 6 {
        sundays++
      }
    }
  }
  if (day == 6) {
    sundays--
  }
	fmt.Print(sundays)
}
