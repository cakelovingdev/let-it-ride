package main

import "fmt"
import "rsc.io/sampler"


func add(x string, y string){
  fmt.Println(y,x)
}

func main() {

  // declaring variable without value
  var (
    a int
    b int
  )

  // Variable declaration - Inferred
  name := "kaoruko"

  // variable declared normally with type
  var hello string = "konichiwa"

  // declaring multiple variables
  c, d := "Hello", "World"


  // adding variables
  fmt.Println(a+b)
  fmt.Println(c+"\n"+d)

  // add variable to string
  fmt.Println("", hello)
  fmt.Println(""+hello)

  fmt.Print(name)
  fmt.Println(sampler.Hello())
  
  add("rintaro","tsumugi")


  // storing a result in a var
  e := c+d

  fmt.Println(e)

  // incrementing a var
  a = a + 1  
  fmt.Println(a)

  fmt.Println(a+b)

  fmt.Printf("%8s\n", name)
  fmt.Printf("%-8s\n", name)
  fmt.Printf("%16s\n", name) 

  b = 18
  fmt.Printf("%bs\n", a) // shows binary

  var h8 int8 = 129 - 2
  fmt.Println(h8)
} 
