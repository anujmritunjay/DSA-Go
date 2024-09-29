package main

import "fmt"

func printSum(n int, sum int) int {
	if n == 0 {
		return sum
	}
	sum = sum + n
	return printSum(n-1, sum)
}

func SumOfFirstN() {
	n := 10

	fmt.Println(printSum(n, 0))

}
