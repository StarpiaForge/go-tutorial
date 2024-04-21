package main

import (
	"fmt"
)

func main() {
	// for 문은 기본적으로 이런 형태로 사용합니다.
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d번째 반복입니다.\n", i)
	}

	// while 문의 기능도 수행할 수 있습니다.
	i := 1
	for i <= 5 {
		fmt.Printf("%d번째 반복입니다.\n", i)
		i++
	}

	// 무한 루프도 가능합니다.
	// 다음 코드는 주석처리해두었습니다.
	// 실행시키면 무한 반복입니다.를 멈추지 않고 출력합니다.
	/*
		for {
			fmt.Println("무한 반복입니다.")
		}
	*/
}

// 실행 결과
// 1번째 반복입니다.
// 2번째 반복입니다.
// 3번째 반복입니다.
// 4번째 반복입니다.
// 5번째 반복입니다.
// 1번째 반복입니다.
// 2번째 반복입니다.
// 3번째 반복입니다.
// 4번째 반복입니다.
// 5번째 반복입니다.