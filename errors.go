package main

import (
	"errors"
	"fmt"
)

func main() {
	Errors()
}

func Errors() {
	err := errors.New("error golang")
	if err != nil {
		fmt.Println(err)
	}
}
