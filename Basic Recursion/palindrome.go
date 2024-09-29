package main

import "fmt"

func findPalindrome(str string, i int) bool {
	if i >= len(str)/2 {
		return true
	}

	if str[i] != str[len(str)-i-1] {
		return false
	}
	return findPalindrome(str, i+1)
}

func Palindrome() {
	str := "MADAMA"

	fmt.Println(findPalindrome(str, 0))
}
