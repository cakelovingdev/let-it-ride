package operators

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var c int = 2

	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	fmt.Println("c is:", c)

	// Arithmetic Operators
	fmt.Println("a+b:", a+b)
	fmt.Println("a-b:", a-b)
	fmt.Println("a*b:", a*b)
	fmt.Println("b/a:", b/a)
	fmt.Println("a%b:", a%b)
	fmt.Println("a|b:", a|b)
	fmt.Println("a^b:", a^b)
	fmt.Println("a<<b:", a<<b)
	fmt.Println("a>>b:", a>>b)

	// Comparision Operators
	fmt.Println("a==b:", a == b)
	fmt.Println("a!=b:", a != b)
	fmt.Println("a>b:", a > b)
	fmt.Println("a<b:", a < b)
	fmt.Println("a>=b:", a >= b)
	fmt.Println("a<=b:", a <= b)

	fmt.Println("c==b:", c == b)
	fmt.Println("c!=b:", c != b)
	fmt.Println("c<=b:", c <= b)
	fmt.Println("c>=b:", c >= b)
}
