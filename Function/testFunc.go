package main

import "fmt"

func main() {
	// b1 := calculate1(3, 5, "*")
	// fmt.Println(b1)

	// b2 := operate(10,5, minus)
	// fmt.Println(b2)

	// b3 := []int{32,12,41,55,71}
	// fmt.Println(sum1(b3...))

	b4 := []string {"Hello", "Minh", "Duc"}
	fmt.Println(concat(b4...))
}

/*b1
Viết một hàm calculate1 nhận hai số nguyên và một chuỗi biểu diễn phép toán ("+", "-", "*", "/")
,sau đó trả về kết quả của phép toán.
*/
func calculate1(num1, num2 int, str string) int {
	switch str{
	case "+" : return plus(num1,num2)
	case "-" : return minus(num1, num2)
	case "*" : return multiply(num1, num2)
	case "/" : return divide(num1,num2)
	default : fmt.Println("Phep toan khong hop le", str)
	return 0
	}
}
func plus(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}
//-------------------------------------//

/*b2
Viết một hàm operate nhận vào một hàm khác như tham số (hàm này thực hiện phép toán trên hai số nguyên). 
Sau đó, operate sẽ gọi hàm này với hai số nguyên bất kỳ.
*/
func operate (num1, num2 int, op func(int,int) int) int{
	return op(num1,num2)
}
//-------------------------------------//

/*b3
Viết một hàm sum sử dụng variadic function để tính tổng của một danh sách các số nguyên.
*/
func sum1(nums... int) int{
	result := 0
	for _, i := range nums{
		result += i
	}
	return result
}
//-------------------------------------//

/*b4
Viết một hàm concat nhận nhiều chuỗi và trả về một chuỗi duy nhất bằng cách nối chúng lại với nhau, cách nhau bởi dấu cách.
*/
func concat(strings ... string) string{
	str := ""
	for _, i := range strings{
		str += i + " "
	}
	return str
}