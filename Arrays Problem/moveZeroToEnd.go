package main

import "fmt"

func moveZeroBrute(arr []int) {
	nz := make([]int, 0)
	for _, val := range arr {
		if val != 0 {
			nz = append(nz, val)
		}
	}
	copy(arr, nz)
	for i := len(nz); i < len(arr); i++ {
		arr[i] = 0
	}
	fmt.Println(arr)
}

func moveZeroOptimal(arr []int) {
	j := -1

	for i, val := range arr {
		if val == 0 {
			j = i
			break
		}
	}

	for i := j + 1; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	fmt.Println(arr)
}

func MoveZeroToEnd() {
	fmt.Println("Move Zero to end")
	arr := []int{1, 0, 2, 3, 2, 0, 0, 4, 5, 1}
	moveZeroOptimal(arr)
}
