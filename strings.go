package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//ToLower()
	//ToUpper()
	//ToUpperSpecial()
	CrossLineString()
}

//大写转小写
func ToLower() {
	fmt.Println(strings.ToLower("ABC"))
}

//小写转换大写
func ToUpper() {
	fmt.Println(strings.ToUpper("abc"))
}

//该函数把s字符串里面的每个单词转化为大写，但是调用的是unicode.SpecialCase的ToUpper方法
func ToUpperSpecial() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToUpperSpecial(SC, "Gopher"))
}

//字符串跨行连接，“+”必须在上一个字符串结尾
func CrossLineString() {
	s := "Hello" +
		"World!"
	fmt.Println(s)

	/*
		s2 := "Hello"
		+"World!"
		fmt.Println(s2)
			invalid operation: + untyped string
	*/
}
