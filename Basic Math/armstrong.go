package main

import (
	"fmt"
	"math"
	"strconv"
)

func isArmstrong(num int) bool {
	k := len(strconv.Itoa(num))
	temp := num

	sums := 0

	for temp > 0 {
		last := temp % 10
		sums = sums + int(math.Pow(float64(last), float64(k)))
		temp = temp / 10
	}
	return sums == num
}

func ArmstrongNumber() {
	num := 153

	fmt.Println(isArmstrong(num))
}
