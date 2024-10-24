package main

import (
	"fmt"
	"math"
	"sort"
)

func bruteForce(arr []int) {
	sort.Ints(arr)
	fmt.Printf("Second Lowest: %d\n", arr[1])
	fmt.Printf("Second Highest: %d\n", arr[len(arr)-2])
}

func better(arr []int) {
	smallest := math.MaxInt
	secondSmallest := math.MaxInt
	highest := math.MinInt
	secondHighest := math.MinInt

	for _, val := range arr {
		smallest = int(math.Min(float64(val), float64(smallest)))
		highest = int(math.Max(float64(val), float64(highest)))
	}

	for _, val := range arr {
		if val < secondSmallest && val != smallest {
			secondSmallest = val
		}
		if val < secondSmallest && val != smallest {
			secondSmallest = val
		}
	}

	fmt.Println(secondSmallest, secondHighest)
}

func sHighest(arr []int) int {
	highest := math.MinInt
	secondHighest := math.MinInt

	for _, val := range arr {
		if val > highest {
			secondHighest = highest
			highest = val
		} else if val > secondHighest && val != highest {
			secondHighest = val
		}
	}
	return secondHighest

}

func sLowest(arr []int) int {
	low := math.MaxInt
	secondLowest := math.MaxInt

	for _, val := range arr {
		if val < low {
			secondLowest = low
			low = val
		} else if val < secondLowest && val != low {
			secondLowest = val
		}
	}
	return secondLowest
}

func optimal(arr []int) {
	secondHighest := sHighest(arr)
	secondLowest := sLowest(arr)
	fmt.Println(secondHighest, secondLowest)
}

func SecondLargest() {
	arr := []int{1, 2, 4, 6, 7, 5}
	bruteForce(arr)
	better(arr)
	optimal(arr)

}
