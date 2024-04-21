package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2

	// b = b + a
	b += a
	fmt.Println(b) // 3

	// b = b - a
	b -= a
	fmt.Println(b) // 2

	// b = b * a
	b *= a
	fmt.Println(b) // 2

	// b = b >> a
	b >>= a
	fmt.Println(b) // 1

	// b = b | a
	b |= a
	fmt.Println(b) // 1

	// b = b & a
	b &= a
	fmt.Println(b) // 1
}
