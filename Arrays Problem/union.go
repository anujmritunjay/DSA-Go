package main

import "fmt"

func findUnion1(arr1 []int, arr2 []int) {
	mp := make(map[int]int)

	for _, v := range arr1 {
		_, ok := mp[v]
		if !ok {
			mp[v] = 1
		}
	}
	for _, v := range arr2 {
		_, ok := mp[v]
		if !ok {
			mp[v] = 1
		}
	}
	unionArr := make([]int, 0)
	for key, _ := range mp {
		unionArr = append(unionArr, key)
	}
	fmt.Println(unionArr)

}

func FindUnion() {
	fmt.Println("Welcome to union")
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 3, 4, 4, 5}
	findUnion1(arr1, arr2)
}
