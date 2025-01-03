package main

import "fmt"

func main() {
	// var arrayTemp [3]string // hoặc này, var chưa cần gán giá trị liền
	// arrayTemp[0] = "Nguyen"
	// arrayTemp[1] = "Minh"
	// arrayTemp[2] = "Duc"
	// hoặc arrayTemp = {"Nguyen", "Minh", "Duc"}

	// num := []int{2,5,6,123,561,87} // hoặc này, := phải gán giá trị liền
	// fmt.Println(num)

	// fmt.Println(arrayTemp)

	// for i := 0; i < len(arrayTemp); i++ {
	// 	fmt.Print(arrayTemp[i])
	// }

	//--------------------------------------------
	arrays := []int{2,4,6,8,10}
	fmt.Println("Arrays before add: ", arrays)

	arrays = append(arrays, 4, 321,414,321)
	fmt.Print("Arrays after add: ", arrays)

}