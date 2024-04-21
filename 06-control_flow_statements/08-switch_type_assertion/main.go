package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	switch v := i.(type) {
	case int:
		fmt.Println("i의 두 배는", v*2)
	case string:
		fmt.Println(v + " world!")
	default:
		fmt.Printf("타입 %T 에 대해선 모르겠습니다!\n", v)
	}
}

// 실행 결과
// hello world!
