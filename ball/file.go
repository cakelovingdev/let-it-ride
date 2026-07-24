package main

import (
	"fmt"
	"os"
)

const path = "ball/hehe.txt"

func CreateFile() {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error")
	}
	defer file.Close()
}

func FileExists() bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func main() {
	if FileExists() == true {
		fmt.Println("file exists")
	} else if FileExists() == false {
		fmt.Println("file doesnt exists")
		CreateFile()
	}
}
