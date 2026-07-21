package main

import (
	"fmt"
)

func isItEqual(x float64, y float64) bool {
	const diff = 1e-9

	var mag float64

	if x <= y {
		mag = y - x
	} else {
		mag = x - y
	}

	fmt.Println(mag)
	fmt.Println(diff)

	if mag <= diff {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Printf("// Test - to show the %% \n")

	var a float64 = 2
	var b float64 = 2.11

	fmt.Printf("%.2f \n", a+b)

	if isItEqual(a+b, 4.11) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
