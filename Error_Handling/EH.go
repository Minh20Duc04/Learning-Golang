package main

import (
	"errors"
	"fmt"
)

func main() { 
	// temp, err := divide1(10,0)
	// if err != nil{
	// 	fmt.Println("Error:", err)
	// }else{
	// 	fmt.Println(temp)
	// }


	id := findItem(2)
	if errors.Is(id, errNotFound){
		fmt.Println("Item not found") //kiểm tra xem có phải lỗi đã định nghĩa trước không
	}


}

func divide1(a,b int) (int, error){ //error-handling thường đc xem như kiểu trả về thứ 2 	
	if b == 0{
		return 0, errors.New("Divide by zero")
	}
	return a/b, nil
}

//---------------------------------------------//
var errNotFound = errors.New("not found") 

func findItem(id int) error{
	if id != 1{
		return errNotFound
	}
	return nil
}