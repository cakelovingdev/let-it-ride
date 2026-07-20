package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps\n----\n")

	var cards = map[string]int{"Ace": 13, "Spades": 13, "Diamonds": 13, "Hearts": 13}

	fmt.Println(cards)

	v1 := cards["Ace"]
	v2, o2 := cards["Spades"]
	_, o3 := cards["Diamonds"]

	fmt.Println(v1)
	fmt.Println(v2, o2)
	fmt.Println(o3)

	delete(cards, "Ace")

	fmt.Println(cards)

	cards["Clover"] = 13

	fmt.Println(cards)

	var emptymap = map[string]int{}

	fmt.Println(emptymap)

	// empty map using make function
	var a = make(map[string]int) // preferred way to create empty maps

	fmt.Println(a)

	var b map[string]int

	fmt.Println(b)

	if a == nil {
		fmt.Println("a is nil")
	}

	if b == nil {
		fmt.Println("b is nil")
	}

	if emptymap == nil {
		fmt.Println("empty map is nill")
	}

	// b["Car"] = 1
	// fmt.Println("this line doesnt run because of write panic?")

	b = make(map[string]int)

	fmt.Println("This runs now:", b)

	for key, val := range cards {
		// fmt.Println(key, "=", val)
		fmt.Printf("%v = %v \n", key, val)
	}

}
