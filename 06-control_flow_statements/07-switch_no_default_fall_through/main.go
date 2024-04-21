package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a :", a)

	switch a {
	case 1:
		fmt.Println("a 값은 1 입니다.")
	default:
		fmt.Println("이 코드는 실행되지 않습니다.")
	}

	fmt.Println("switch 문에서 탈출했습니다.")
}

// 실행 결과
// a : 1
// a 값은 1 입니다.
// switch 문에서 탈출했습니다.
