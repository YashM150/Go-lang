package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot be dividable")
	}
	return a / b, nil
}

func main() {
	fmt.Println("This is error handling demo")
	ans, _ := divide(10, 0)
	fmt.Println("Division :", ans)

}
