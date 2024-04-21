package main

import "fmt"

func main() {
	var a, b, c int
	// ...(상위 비트 표기 생략)
	a = 1 // ... 0001
	b = 2 // ... 0010

	// OR 게이트
	// a, b 각각의 비트를 비교하여 하나라도 1이면 1로 설정
	c = a | b
	fmt.Printf("%04b\n", c) // ... 0011 = 3

	// AND 게이트
	// a, b 각각의 비트를 비교하여 하나라도 0이면 0으로 설정
	c = a & b
	fmt.Printf("%04b\n", c) // ... 0000 = 0

	// Left Shift
	// a 비트 전체를 N번 왼쪽으로 이동
	c = a << 2
	fmt.Printf("%04b\n", c) // ... 0100 = 4

	// Right Shift
	// b 비트 전체를 N번 오른쪽으로 이동
	c = b >> 1
	fmt.Printf("%04b\n", c) // ... 0001 = 0
}
