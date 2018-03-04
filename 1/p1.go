package main

import "fmt"

func sumz(step, max int) int {
	sum := 0
	for i := 0; i < max; i += step {
		sum += i
	}
	return sum
}

func main() {
	threes := sumz(3, 1000)
	fives := sumz(5, 1000)
	fivteens := sumz(15, 1000)
	fmt.Println(threes + fives - fivteens)
}
