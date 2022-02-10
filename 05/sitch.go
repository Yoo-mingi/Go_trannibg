package main

import (
	"fmt"
)

func main() {
	a := 30
	switch a {
	case 31:
		fmt.Printf("zzz")
	case 30:
		fmt.Print("zzzzzzz")
	}
}

//a 생략시 true로 간주
