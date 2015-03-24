package main

import (
	"fmt"
)

type data struct {
	a int
}

func main() {
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p,%v \n", p, p.a)
}
