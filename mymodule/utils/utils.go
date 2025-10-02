package utils

import (
	"errors"
	"fmt"
	"math/rand"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("NAME REQUIRED")
	}
	msg := fmt.Sprintln(randfunc(), name)
	return msg, nil
}

func randfunc() string {
	formats := [](string){
		"Hi, Welcome!",
		"Great to see you, !",
		"Hail,  Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func SayHelloArray(name []string) (map[string]string, error) {
	message := make(map[string]string, len(name))
	for _, val := range name {
		returnval, err := SayHello(val)
		if err != nil {
			return map[string]string{}, err
		}
		message[val] = returnval
	}
	fmt.Println("Running the sayhellowfunc")
	return message, nil
}
