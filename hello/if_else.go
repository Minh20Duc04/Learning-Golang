package main

import "fmt"

func main() {
	// isFast := true
	// if isFast {
	// 	fmt.Println("Going so fast mate, slow down!")
	// }else if !isFast {
	// 	fmt.Println("Ok this is acpectable!")
	// }


	//cho 2 bất biến lowerSpeedLimit và upperSpeedLimit, sau đó cho thử 1 speed để kiểm tra
	const upperSpeedLimit = 130
	const lowerSpeedLimit = 70
	var speedActual uint16
	fmt.Print("You 're driving at: ")
	fmt.Scanln(&speedActual)

	if speedActual < lowerSpeedLimit{
		fmt.Println("You 're going to slow, speed up mate!")
	}else if speedActual > upperSpeedLimit {
		fmt.Println("You 're going to fast, slow down mate!")
	}else{
		fmt.Println("Thank you for complying the law")
	}
}