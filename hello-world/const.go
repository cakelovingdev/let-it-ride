package main

import "fmt"

func main() {
	fmt.Println("Hello..")

	const (
		a = 10
		b
		c
	)

	fmt.Println(a, b, c)

	const d int = 12
	fmt.Println(d)

	const (
		x = iota
		w = iota
		y
		z
	)

	fmt.Println(x, y, z, w)
}
