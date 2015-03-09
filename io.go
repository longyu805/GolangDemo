package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Copy()
}

//这个函数是从一个文件读取拷贝到另外一个文件，一直拷贝到读取文件的EOF，所以不会返回io.EOF错误，参数是写入目标器和读取目标器，返回int64的拷贝字节数和err信息
func Copy() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("Write.txt")
	num, err := io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num) //返回int64的11 打开我的write.txt正是test.txt里边的hello widuu
}
