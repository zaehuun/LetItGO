package main

/*
1. main이 진입점
2. 컴파일러는 패키지의 이름이 main인 것부터 찾아냄
3. 라이브러리라면 굳이 컴파일 필요 없기에 main이 아니어도 됨
*/

import (
	"fmt"
	"strings"
)

/*
21/10/01 함수
1. 곱하기 함수를 생성한다.
*/
func multiply(a int, b int) int {
	return a * b
}

/*
21/10/01 함수
2. 인자의 타입이 같으면 아래도 가능
*/
func divide(a, b int) int {
	return a / b
}

/*
21/10/01 함수
3. 여러 개 리턴 가능
*/
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func nakedReturn(name string) (length int, uppercase string) {
	length = len(name) //length와 uppercase는 위에서 이미 만듦
	uppercase = strings.ToUpper(name)
	return
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
	fmt.Println(divide(2, 2))
	totalLength, upperName := lenAndUpper("zaehuun")
	tmpValue, _ := lenAndUpper("zaehuun") // _ 를 사용하면 무시 가능
	fmt.Println(totalLength, upperName, tmpValue)
	repeatMe("A", "B", "C", "D")
}
