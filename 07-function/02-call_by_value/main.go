package main

import "fmt"

func callByValue(x int) {
	// 매개변수의 값을 변경해도 원본 값에는 영향을 주지 않습니다.
	x = 10
}

func main() {
	value := 5
	fmt.Printf("Call By Value 호출 전: %d\n", value)
	callByValue(value)
	fmt.Printf("Call By Value 호출 후: %d\n", value)
}

// 실행 결과
// Call By Value 호출 전: 5
// Call By Value 호출 후: 5
