package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func fun1(x, y int) (int, int) {
	return y, x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d+%d=%d \n", i, i+2, i+i+2)
	}
}
