package main

import "fmt"

func removeDuplicate(arr []int) {
	seen := make(map[int]bool)
	uniqueArr := []int{}
	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			uniqueArr = append(uniqueArr, val)
		}

	}
	fmt.Println("The unique array is: ", uniqueArr)
}

func RemoveDuplicate() {
	arr := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}
	removeDuplicate(arr)
}
