package main

import (
	"fmt"
	"net"
)

func main() {
	GetMAC()
}

//获取本机的MAC地址

func GetMAC() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Err:" + err.Error())
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址，硬件地址
		fmt.Println("MAC=", mac)
	}
}
