package main

import (
	"fmt"
)

func main() {
	//tạo mảng trước rồi duyệt tuần tự
	bodyTypes := []string{"Sedan", "SUV", "Converible"}

	for i, bodyTypes := range bodyTypes { //i là chỉ số
		fmt.Printf("Index: %v. Item: %v \n", i, bodyTypes)
	}

	carInventory := map[string]int{
		"Sedan" : 20,
		"SUV" : 12,
		"Converible" : 16,
	}

	for i, carInventory := range carInventory{ //i là key
		fmt.Printf("This %v has %v \n",i,carInventory)
	}

	countInventory := 0

	for _, count := range carInventory{ //bỏ qua key
		countInventory += count
	}

	fmt.Println(countInventory)

	
}