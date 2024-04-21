package main

import "fmt"

func main() {
	// 상수는 이렇게 선언합니다.
	// 상수는 중간에 변경할 수 없기 때문에 선언할때 값을 반드시 지정해야 합니다.
	// const <변수 이름> <변수 타입> = <값>
	const number int = 10

	// 이 다음 줄처럼 상수를 변수처럼 값을 변경하려고 하면 참조 오류가 발생합니다.
	// number = 20

	// number에 무슨 값이 들어있는지 출력합니다.
	fmt.Println(number)
}
