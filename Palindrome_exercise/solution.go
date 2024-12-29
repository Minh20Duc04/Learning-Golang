package main

import (
	"fmt"
)

func main() {
	//Nhập vào 1 chuỗi, kiểm tra xem nó có phải chuỗi đối xứng hay không
	var inputString string
	fmt.Println("Enter your string to check: ")
	fmt.Scanln(&inputString)

	fmt.Println("Result of your input string is :", checkPalindrome(inputString))
}

func checkPalindrome(strCheck string) bool {
	// var tempCheck string

	// i:=0
	// for{
	// 	check1 := strCheck[i]
	// 	check2 := strCheck[len(strCheck) - i]

	// 	if check1 != check2{
	// 		break
	// 	}

	// 	if i > len(strCheck) -1 {
	// 		break;
	// 	}
	// 	i++
	// }

	// return tempCheck == strCheck


	//Cách trên chưa đúng
	i := 0
	for i < len(strCheck)/2 {
		check1 := strCheck[i]
		check2 := strCheck[len(strCheck)-1-i]
		
		if check1 != check2 {
			// Nếu các ký tự không khớp, trả về false
			return false
		}
		
		i++
	}
    
    return true
}
