package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
	fmt.Println("This is file operation demo:")
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("There is error while creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Hello world"
	// _, error := io.WriteString(file, content)
	// if error != nil {
	// 	fmt.Println("There is error while writing in file:", err)
	// }
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("There is error while parsing the file:", err)
	// 	return
	// }
	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println("Error while parsing the file:", err)
	// 		}
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("There is err while parsing", err)
	}
	fmt.Println("Contents of file are:", string(content))

	// defer file.Close()
	// fmt.Println("File is created!")
}
