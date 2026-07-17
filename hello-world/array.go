
package main

import "fmt"

func main(){
  fmt.Println("Array")

  var vowels = [5]string{"a","e","i","o","u"}
  
  fmt.Println(vowels)

  fmt.Println(vowels[1])
  fmt.Println(vowels[1+2])
  
  i := 2
  fmt.Println(vowels[i])

  var fav_food = [...]string{"cake","sweets","chocolate"}
  fmt.Println(fav_food)


  fav_food[2] = "stew"
  fmt.Println(fav_food[1+1])

  var consonants = []string{1:"b",2:"c",3:"d",5:"f",6:"g",7:"h",9:"j",10:"k",11:"l",12:"m",13:"n",15:"p",16:"q",17:"r",18:"s",19:"t",21:"v",22:"w",23:"x",24:"y",25:"z"}

  fmt.Println(consonants)

  fmt.Println("The no of vowels are", len(vowels))
  fmt.Println("The no of consonants are", cap(consonants))
  fmt.Println(len(fav_food))

  fmt.Println(cap(consonants))


  var const_slice = consonants[0:3]

  fmt.Println(const_slice)
  fmt.Println(len(const_slice), cap(const_slice))


  vowels_extended := vowels[0:5]
  fmt.Println(vowels_extended, "Capacity:", cap(vowels_extended), "Length:", len(vowels_extended))
  

  vowels_extended = append(vowels_extended, "vow", "els")

  fmt.Println("The new ext:", vowels_extended, "Capacity:", cap(vowels_extended), "Length:", len(vowels_extended))
 
  var arr1 = [...]int{1,2,3,4,5}
  var arr2 = [...]int{6,7,8,9}

  // cant append arrays so converting them to slices
  var arr3 = append(arr1[:], arr2[:]...)

  fmt.Println(arr3)

}
