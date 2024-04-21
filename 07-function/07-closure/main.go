package main

import "fmt"

// 클로저 함수
func outer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	counter := outer()
	fmt.Println("클로저 카운트:", counter())
	fmt.Println("클로저 카운트:", counter())
}

// 실행 결과
// 클로저 카운트: 1
// 클로저 카운트: 2
