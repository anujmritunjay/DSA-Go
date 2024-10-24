package main

import "fmt"

func bruteForceMissing(arr []int) {

	for i := range arr {
		flag := false
		for j := range arr {
			if arr[j] == i {
				flag = false
				break
			}
		}
		if flag == false {
			fmt.Println(i)
			return
		}

	}
	return
}

func MissingNum() {
	arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	bruteForceMissing(arr)
}
