package main

import (
	"fmt"
)

func main() {
label1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break label1
			}
			fmt.Println(i)
		}
	}
	fmt.Println("End")
}
