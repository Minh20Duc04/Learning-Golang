package main

import "fmt"

func main() {

	//địa chỉ và con trỏ
	// x := 35
	// var ptr *int = &x
	// fmt.Println(*ptr)

	//-------------------------------
	// var name string
	// fmt.Println("What's your name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Hello", name)

	str := "Hello Go"
	sayHello(str)
	fmt.Println("Method sayHello: ", str)

	pointerSayHello(&str)
	fmt.Println("Method pointerSayHello: ", str)

}

func sayHello(s string){
	s = "Hello World"
}

func pointerSayHello(p *string){
	*p = "Hello World"
}
