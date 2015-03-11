package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	//Copy()
	//CopyN()
	//ReadFull()
	//WriteString()
	//LimitedReader()
	//Chtimes()
	//Environ()
	//Exit()
	Hostname()
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
	defer w.Close()
	b, _ := ioutil.ReadFile("write1.txt")
	fmt.Println("write total", n)
	fmt.Println(string(b))
}

/*
结构type LimitedReader
查看源代码打印帮助
type LimitedReader struct {
    R Reader // 读取器了
    N int64  // 最大字节限制
}
*/

func LimitedReader() {
	reader, _ := os.Open("test.txt")
	limitedreader := io.LimitedReader{
		R: reader,
		N: 20, //，不能缺少
	}
	p := make([]byte, 10)
	var total int
	for {
		n, err := limitedreader.Read(p)
		if err == io.EOF {
			fmt.Println("read total", total)
			fmt.Println("read value", string(p))
		}
		total = total + n
	}
}

//访问时间 创建时间 返回的是error接口信息
func Chtimes() {
	err := os.Chtimes("io.go", time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fi, _ := os.Stat("io.go")
	fmt.Println(fi.ModTime())
}

//Environ()获取系统的环境变量
func Environ() {
	data := os.Environ()
	fmt.Println(data)
}

func Exit() {
	for {

		fmt.Println("Exit")
		os.Exit(1)
	}
}

//
func Hostname() {
	hosename, _ := os.Hostname()
	fmt.Println(hosename)
}
