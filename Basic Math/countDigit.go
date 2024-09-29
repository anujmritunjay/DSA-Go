package main

import "fmt"

func CountDigit() {
	n := 2446

	temp := n
	count := 0

	for temp > 0 {
		last := temp % 10
		if last != 0 && n%last == 0 {
			count++
		}

		temp = temp / 10
	}
	fmt.Println(count)
}
