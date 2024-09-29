package main

import (
	"fmt"
	"math"
)

func findDivisor(num int) []int {
	divisor := make([]int, 0)

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisor = append(divisor, i)

			if i != num/i {
				divisor = append(divisor, num/i)
			}
		}
	}

	return divisor
}

func Divisor() {
	num := 36
	fmt.Println(findDivisor(num))
}
