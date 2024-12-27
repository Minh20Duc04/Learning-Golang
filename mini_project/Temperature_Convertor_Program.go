package main

import (
	"fmt"
	"strings"
)

func main() {
	//nhập vào nhiệt độ muốn đổi và chuyển đổi giữa các độ F, K, C
	var temperature float64
	fmt.Println("Enter temperature: ")
	fmt.Scanln(&temperature)

	var unit, targetUnit string
	fmt.Println("Enter your current unit: ")
	fmt.Scanln(&unit)
	unit = strings.ToUpper(unit)

	fmt.Println("Enter your target unit: ")
	fmt.Scanln(&targetUnit)
	targetUnit = strings.ToUpper(targetUnit)


	switch unit {
	case "C": 
	if targetUnit == "K"{
		temperature += 273.15
	}else if targetUnit == "F"{
		temperature = temperature * 9/5 + 32
	}else{
		fmt.Println("This targetUnit can't not convert: ", targetUnit)
	}
	case "F" :
		if targetUnit == "C"{
			temperature = (temperature - 32) * 5/9
		}else if targetUnit == "K"{
			temperature = (temperature - 32) * 5/9 +273.15
		}else{
		fmt.Println("This targetUnit can't not convert: ", targetUnit)
	}
	case "K":
		if targetUnit == "C"{
			temperature -= 273.15
		}else if targetUnit == "F"{
			temperature = (temperature - 273.15) * 9/5 + 32
		}else{
			fmt.Println("This targetUnit can't not convert: ", targetUnit)
		}
	}
	
	fmt.Println("After converting, your current temperature is: ", temperature)

}