package main

import (
	"fmt"
	"strconv"
)


func main()  {
	// //doi tu độ K - > C
	// temperatureK := 346.3
	// temperatureC := int8(temperatureK - 273.15)

	// fmt.Println(temperatureC)

	// //làm tròn bằng math
	// var num float64 = 12.75
	// fmt.Println("Num: ", num)
	// num = math.Round(num)
	// fmt.Println("Num after round: ", num)

	//------------------------------------------------

	//chuyển chuỗi thành số
	// num := "28"
	// numToString, err := strconv.Atoi(num)
	// if(err != nil){
	// 	fmt.Println("Error with: ", err)
	// }else{
	// 	fmt.Println("Convert to string:", numToString)
	// 	fmt.Printf("%T \n", numToString)
	// }

	//------------------------------------------------
	var myFloatStr string = "3.143826"
	myFloatFromString, _ := strconv.ParseFloat(myFloatStr, 64)
	fmt.Printf("%T \n", myFloatFromString)


	

}