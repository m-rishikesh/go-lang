package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("error while creating a file")
	}
	defer file.Close()
	file.WriteString("created using golang")
	stream, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("error while reading a file")
	}
	fmt.Println(string(stream))
	webserver()
}
