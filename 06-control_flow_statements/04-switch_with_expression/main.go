package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a :", a)

	// Go 에서는 이렇게 switch 뒤에 계산식을 포함시킬 수 있습니다.
	switch x := a * 10; x {
	case 1:
		fmt.Println("x 값은 1 입니다.")
	case 2, 3, 4:
		fmt.Println("x 값은 2 혹은 3 혹은 4 입니다.")
	default:
		// switch 조건식의 결과 값 x가 10이기 때문에 이 코드 블럭이 실행됩니다.
		fmt.Println("x 값은 1, 2, 3, 4 가 아닙니다.")
		fmt.Println("x :", x)
	}
}

// 실행 결과
// a : 1
// x 값은 1, 2, 3, 4 가 아닙니다.
// x : 10
