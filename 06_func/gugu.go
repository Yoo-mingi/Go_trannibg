package main

import "fmt"

func main() {
	first_num := 1
	last_num := 1

	for first_num = 1; first_num <= 9; first_num++ {

		for last_num = 1; last_num <= 9; last_num++ {
			if first_num == 3 && last_num == 5 {
				continue
			}
			fmt.Printf("%d x %d = %d \n", first_num, last_num, first_num*last_num)
		}
	}
}
