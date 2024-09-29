package main

import (
	"fmt"
	"slices"
)

func reversSliceASpace(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func reveredSliceBrute(nums []int) []int {
	ans := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}
	return ans
}

func reverseRecursive(nums []int, start int, end int) []int {
	if start < end {
		nums[start], nums[end] = nums[end], nums[start]
		return reverseRecursive(nums, start+1, end-1)
	}
	return nums
}

func reverseInbuilt(nums []int) []int {
	slices.Reverse(nums)
	return nums
}

func Reverse() {
	var nums = []int{5, 4, 3, 2, 1}
	// reveredSlice := reversSliceASpace(nums)
	// reveredSlice := reveredSliceBrute(nums)
	// reveredSlice := reverseRecursive(nums, 0, len(nums)-1)
	reveredSlice := reverseInbuilt(nums)
	fmt.Println(reveredSlice)
}
