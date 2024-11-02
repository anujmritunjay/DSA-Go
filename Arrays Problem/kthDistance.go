package main

import "fmt"

func findKthDistance(arr []int, k int) bool {
	mp := make(map[int]int)

	for i, v := range arr {
		_, ok := mp[v]

		if ok && i-mp[v] <= k {
			return true
		}
		mp[v] = i
	}
	return false
}

func KthDistance() {
	arr := []int{1, 2, 3, 1, 4, 5}
	k := 3

	fmt.Println(findKthDistance(arr, k))
}
