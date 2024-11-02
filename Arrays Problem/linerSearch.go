package main

import "fmt"

func linearSearch(arr []int, n int) {
	for i, v := range arr {
		if v == n {
			fmt.Println(n, " found at index ", i)
			break
		}
	}

}

func LinearSearch() {
	arr := []int{1, 2, 3, 4, 5}
	n := 2
	linearSearch(arr, n)
}
