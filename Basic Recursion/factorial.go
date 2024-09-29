package main

import "fmt"

func findFactorial(num int, fact int) int {
	if num == 0 {
		return fact
	}
	fact = fact * num
	return findFactorial(num-1, fact)
}

func Factorial() {
	num := 5

	fmt.Println(findFactorial(num, 1))
}
