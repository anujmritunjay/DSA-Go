package main

import "fmt"

func isSorted(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("Array is not sorted")
			return
		}
	}
	fmt.Println("Array is sorted")
}

func IsSorted() {
	arr := []int{6, 1, 2, 3, 4, 5}
	isSorted(arr)

}
