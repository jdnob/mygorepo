package main

import (
	"fmt"
)

func fibonacci(k int) int {
	if k < 0 {
		return 0
	} else if k == 0 {
		return 1
	} else {
		return fibonacci(k-1) + fibonacci(k-2)
	}
}

func fact(k int) int {
	if k == 0 {
		return 1
	}
	return k * fact(k-1)
}

/**
 */
func quickmath() (x, y int) {
	x = 12
	y = 13
	return
}

func main() {
	fmt.Println(quickmath())
}
