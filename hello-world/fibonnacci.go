package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b

		fmt.Println("The value of a:", a)
		fmt.Println("The value of b is:", b)
		fmt.Println("iteration i:", i)
	}

	return b
}

func main() {
	fmt.Println(fib(6))
}
