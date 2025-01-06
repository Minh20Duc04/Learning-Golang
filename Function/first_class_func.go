package main

import (
	"fmt"
	"math"
)

//nhận hàm như 1 tham số, sẽ gồm tham số khác xong dùng hàm nhận để xử lý tham số đó
func main(){
	fmt.Println(calculate(7, double))
	fmt.Println(calculate(6, half))

	result := calculate(5, func(intputNum int) int { //anonymous func
		return int(math.Pow(float64(intputNum) , 2))
	})

	fmt.Println(result)
}

func calculate (num int,  operation func(intputNum int) int) int{ //cứ coi operation như 1 biến bth, tên + kdl
														          //func như kiểu khung sườn chứ k cố định là hàm nào
	return operation(num)
}

func double (num int) int{
	return num * 2
}

func half (num int) int{
	return num / 2
}