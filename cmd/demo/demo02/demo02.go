package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定義變量
	// var i int

	// 賦值
	// i = 10
	// fmt.Println("i =", i)
	// fmt.Printf("i = %d\n", i)

	// //系統會根據輸入值自動判定變數類型
	// var num = 10.11
	// //fmt.Println("num 的數據類型是 %T ", num)
	// fmt.Printf("num 的數據類型是 %T\n", num)

	// 簡潔語法
	// 透過以下語法同時處理定義與賦值
	// name := "tom"
	// fmt.Println("name=", name)
	// var address string = "台灣高雄85 大樓 \nhello world"
	// fmt.Println(address)

	// string2 := `這是一個\n 忽略器的實作`
	// fmt.Println(string2)

	//基本數據類型轉換的實作
	// var i int32 = 100

	// 將 i 轉成float 這樣就可以瘋狂捉i 了
	// var n1 float32 = float32(i)
	// var n2 int8 = int8(i)
	// fmt.Printf("i=%v n1=%v n2=%v\n", i, n1, n2)

	// 基本數據類型轉string 的應用
	// var num1 int = 99
	// var num2 float64 = 23.456
	// var b bool = true
	// var myChar byte = 'h'
	// var str string //空的str

	// 使用第一種方法來轉換
	// str = fmt.Sprintf("%d", num1)
	// fmt.Printf("str type %T str=%v \n", str, str)

	// str = fmt.Sprintf("%f", num2)
	// fmt.Printf("str type %T str=%v \n", str, str)

	// str = fmt.Sprintf("%t", b)
	// fmt.Printf("str type %T str=%q \n", str, str)

	// str = fmt.Sprintf("%c", myChar)
	// fmt.Printf("str type %T str=%q \n", str, str)

	// //第二種方式strconv
	// var str string
	// var num3 int = 99
	// var num4 float64 = 23.456
	// var b2 bool = true
	// // var str string //空的str
	// str = strconv.FormatInt(int64(num3), 10)
	// fmt.Printf("str type %T str=%q\n", str, str)

	// str = strconv.FormatFloat(num4, 'f', 10, 64)
	// fmt.Printf("str type %T str=%q\n", str, str)

	// str = strconv.FormatBool(b2)
	// fmt.Printf("str type %T str=%q\n", str, str)

	//將 string 轉換為基本數據類型
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "1234567890"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("n1 type %T n1=%v\n", f1, f1)

}
