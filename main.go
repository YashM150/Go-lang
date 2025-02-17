package main

import (
	"fmt"
	myutil "learngo/myutils"
)

func main() {
	fmt.Println("Learn go")
	myutil.PrintMessage("Myutil: Hello World")

	var num int = 1
	var name string = "Yash"
	var numb float64 = 20.21
	var status bool = false
	rando := "yash"
	fmt.Println("Variables", num, name, numb, status, rando)
}
