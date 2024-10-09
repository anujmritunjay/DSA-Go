package main

import (
	"fmt"
	"math"
)

func HighLowFreq() {
	arr := []int{10, 5, 10, 15, 10, 5}

	mp := make(map[int]int)

	for _, value := range arr {
		if mp[value] == 0 {
			mp[value] = 1
		} else {
			mp[value] = mp[value] + 1
		}
	}

	var maxi, mani int
	maxFreq := math.Inf(-1)
	minFreq := math.Inf(1)

	// Iterate over the map to find the numbers with the highest and lowest frequency
	for key, freq := range mp {
		if float64(freq) > maxFreq {
			maxFreq = float64(freq)
			maxi = key
		}

		if float64(freq) < minFreq {
			minFreq = float64(freq)
			mani = key
		}
	}

	fmt.Println("Number with highest frequency:", maxi)
	fmt.Println("Number with lowest frequency:", mani)
}
