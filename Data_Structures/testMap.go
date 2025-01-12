package main

import "fmt"

func main() {

	//1. Đếm tần suất các phần tử trong mảng
	/* 
	vd: [1, 2, 2, 3, 3, 3] 
	->
	1: 1
	2: 2
	3: 3
	*/
	// inputString := []int{1, 2, 2, 3, 3, 3}
	// countMap := make(map[int]int)

	// //lấy giá trị ko phải vị trí của string nên bỏ qua vị trí (_)
	// for _,num := range inputString{
	// 	//kiểm tra key xem có thêm vào countMap chưa
	// 	if _, ok := countMap[num]; ok{
	// 		countMap[num]++
	// 	}else{
	// 		countMap[num] = 1
	// 	}
	// }

	// for key, value := range countMap{
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	//2. Đảo ngược map
	// Yêu cầu:
	// Tạo một map với khóa là tên của các quốc gia và giá trị là mã quốc gia (ví dụ: {"Vietnam": "VN", "United States": "US"}).
	// Đảo ngược map để mã quốc gia trở thành khóa và tên quốc gia trở thành giá trị.

	// map1 := map[string]string{
	// 	"Vietnam" : "VN",
	// 	"United States" : "US",
	// }

	// fmt.Println("Map before reversed: ", map1)
	// map2 := make(map[string]string)
	// for key, value := range map1{
	// 	map2[value] = key
	// }
	// fmt.Println("After reversed: ", map2)

	//3. Tính tổng điểm sinh viên
	// Yêu cầu:
	// Tạo một map với tên sinh viên làm khóa và mảng điểm của họ làm giá trị.
	// Sử dụng for-range loop để tính tổng điểm của mỗi sinh viên.
	// In ra tên sinh viên và tổng điểm.

	map1 := map[string][]int{
		"Alice": {85, 90, 78},
    	"Bob":   {88, 72, 95},
	}
	for k, v := range map1{
		var temp int = 0
		for _, loopInMap := range v{
			temp += loopInMap
		}
		fmt.Printf("%v: %v\n", k, temp)
	}
	

}