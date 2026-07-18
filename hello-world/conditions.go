package main

import "fmt"

func main() {
	var x int = 1
	var y int = 2

	var age int = -18 // uint can be used
	var name string = "kaoruko"

	if x < y {
		fmt.Println("Hell yea!")
	}

	if age >= 18 {
		if name == "kaoruko" {
			fmt.Println("Good to go")
		} else {
			fmt.Println("Name doesnt match")
		}

	} else if age < 18 && age >= 0 {
		fmt.Println("you are not an adult")
	} else {
		fmt.Println("Not a valid age!")
	}

	// switch case

	var mark uint = 99
	var grade string

	if mark >= 90 && mark <= 100 {
		grade = "A"
	} else if mark >= 80 && mark < 90 {
		grade = "B"
	} else if mark >= 70 && mark < 80 {
		grade = "C"
	} else if mark >= 60 && mark < 70 {
		grade = "D"
	} else if mark >= 50 && mark < 60 {
		grade = "E"
	} else {
		grade = "F"
	}

	switch grade {
	case "A":
		fmt.Println("Your Grade is A")
	case "B":
		fmt.Println("Your Grade is B")
	case "C":
		fmt.Println("Your Grade is C")
	case "D":
		fmt.Println("Your Grade is D")
	case "E":
		fmt.Println("Your Grade is E")
	default:
		fmt.Println("Your Grade is F")
	}
}
