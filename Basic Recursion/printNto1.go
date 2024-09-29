package main

import "fmt"

func prints(n int) {
	if n < 1 {
		return
	}
	fmt.Println(n)
	prints(n - 1)
}

func Nto1() {
	n := 10
	prints(n)

}
