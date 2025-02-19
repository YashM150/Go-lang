package main

import "fmt"

func main() {
	fmt.Println("\t \t \t This is if-else demo:")
	var x int
	var y int

	fmt.Printf("\nEnter value of x: \t")
	fmt.Scan(&x)
	fmt.Printf("\nEnter value of y: \t")
	fmt.Scan(&y)

	if x > 9 {
		fmt.Println("X is greater than 9")
	} else if x > 9 && y > 9 {
		fmt.Println("x & y are greater than 9")
	} else if y > 9 {
		fmt.Println("Y is greater than 9")
	} else {
		fmt.Println("x & y are less than 9")
	}
}
