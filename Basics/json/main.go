package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	var person Person
	fmt.Println("This is json demo:")
	fmt.Printf("\n Enter the details of the person:")

	fmt.Printf("\n 1. Enter the name:\t")
	fmt.Scan(&person.Name)
	fmt.Printf("\n 2. Enter the Age:\t")
	fmt.Scan(&person.Age)
	fmt.Printf("\n 3. Is the person Adult:\t")
	var choice string
	fmt.Scan(&choice)
	if choice == "yes" {
		person.IsAdult = true
	} else if choice == "no" {
		person.IsAdult = false
	} else {
		fmt.Println("Enter correct choice!!")
		return
	}

	// encoding the json
	fmt.Println("Converting to json:")
	jsondata, err := json.Marshal(person)
	if err != nil {
		fmt.Println("There is error while converting to json", err)
		return
	}
	fmt.Println("The json data is:", string(jsondata))

	//decoding the json
	fmt.Println("Decoding the json")
	var person1 Person
	errors := json.Unmarshal(jsondata, &person1)
	if errors != nil {
		fmt.Println("The error is:", err)
		return
	}
	fmt.Println("The decoded value is:", person1)

}
