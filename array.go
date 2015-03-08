package main

import (
	"fmt"
)

func main() {
	/*a := [2][3]int{
		{1, 3, 4},
		{12, 13, 14}}
	fmt.Println(a)*/
	Sort()
}

func Sort() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
