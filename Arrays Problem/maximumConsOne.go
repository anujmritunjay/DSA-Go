package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findConsOne(arr []int) {
	count := 0
	maxCount := 0
	for _, val := range arr {
		if val == 1 {
			count++
			maxCount = max(count, maxCount)
		} else {
			count = 0
		}
	}
	fmt.Println(maxCount)
}

func ConsecutiveOne() {
	arr := []int{1, 0, 1, 1, 0, 1}
	fmt.Println("Program to find maximum consecutive one")
	findConsOne(arr)

}
