package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is loop demo:")
	var counter int
	fmt.Println("Enter a value for counter:\t")
	fmt.Scan(&counter)

	fmt.Println("\n\n1.Normal for:")
	for i := 0; i < counter; i++ {
		fmt.Printf("This is %d iteration.\n", i)
	}

	fmt.Println("\n\n2.Infinite loop for:")
	count := 0
	for {
		fmt.Println("This is Infinite loop")
		count++
		if count == counter {
			break
		}
	}

	fmt.Println("\n\n3.Range keyword int for:")
	var length int
	fmt.Println("Enter length of array:\t")
	fmt.Scan(&length)
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		var temp int
		fmt.Printf("\nEnter the value of %d element of array:", (i + 1))
		fmt.Scan(&temp)
		numbers = append(numbers, temp)
	}
	fmt.Println("Printing the array")
	for index, value := range numbers {
		fmt.Printf("\nThe value at %d index is %d", index, value)
	}

	fmt.Println("\n\n4.Range keyword char for:")
	String := "Hello World"
	fmt.Println("The string is: ", String)
	for index, char := range String {
		fmt.Printf("\nThe value at %d index is %c", index, char)
	}
}
