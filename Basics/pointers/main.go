package main

import "fmt"

func swapping(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("This is pointer demo:")
	fmt.Printf("\nEnter the value:")
	var num int
	fmt.Scan(&num)

	ptr := &num
	fmt.Printf("\n Value of num: %d", num)
	fmt.Printf("\n Num is stored in: %p", ptr)
	fmt.Printf("\n So Value stored in %p is %d", ptr, *ptr)

	fmt.Println("\nLet us implement the simple swapping algorithm using the pointers!!")
	fmt.Printf("\n Enter another number:\t")
	var num1 int
	fmt.Scan(&num1)

	fmt.Printf("\n So num := %d and num1:= %d", num, num1)

	swapping(&num, &num1)
	fmt.Printf("\n after swapping num:= %d and num1:=%d", num, num1)

}
