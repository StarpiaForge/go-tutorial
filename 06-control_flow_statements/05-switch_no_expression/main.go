package main

import "fmt"

func main() {
	a := 1
	fmt.Println("a :", a)

	// switch 문 뒤에 변수나 계산식이 없습니다.
	switch {
	case a < 2:
		fmt.Println("a 값은 2 보다 작습니다.")
	case a == 2:
		fmt.Println("a 값은 2 입니다.")
	default:
		fmt.Println("a 값은 2 보다 큽니다.")
	}
}

// 실행 결과
// a : 1
// a 값은 2 보다 작습니다.
