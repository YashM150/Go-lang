package main

import "fmt"

func main() {
	fmt.Println("This is map demo:")
	grades := make(map[string]int)
	var choice int
	counter := 0
OuterLoop:
	for {
		counter++
		fmt.Printf("\n\n-------------------Menu--------------------------\n")
		fmt.Printf("1. Add marks\n 2. View the marks\n 3.Particular student marks\n 4. Delete the student\n 5. Break the loop\n")
		fmt.Printf("\nEnter the choice:\t")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if len(grades) == 0 {
				var len int
				fmt.Printf("Enter the no of students:\t")
				fmt.Scan(&len)
				for i := 0; i < len; i++ {
					var name string
					var marks int
					fmt.Printf("\nEnter name:\t")
					fmt.Scan(&name)
					fmt.Printf("\nEnter marks:\t")
					fmt.Scan(&marks)
					grades[name] = marks
				}
			} else {
				var name string
				var marks int
				fmt.Printf("\nEnter name:\t")
				fmt.Scan(&name)
				fmt.Printf("\nEnter marks:\t")
				fmt.Scan(&marks)
				grades[name] = marks
			}
		case 2:
			fmt.Println("Marks are:")
			for value, index := range grades {
				fmt.Printf("\nMarks for %s are %d", value, index)
			}
		case 3:
			var name string
			fmt.Printf("\nEnter the name:\t")
			fmt.Scan(&name)

			value, exists := grades[name]

			if exists {
				fmt.Printf("\nMarks of %s are %d", name, value)
			} else {
				fmt.Printf("No marks for %s", name)
			}
		case 4:
			var name string
			fmt.Printf("\nEnter the name:\t")
			fmt.Scan(&name)

			value, exists := grades[name]
			if exists {
				fmt.Printf("\nMarks of %s i.e %d is deleted!", name, value)
				delete(grades, name)
			}
		case 5:
			if counter >= 1 {
				break OuterLoop
			}
		default:
			fmt.Println("Enter correct choice.")
		}
	}

}
