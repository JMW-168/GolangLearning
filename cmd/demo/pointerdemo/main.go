package main

import (
	"fmt"
)

func main() {
	var i int = 10
	fmt.Println("i 的地址＝", &i)

	//下面的  var ptr *int = &i
	//1. ptr 是一個指針變量
	//2. ptr 的類型 *int
	//3. ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr 的數值=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)
}
