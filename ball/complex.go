package main

import (
	"fmt"
)

func main() {
	fmt.Println("Complex:")

	c := 10 + 11i

	realp := real(c)
	imagp := imag(c)

	fmt.Printf("%v and %v\n", realp, imagp)

	d := complex(10, 2)

	fmt.Println(c + d)
}
