package main

import "fmt"

func main() {
	fmt.Println("This is arrays demo")
	fmt.Println("This is integer arrays:")
	var arr [5]int

	arr[0] = 1
	arr[3] = 4

	fmt.Println("This is String arrays:")
	var arr1 [5]string

	arr1[0] = "Yash"
	arr1[3] = "Moharil"

	fmt.Printf("The array is:%q\n", arr1)
}
