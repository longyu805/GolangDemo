package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Copy()
	//CopyN()
	//ReadFull()
	WriteString()
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

//将N个字符，复制到Write1.txt文件当中
func CopyN() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("Write1.txt")
	num, err := io.CopyN(w, r, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

//是读取len(buf)个，放在buf中
func ReadFull() {
	r, _ := os.Open("write.txt")
	b := make([]byte, 20)
	num, err := io.ReadFull(r, b)
	defer r.Close()
	fmt.Println(num)
	if err != nil {
		fmt.Println(err)
		if err == io.EOF {
			fmt.Println("Read end total", num)
		}
		if err == io.ErrUnexpectedEOF {
			fmt.Println("Read fewer value:", string(b[:num])) //Read fewer value: hello widuu，依然是buf长度大于读取的长度

			return
		}
	}
	fmt.Println("Read value:", string(b))
}

//向写入目标中写入字符创，返回是写入的字节数还有error错误
func WriteString() {
	w, _ := os.OpenFile("write1.txt", os.O_RDWR, os.ModePerm)
	n, err := io.WriteString(w, "ni hao ma")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
