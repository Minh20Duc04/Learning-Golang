package main

import "fmt"

func main() {
	//tạo mảng gán giá trị luôn, sau đó tạo mảng mới từ index của mảng cũ

	fuelTypes := []string{"Gasoline", "Diesel", "Electric", "Hybrid", "Hydrogen"}

	fuelTypesCopy := fuelTypes[0:2]
	fmt.Println("fuelTypesCopty", fuelTypesCopy)

}