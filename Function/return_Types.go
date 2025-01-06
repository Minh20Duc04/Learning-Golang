package main

import "fmt"

var carInventory = make(map[string]int)

func main() {
	addToInventory("VinFast", 10)
	addToInventory("Sedan", 15)

	addToInventory("VinFast", 15)
	fmt.Println("Updated inventory:", carInventory)

	fmt.Println(getCount("VinFast"))

	fmt.Println(getStatus("VinFast"))
}

func addToInventory(bodyType string, number int) { //void
	carInventory[bodyType] += number
	fmt.Println(bodyType, "added. New count:", carInventory[bodyType]) //giống như get(Object) trong Java, là lấy giá trị value dựa trên key
} 

func getCount(bodyType string) int{ //return 
	return carInventory[bodyType]
}

func getStatus(bodyType string) (string, int){ //return 2 parameters
	carType := bodyType
	count := carInventory[bodyType]
	return carType, count
}