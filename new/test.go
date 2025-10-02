package main

import (
	"fmt"
	"log"

	"github.com/m-rishikesh/go-lang/mymodule/utils"
)

func Mymod() {
	mapp, err := utils.SayHelloArray([]string{"ram", "sita", "geeta"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mapp)

}
