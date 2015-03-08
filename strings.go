package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	ToLower()
	ToUpper()
	ToUpperSpecial()
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
