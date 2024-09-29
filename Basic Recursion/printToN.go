package main

import "fmt"

func printFunc(i int, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	printFunc(i+1, n)
}

func PrintTillN() {
	n := 10
	printFunc(1, n)
}
