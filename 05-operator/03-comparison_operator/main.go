package main

import "fmt"

func main() {
	var a, b int
	var c bool
	a = 1
	b = 2

	// a, b 값이 같으면 true 아니면 false
	c = a == b
	fmt.Println(c) // false

	// a 값이 더 크면 true 아니면 false
	c = a > b
	fmt.Println(c) // false

	// a 값이 더 크거나 b 와 같다면 true 아니면 false
	c = a >= b
	fmt.Println(c) // false

	// a 값이 더 작으면 true 아니면 false
	c = a < b
	fmt.Println(c) // true

	// a 값이 더 작거나 b 와 같다면 true 아니면 false
	c = a <= b
	fmt.Println(c) // true

	// a, b 값이 다르면 true 아니면 false
	c = a != b
	fmt.Println(c) // true
}
