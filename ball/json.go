package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	PageCount int
	id        int // not exported since it doesnt start with a uppercase letter
}

func main() {

	fmt.Println("JSON")

	var book1 Book

	book1.Title = "The God of Small Things"
	book1.Author = "Arundhati Roy"
	book1.PageCount = 339
	book1.id = 2

	book2 := Book{
		Title:     "Not Quite Dead Yet",
		Author:    "Holly Jackson",
		PageCount: 429,
		id:        1,
	}

	var book3 Book

	book3 = Book{
		Title:     "Days at the Torunka Cafe",
		Author:    "Satoshi Yagisawa",
		PageCount: 229,
		id:        3,
	}

	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)

	book1_out, _ := json.Marshal(book1)
	book2_out, _ := json.Marshal(book2)
	book3_out, _ := json.Marshal(book3)

	fmt.Println(string(book1_out))
	fmt.Println(string(book2_out))
	fmt.Println(string(book3_out))

	var out = make([]Book, 0)

	out = append(out, book1)
	out = append(out, book2)
	out = append(out, book3)

	fmt.Println(out)

	all_out, _ := json.Marshal(out)
	fmt.Println(string(all_out))
}
