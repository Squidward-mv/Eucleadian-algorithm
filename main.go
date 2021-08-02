package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Enter the number 1: ")
	fmt.Scan(&num1)

	var num2 int
	fmt.Print("Enter the number 2: ")
	fmt.Scan(&num2)

	var nod int

	for ;num1 != 0 && num2 != 0; {
		if num1 >= num2 {
			nod = num2
			num1 = num1 % num2
		}else {
			nod = num1
			num2 = num2 % num1
		}
	}

	fmt.Print(nod)
}