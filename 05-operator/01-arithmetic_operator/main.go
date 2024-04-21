package main

import "fmt"

func main() {
	var a int
	// 더하기 연
	a = 1 + 1
	fmt.Println(a) // 2

	// 빼기 연산
	a = 2 - 1
	fmt.Println(a) // 1

	// 곱하기 연산자
	a = 3 * 1
	fmt.Println(a) // 3

	// 나누기 연산자
	a = 4 / 2
	fmt.Println(a) // 2

	// 나머지 연산자
	a = 1 % 4
	fmt.Println(a) // 1

	// 증감 연산자
	// 1 증가
	b := 5
	b++
	fmt.Println(b) // 6

	// 1 감소
	c := 5
	c--
	fmt.Println(c) // 4
}
