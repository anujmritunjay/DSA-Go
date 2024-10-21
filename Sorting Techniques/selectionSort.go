package main

import "fmt"

func sortArr(arr []int) {
	for index, _ := range arr {
		for i, _ := range arr {
			if arr[index] > arr[i] {
				temp := arr[index]
				arr[index] = arr[i]
				arr[i] = temp
			}
		}

	}
	fmt.Println(arr)
}

func SelectionSort() {
	// var arr = []int{13, 46, 24, 52, 20, 9}
	arr := []int{13, 46, 24, 52, 20, 9}
	sortArr(arr)
}
