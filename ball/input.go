package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a number base-10 to get the base-2 version of it:")

	var x int
	var slice = make([]int, 0)

	fmt.Scanf("%d", &x)

	var copy_x = x

	for i := 0; x != 0; i++ {
		if x%2 == 0 {
			// fmt.Print(0)
			slice = append(slice, 0)
		} else {
			// fmt.Print(1)
			slice = append(slice, 1)
		}
		x /= 2
	}
	// fmt.Println("\n", slice)

	// fmt.Println("the len of slice is:", len(slice))

	var reverse_slice = make([]int, 0)

	for j := len(slice) - 1; j >= 0; j-- {
		reverse_slice = append(reverse_slice, slice[j])
		// fmt.Println(slice[j], j)
	}

	fmt.Printf("\nbase 2 version of %v: %v", copy_x, reverse_slice)

}
