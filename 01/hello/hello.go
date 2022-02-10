/*
package main
이 패키지의 이름은 main 이다.
main은 프로젝트의 시작점을 뜻한다. 처음 로드 될때 실행된다.
즉, 라이브러리에는 main을 쓰지 않는다.
*/
package main

/*
fmt는 포멧의 약자로 golang을 제공하는 표준패키지이다.
go에서 제공하는 기본패키지, 그 외 의 외부패키지가 있다.
*/
import "fmt"

/*
main 펑션은 별도의 호출 없이도 실행된다.
킹냐! main이니까.
*/
func main() {
	fmt.Println("hello Go")
}
