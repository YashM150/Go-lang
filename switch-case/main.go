package main

import (
	"fmt"
	myutil "learngo/myutils"
)

func main() {
	myutil.PrintMessage("This is Switch case demo:")
	var day int
	fmt.Println("Enter the number:\n 1.Monday\n 2.Tuesday\n3.Wednesday\n4.Thursday\n5.Friday\n6.Saturday\n7. Sunday")
	fmt.Printf("Enter:\t")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Enter correct day!")
	}
}
