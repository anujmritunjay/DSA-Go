package main

import (
	"fmt"
	"math"
)

func findGcd(num1 int, num2 int) int {
	gcd := 1

	// for i := 1; i <= int(math.Min(float64(num1), float64(num2))); i++ {
	// 	if num1%i == 0 && num2%i == 0 {
	// 		gcd = i
	// 	}
	// }

	for i := int(math.Min(float64(num1), float64(num2))); i > 1; i-- {
		if num1%i == 0 && num2%i == 0 {
			fmt.Println((i))
			gcd = i
		}
	}

	return gcd

}

func LCMGCD() {
	fmt.Println("Program to find LCM and GCD")

	num1 := 5
	num2 := 10
	fmt.Println(findGcd(num1, num2))
}
