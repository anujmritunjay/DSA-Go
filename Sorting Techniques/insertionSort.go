package main

import "fmt"

func sortArrInsertion(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		fmt.Println("Not enough element to sort")
	}

	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[i]
			arr[i] = arr[j-1]
			arr[j-i] = temp
			j--
		}
	}
	fmt.Println(arr)

}

func InsertionSort() {
	arr := []int{13, 46, 24, 52, 20, 9}
	sortArrInsertion(arr)

}
