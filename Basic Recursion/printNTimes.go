package main

import "fmt"

func print(n int) {
	if n < 1 {
		return
	}
	fmt.Println("Hello Recursion")
	print(n - 1)
}

func PrintNTimes() {
	n := 5

	print(n)
}
