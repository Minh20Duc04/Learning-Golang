package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*Yêu cầu: Viết một hàm analyzeString(s string) (int, error) nhận một chuỗi s và trả về độ dài của chuỗi. Nếu chuỗi rỗng, hàm phải trả về một lỗi với thông báo "Empty string provided".
Nếu chuỗi chứa ký tự đặc biệt (chẳng hạn như dấu chấm, dấu chấm hỏi),
hàm phải trả về một lỗi với thông báo "Special characters detected".
*/
func analyzeString(s string) (int, error) { 
	if s == ""{
		return 0, errors.New("Empty string provided")
	}else if strings.ContainsAny(s,".!:?"){
		return 0, errors.New("Special characters detected")
	}
	return len(s), nil
}
/*Yêu cầu: Viết một hàm convertToInt(s string) (int, error) nhận một chuỗi s và chuyển đổi nó thành số nguyên.
 Nếu chuỗi không phải là một số hợp lệ (ví dụ: chứa ký tự không phải là số), 
 hàm phải trả về một lỗi với thông báo "Invalid number format".
 */
func convertToInt(s string) (int, error){
	result, err := strconv.Atoi(s)
	return result, err
}


func main() {
	str := "This is my week 3 of learning Go"
	length, err := analyzeString(str)

	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println(length)
	}

	str1 := "31"
	result, err := convertToInt(str1)
	if err != nil{
		fmt.Println("Invalid number format")
	}else{
		fmt.Println(result)
	}
	

}
