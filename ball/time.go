package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Printf("%T is the type and the value is %v\n\n", now, now)
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Printf("Time formatted in 2006-01-02 15:04:05 format is:%v", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Time formatted in 03:04PM format is:%v", now.Format(time.Kitchen))
}
