package main

import "fmt"

func main() {

	var j int

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for j = 0; j <= 5; j++ {
		if j == 0 {
			continue
		} else if j == 4 {
			break
		}
		fmt.Println(j)
	}

	var i int

	fmt.Println("No way!", i, j)

	// range
	var cakes = [...]string{"shortcake", "cheesecake", "tiramisu"}

	for _, values := range cakes {
		fmt.Println("Cakes:", values)
	}

	// while loop??
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("After while loop i value is:", i)
	i = 0
	fmt.Println("Now i value is:", i)

	// do while loop??
	for {
		fmt.Println(i)
		i += 1
		if i < 10 {
			break
		}
	}

}
