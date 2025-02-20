package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("This is web demo:")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("There is error while getting the response:", err)
		return
	}
	fmt.Println("Hrre is response from the api:", res)

	defer res.Body.Close()

	data, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println("There is error while reading the body", error)
		return
	}
	fmt.Println("The response body is:", string(data))

}
