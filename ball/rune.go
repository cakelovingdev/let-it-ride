package main

import (
	"fmt"
)

func main() {
	fmt.Println("Runes \n-----\n")

	var msg string = "Happy Birthday! 🍰"

	for i, r := range msg {
		fmt.Printf("\nThe Index is %v and the rune is %v", i, r)
	}
	fmt.Print("\n")

	cake := 127856
	fmt.Println(string(rune(cake)))

	woah := '🎉'
	fmt.Printf("%v is the type of: %T", woah, woah)
}
