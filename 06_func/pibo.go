package main

import "fmt"

func f(x int) int {
	if x == 0 {
		return 1
	} else if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}

func main() {
	rst := f(10)
	fmt.Print(rst)
}
