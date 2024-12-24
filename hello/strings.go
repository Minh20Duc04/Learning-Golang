package main

import (
	"fmt"
	"strings"
)

func main() {

	// str := "Nguyen Minh Duc la toi, oki"
	// lenght := len(str)
	// fmt.Println(lenght)

	// str1 := "Go Programing"
	// str2 := "go programing"

	// //kiểm tra đúng định dạng UTF-8
	// checkString := strings.EqualFold(str1, str2)
	// fmt.Println(checkString)

	// //kiểm tra vị trí string tại index
	// fmt.Println(strings.Index(str1, "gram"))

	// //tạo mảng phụ dựa trên index của 1 string cho trước
	// str4 := "Hello, World"
	// str5 := str4[7:12] //Word
	// str5 = str4[4:]    //, World
	// str5 = str4[:4]    //Hell
	// fmt.Println(str5)

	//replace:
	// str := "Hello, Go"
	// str = strings.Replace(str, "Go", "World", 1)
	// fmt.Println(str)

	str1 := "Go Programing"
	fmt.Println(strings.ToUpper(str1))	
	fmt.Println(strings.ToLower(str1))	

	fmt.Println(strings.Contains(str1, "Go"))
}