package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a :", a)

	// Go 에서는 이렇게 boolean 이 아닌 값이 조건식으로 올 수 없습니다.
	// 이 경우엔 컴파일 단계에서 에러가 발생합니다.
	if a {
		fmt.Println("a 값이 b 값보다 더 작습니다.")
	}
}

// 실행 결과
// non-boolean condition in if statement
