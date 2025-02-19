package main

import "fmt"

func simplefunc() {
	fmt.Println("Simple print function")
}

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("This is demo of functions in go:")
	simplefunc()
	fmt.Println("Addition of two number", add(2, 3))
}
