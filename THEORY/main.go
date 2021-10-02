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

func multiply(a int, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func nakedReturn(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") //defer를 사용하면 함수가 끝나고 실행시킴
	length = len(name)            //length와 uppercase는 위에서 이미 만듦
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func makeArray() {
	names := [5]string{"hi"}
	names[1] = "good"
	names[2] = "bye"
	names[3] = "hello"
	names[4] = "workd"
	fmt.Println(names)

	namesNoLength := []string{"hi"}
	namesNoLength = append(namesNoLength, "new")
	fmt.Println(namesNoLength)
}

func makeMaps() {
	zaehuun := map[string]string{}
	zaehuun["name"] = "hoon"
	zaehuun["age"] = "26"
	for key, value := range zaehuun {
		fmt.Println(key, value)
	}
}

func main() {
	makeMaps()
}
