package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	fmt.Println("a :", a)
	fmt.Println("b :", b)

	// Go 에서는 이렇 블록 시작 괄호가 조건식의 다음줄에 들어오는 것을 허용하지 않습니다.
	// 이 경우엔 컴파일 단계에서 에러가 발생합니다.
	if a < b {
		fmt.Println("a 값이 b 값보다 더 작습니다.")
	}
}

// 실행 결과
// unexpected newline, expecting { after if clause
