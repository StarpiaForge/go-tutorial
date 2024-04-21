package main

import "fmt"

// divide 함수는 int, int, error 3개의 값을 동시에 반환하는 함수입니다.
func divide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("0으로 나눌 수 없습니다")
	}
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder, nil
}

func main() {
	// Go 에서는 이처럼 여러개의 반환값을 가지는 함수를 만들고 사용할 수 있습니다.
	q, r, err := divide(10, 3)
	if err != nil {
		fmt.Printf("에러 발생: %v\n", err)
	} else {
		fmt.Printf("나눗셈 결과: 몫(%d)/나머지(%d)\n", q, r)
	}
}

// 실행 결과
// 나눗셈 결과: 몫(3)/나머지(1)
