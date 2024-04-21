package main

import "fmt"

const (
	Monday    = iota // Monday = 0
	Tuesday          // Tuesday = 1
	Wednesday        // Wednesday = 2
	Thursday         // Thursday = 3
	Friday           // Friday = 4
	Saturday         // Saturday = 5
	Sunday           // Sunday = 6
	// 무언가 추가하고 싶다면 더 할수도 있습니다.
)

func main() {
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("Sunday:", Sunday)
}
