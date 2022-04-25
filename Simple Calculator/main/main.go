package main

//fmt : Print
import (
	"fmt"
)

func main() {
	var a, b int
	var center string
	fmt.Print("입력 : ")
	fmt.Scanln(&a, &center, &b)
	fmt.Print("결과값: ")
	if center == "*" {
		fmt.Println(a * b)
	} else if center == "/" {
		fmt.Println(a / b)
	} else if center == "-" {
		fmt.Println(a - b)
	} else if center == "+" {
		fmt.Println(a + b)
	} else {
		fmt.Print("연산자 오류 : ")
		fmt.Println(center)
	}
}
