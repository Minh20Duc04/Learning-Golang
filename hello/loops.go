package main

import "fmt"

func main() {

	//for: 
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Simple loop:", i)
	// }
	
	//while:
	i := 1
	for{
		fmt.Println("While loop", i)
		if i == 5{
			break
		}
		i++
	}
}