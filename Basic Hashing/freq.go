package main

import "fmt"

func Freq() {
	var arr = []int{10, 5, 10, 15, 10, 5}
	mp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		ele := arr[i]
		if mp[ele] == 0 {
			mp[ele] = 1
		} else {
			mp[ele] = mp[ele] + 1
		}
	}

	for key, value := range mp {
		fmt.Printf("%d present %d times\n", key, value)
	}
}
