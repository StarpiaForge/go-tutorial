package main

import "fmt"

func main() {
	var a int
	a = 1

	// 포인터는 변수의 데이터 타입 앞에 * 을 붙혀서 선언합니다.
	var b *int
	// & 연산자를 이용하여 포인터 변수에 a 변수의 주소값을 할당합니다.
	b = &a

	// * 연산자를 이용하여 포인터 변수가 가르키는 변수 공간을 참조할 수 있습니다.
	*b = *b + 1
	fmt.Printf("pointing variable b : %d\n", *b)
	fmt.Printf("orginal variable a : %d\n", a)

	// Go 에서 포인터 산술 기능은 허용되지 않습니다.
	// 따라서 아래의 코드는 컴파일 에러를 발생시킵니다.
	b = b + 1
}
