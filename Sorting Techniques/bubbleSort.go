package main

import "fmt"

func sortArrBubble(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}

func BubbleSort() {
	arr := []int{13, 46, 24, 52, 20, 9}
	sortArrBubble(arr)

}
