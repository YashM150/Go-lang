package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Your Name is:", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Entered Text is:", name)
}
