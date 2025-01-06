package main

import "fmt"

func main() {
	tc1 := sum(1, 2, 4, 6, 2, 5, 76)
	fmt.Println(tc1)

	tc2 := []int {2,1,41,120,77,58}
	fmt.Println(sum(tc2...))

}

func sum(numbers ...int) int { //nhận nhiều tham số nếu không cần mảng
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
