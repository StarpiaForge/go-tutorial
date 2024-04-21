package main

import "fmt"

func main() {
	// 익명 함수는 이렇게 선언 하고 즉시 호출할 수 있습니다.
	func() {
		fmt.Println("이것은 익명 함수입니다.")
	}()

	// 익명 함수는 변수에도 할당하여 사용할 수 있습니다.
	anonymousFunc := func() {
		fmt.Println("익명 함수는 이렇게 변수에 할당할 수 있습니다.")
	}
	anonymousFunc()
}

// 실행 결과
// 이것은 익명 함수입니다.
// 익명 함수는 이렇게 변수에 할당할 수 있습니다.
