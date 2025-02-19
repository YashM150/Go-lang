package main

import "fmt"

type student struct {
	Firstname string
	Lastname  string
	Marks     int
	Grades    string
}
type contact_details struct {
	Phone string
	Email string
}

type details struct {
	Academic_details   student
	Additional_details contact_details
}

func main() {
	fmt.Println("This is struct demo:")
	fmt.Printf("\nEnter the no of students:\t")
	var nstudents int
	fmt.Scan(&nstudents)

	arr := make([]details, nstudents)

	fmt.Printf("\nEnter the details of students:")
	for i := 0; i < nstudents; i++ {
		fmt.Printf("\nFor student %d", i+1)
		fmt.Printf("\nEnter first name:\t ")
		fmt.Scan(&arr[i].Academic_details.Firstname)
		fmt.Printf("Enter last name:\t ")
		fmt.Scan(&arr[i].Academic_details.Lastname)
		fmt.Printf("Enter email:\t ")
		fmt.Scan(&arr[i].Additional_details.Email)
		fmt.Printf("Enter phone number:\t ")
		fmt.Scan(&arr[i].Additional_details.Phone)
		fmt.Printf("Enter Marks:\t ")
		fmt.Scan(&arr[i].Academic_details.Marks)
		if arr[i].Academic_details.Marks >= 80 {
			arr[i].Academic_details.Grades = "A"
		} else if (arr[i].Academic_details.Marks >= 40) && arr[i].Academic_details.Marks < 80 {
			arr[i].Academic_details.Grades = "B"
		} else {
			arr[i].Academic_details.Grades = "C"
		}
	}

	fmt.Printf("\n\nEntered details are:")
	for i := 0; i < nstudents; i++ {
		fmt.Printf("\n\nDetails of Student %d:", i+1)
		fmt.Printf("\n1. First Name: %s", arr[i].Academic_details.Firstname)
		fmt.Printf("\n2. Last Name: %s", arr[i].Academic_details.Lastname)
		fmt.Printf("\n3. Email:%s", arr[i].Additional_details.Email)
		fmt.Printf("\n4. Phone:%s", arr[i].Additional_details.Phone)
		fmt.Printf("\n5. Marks: %d", arr[i].Academic_details.Marks)
		fmt.Printf("\n6. Grades: %s", arr[i].Academic_details.Grades)

	}

}
