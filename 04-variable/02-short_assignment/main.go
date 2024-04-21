package main

import "fmt"

func main() {
	// #2 단축 선언 연산자 (Short Assignment Statement)
	//
	// 변수는 이렇게도 선언이 가능합니다.
	// <변수 이름> := <값>
	number := 10

	// number 변수에 10이라는 값을 할당합니다.
	number = 20

	// number에 무슨 값이 들어있는지 출력합니다.
	fmt.Println(number)
}
