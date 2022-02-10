package main

import "fmt"

func main() {
	//for 1번째. 변수와 변화를 빼고 조건만
	x := 0
	for x = 0; x < 10; x++ {
		fmt.Println(x)
	}
	fmt.Print(x)
	fmt.Print(x)
	//야발년아 헷갈리게 하지마 잘 되잖아ㅡㅡ
	/*
		i := 0
		for i < 10 {
			fmt.Println(i)
			i += 1
		}
		//for 2번째. 정---석
		for a := 0; a < 10; a++ {

		}
		//주의해야할 것은, go의 변수의 생명은 중괄호 까지다.
	*/
}
