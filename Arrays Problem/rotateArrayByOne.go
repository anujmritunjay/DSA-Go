package main

import "fmt"

func rotateArrayByOneBrute(arr []int) {
	temp := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		fmt.Println(i)
		temp[i-1] = arr[i]
	}
	temp[len(temp)-1] = arr[0]
	fmt.Println(temp)

}

func RotateArrayByOne() {
	arr := []int{1, 2, 3, 4, 5}
	rotateArrayByOneBrute((arr))

}
