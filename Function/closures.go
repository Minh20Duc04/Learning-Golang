package main

import "fmt"

func main() { //Closure dùng để tạo trạng thái riêng, giải quyết các bài toán cần lưu trữ 
	nextInt := sequence() 
	//nextInt lúc này có thể là kiểu trả về hàm
	
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3
	fmt.Println(nextInt()) //4

	newInts := sequence()
	newInts() //2
	fmt.Println(newInts()) 


}

func sequence() func() int { //kiểu trả về hàm ẩn danh
	i := 0
	return func() int {
		i++
		return i
	}
}