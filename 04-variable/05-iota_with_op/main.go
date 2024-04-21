package main

import "fmt"

const (
	NotStarted float64 = iota * 0.2
	Started
	AlmostHalf
	OverHalf
	AlmostDone
	Done
)

func main() {
	fmt.Printf("NotStarted: %.2f\n", NotStarted)
	fmt.Printf("Started: %.2f\n", Started)
	fmt.Printf("AlmostHalf: %.2f\n", AlmostHalf)
	fmt.Printf("OverHalf: %.2f\n", OverHalf)
	fmt.Printf("AlmostDone: %.2f\n", AlmostDone)
	fmt.Printf("Done: %.2f\n", Done)
}
