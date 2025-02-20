package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("This is start of sayHello function")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("This is end of sayHello function")
}

func sayHi() {
	fmt.Println("This is start of sayHi function")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("This is end of sayHi function")
}

func main() {
	fmt.Println("This is go routine demo")

	go sayHello()
	go sayHi()

	time.Sleep(800 * time.Millisecond)
}
