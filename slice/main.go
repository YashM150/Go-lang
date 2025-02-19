package main

import "fmt"

func main() {

	fmt.Println("This is demo of slice")
	numbers := make([]int, 3, 5)

	fmt.Println("The initial numbers list is:", numbers)
	fmt.Println("Appending 1,2,3,4,5,6")
	numbers = append(numbers, 1, 2, 3, 4, 5, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	fmt.Println("Appending 1,2,3,4,5,6")
	numbers = append(numbers, 1, 2, 3, 4, 5, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

}
