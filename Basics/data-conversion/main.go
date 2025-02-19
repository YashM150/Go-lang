package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This is data conversion demo:")
	var num int
	fmt.Printf("\nEnter the number:")
	fmt.Scan(&num)

	fmt.Printf("\nIn float:\n Num:%f\nType of num:%T\n", float64(num), float64(num))
	fmt.Printf("\nIn String:\n Num:%s\nType of num:%T\n", strconv.Itoa(num), strconv.Itoa(num))

	var str string
	fmt.Printf("\nEnter the String:")
	fmt.Scan(&str)

	float1, _ := strconv.ParseFloat(str, 64)
	num1, _ := strconv.Atoi(str)
	fmt.Printf("\nIn Float:\n Num:%f\nType of num:%T\n", float1, float1)
	fmt.Printf("\nIn Integer:\n Num:%d\nType of num:%T\n", num1, num1)
}
