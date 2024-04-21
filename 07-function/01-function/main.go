package main

import "fmt"

// add 함수
func add(x int, y int) int {
	return x + y
}

func main() {
	a := 1
	b := 2
	fmt.Printf("두 수의 합: %d\n", add(a, b))
}

// 실행 결과
// 두 수의 합: 3
