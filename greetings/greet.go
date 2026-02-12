package greetings

import "fmt"

func Double(x int) int {
	return x * x
}

func Add(lhs, rhs int) int {
	return lhs + rhs
}

func Greet() {
	fmt.Println("Hello from Go!")
}
