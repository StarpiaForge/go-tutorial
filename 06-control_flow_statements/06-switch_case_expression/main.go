package main

import "fmt"

func main() {
	a := 1
	fmt.Println("a :", a)

	// case 문 뒤에 계산식을 사용합니다.
	switch {
	case a < 2:
		fmt.Println("a 값은 2 보다 작습니다.")
	case 2*a == 2:
		fmt.Println("a 값은 1 입니다.")
	default:
		fmt.Println("a 값은 1이 아닙니다.")
	}
}

// 실행 결과
// a : 1
// a 값은 2 보다 작습니다.
// a 값은 1 입니다.
