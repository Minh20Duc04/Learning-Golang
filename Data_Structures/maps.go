package main

import "fmt"

func main() {
	// map1 := map[int]string{1:"Mot", 2:"Hai"} //phải khai báo khi mảng ko rỗng
	// fmt.Println(map1)

	// map1[3] = "Ba" //cách thêm vào map
	// fmt.Println(map1)

	carInventory := make(map[string]int) //tạo mảng rỗng nên ko cần khai báo
	carInventory["Sedan"] = 25 //thêm
	carInventory["SUV"] = 15
	carInventory["Converible"] = 10

	fmt.Println("Car inventory: ", carInventory)

	numOfSedan := carInventory["Sedan"] //lấy value từ key
	fmt.Printf("We have %v Sedan \n", numOfSedan)

	delete(carInventory, "SUV") //xóa dựa trên key
	fmt.Println("Car inventory:", carInventory)

	if numConverible, ok := carInventory["Converible"]; ok{
		fmt.Printf("We have %v Converible", numConverible)
	}
}