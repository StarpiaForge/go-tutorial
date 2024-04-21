package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	fmt.Println("a :", a)
	fmt.Println("b :", b)

	// if 문은 기본적으로 이런 형태로 사용합니다.
	if a > b {
		// if 뒤에 나온 조건이 true 이면 이 코드를 실행
		// a > b == false 이므로 실행되지 않습니다.
		fmt.Println("a 값이 b 값보다 더 큽니다.")
	}

	// if 문은 else-if, else 와 함께 사용할 수 있습니다.
	if a < b {
		// if 뒤에 나온 조건이 true 이면 이 코드를 실행
		// a < b == true 이므로 이 코드가 실행됩니다.
		fmt.Println("a 값이 b 값보다 더 작습니다.")
	} else if a == b {
		// if 조건이 false 이고 else-if 뒤에 나온 조건이 true 이면 이 코드를 실행
		fmt.Println("a, b 값이 서로 같습니다.")
	} else {
		// if, else-if 조건들이 모두 false 이면 이 코드를 실행
		fmt.Println("a 값이 b 값보다 더 작지도, 서로 같지도 않습니다.")
	}
}

// 실행 결과
// a : 1
// b : 2
// a 값이 b 값보다 더 작습니다.
