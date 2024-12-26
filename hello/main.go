package main

import (
	"fmt"
)

func main() {
    // var name string
    // name = "Nguyen Duc"
    // fmt.Println("This is ", name, "speaking !")

    // numOfCars := 50
    // fmt.Println("Our fleet consists of ", numOfCars, " cars")
    // price := 29.99
    // fmt.Printf("Our starting price is %v\n", price)
    // fmt.Println("Take your pick")

    //------------------------------------------------
    //bất biến: 
    const Agency = "Fast Track"
    fmt.Println(Agency)

    //iota:
    const(
        Economy = iota
        Compact
        Standard
        FullSize
        Luxury
    )
    fmt.Println(Economy)
    fmt.Println(FullSize)


}
