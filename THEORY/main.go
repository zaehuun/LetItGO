package main

/*
1. main이 진입점
2. 컴파일러는 패키지의 이름이 main인 것부터 찾아냄
3. 라이브러리라면 굳이 컴파일 필요 없기에 main이 아니어도 됨
*/

import "fmt"

/*
1. 곱하기 함수를 생성한다.
*/
func multiply(a int, b int) int {
	return a * b
}

/*
1. Go 프로그램의 시작점
*/
func main() {
	var text string = "Hello World!"
	text2 := "Hello GO" //타입을 알아서 지정, 함수 내에서만 가능
	fmt.Println(text)
	fmt.Println(text2)
	fmt.Println(multiply(2, 2))
}
