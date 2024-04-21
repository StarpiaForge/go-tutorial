package main

import (
	"fmt"
)

func main() {
	// range 키워드를 사용하여 슬라이스를 순회할 수 있습니다.
	nums := []int{1, 2, 3, 4, 5}
	for idx, num := range nums {
		fmt.Printf("인덱스: %d, 값: %d\n", idx, num)
	}

	// map 데이터 구조를 생성합니다.
	myMap := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	// range 키워드를 사용하여 map을 순회합니다.
	for key, value := range myMap {
		fmt.Printf("키: %s, 값: %d\n", key, value)
	}
}

// 실행 결과
// 인덱스: 0, 값: 1
// 인덱스: 1, 값: 2
// 인덱스: 2, 값: 3
// 인덱스: 3, 값: 4
// 인덱스: 4, 값: 5
// 키: apple, 값: 1
// 키: banana, 값: 2
// 키: cherry, 값: 3
