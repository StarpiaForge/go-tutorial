package main

import "fmt"

// 가변 인자 함수는 이렇게 선언합니다.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// sum 함수는 가변 인자를 가지는 함수이기 때문에
	// 이처럼 전달하는 매개 변수의 개수를 유동적으로 변경할 수 있습니다.
	fmt.Printf("가변 인자 함수 호출 결과 1: %d\n", sum(1, 2, 3))
	fmt.Printf("가변 인자 함수 호출 결과 2: %d\n", sum(1, 2, 3, 4, 5))
}

// 실행 결과
// 출력: 가변 인자 함수 호출 결과 1: 6
// 출력: 가변 인자 함수 호출 결과 2: 15
