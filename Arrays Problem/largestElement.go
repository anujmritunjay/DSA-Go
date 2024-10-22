package main

import (
	"fmt"
	"math"
	"sort"
)

func largestBrute(arr []int) {
	sort.Ints(arr)
	fmt.Printf("Highest element in the array is: %d", arr[len(arr)-1])
}

func largestNoSort(arr []int) {
	maxi := int(math.Inf(-1))
	for _, value := range arr {
		if value > maxi {
			maxi = value
		}
	}
	fmt.Printf("Highest element in the array is: %d", maxi)
}

func LargestElement() {
	arr := []int{2, 5, 1, 3, 0}
	largestBrute(arr)
	largestNoSort(arr)

}
