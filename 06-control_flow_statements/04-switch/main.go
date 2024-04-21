package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("a :", a)

	// switch 문은 기본적으로 이런 형태로 사용합니다.
	switch a {
	// 이런식으로 case 바로 뒤에 a 와 비교할 값을 넣습니다.
	case 1:
		// a == 1 이면 이 코드를 실행
		// 따라서 현재는 이 코드가 실행됩니다.
		fmt.Println("a 값은 1 입니다.")
	// case는 이런식으로 여러개의 값과 비교할 수 도 있습니다.
	case 2, 3, 4:
		// a == 2 혹은 a == 3 혹은 a == 4 이면 이 코드를 실행
		fmt.Println("a 값은 2 혹은 3 혹은 4 입니다.")
	// a 의 값이 case 문에 정의되지 않았다면 이 default 코드 블럭을 실행합니다.
	default:
		fmt.Println("a 값은 1, 2, 3, 4 가 아닙니다.")
	}
}

// 실행 결과
// a : 1
// a 값은 1 입니다.
