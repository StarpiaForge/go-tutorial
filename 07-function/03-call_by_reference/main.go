package main

import "fmt"

func callByReference(x *int) {
	// 포인터를 통해 원본 값을 변경할 수 있습니다.
	*x = 10
}

func main() {
	reference := 5
	fmt.Printf("Call By Reference 호출 전: %d\n", reference)
	callByReference(&reference)
	fmt.Printf("Call By Reference 호출 후: %d\n", reference)
}

// 실행 결과
// Call By Reference 호출 전: 5
// Call By Reference 호출 후: 10
