package main

import "fmt"

func main() {
	var a, b, c bool
	a = true
	b = false

	// OR 연산자
	// a, b 둘 중 하나라도 true 값이면 true
	// 모두 false 값이면 false
	c = a || b
	fmt.Println(c) // true

	// AND 연산자
	// a, b 둘 중 하나라도 false 값이면 false
	// 모두 true 값이면 true
	c = a && b
	fmt.Println(c) // false

	// NOT 연산자
	// a 가 true 값이면 false
	// a 가 false 값이면 true
	c = !a
	fmt.Println(c) // false
}
