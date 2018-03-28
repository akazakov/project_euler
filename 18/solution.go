package main

/*
PROBLEM

   Find the maximum total from top to bottom of the triangle below:

                                       75
                                     95 64
                                    17 47 82
                                  18 35 87 10
                                 20 04 82 47 65
                               19 01 23 75 03 34
                              88 02 77 73 07 63 67
                            99 65 04 28 06 16 70 92
                           41 41 26 56 83 40 80 70 33
                         41 48 72 33 47 32 37 16 94 29
                        53 71 44 65 25 43 91 52 97 51 14
                      70 11 33 28 77 73 17 78 39 68 17 57
                     91 71 52 38 17 14 91 43 58 50 27 29 48
                   63 66 04 68 89 53 67 30 73 16 69 87 40 31
                  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
*/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	//"math"
	//"strings"
	"../myutil"
)


func string_to_nums(s string) ([]int, error) {
	num_strs := strings.Split(s, " ")
	nums := make([]int, len(num_strs))
	for i, s := range num_strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}

func strings_to_nums(ss []string) ([][]int, error) {
	rows := make([][]int, len(ss))
	for i, s := range ss {
		row, err := string_to_nums(s)
		if err != nil {
			return nil, err
		}
		rows[i] = row
	}
	return rows, nil
}

func main() {
	ss, err := myutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, err2 := strings_to_nums(ss)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

  for i := len(rows) - 2; i >= 0; i-- {
    for j, val := range rows[i] {
      rows[i][j] = val + myutil.Max(rows[i+1][j], rows[i+1][j+1])
    }
  }

	fmt.Print(rows[0][0])
}
