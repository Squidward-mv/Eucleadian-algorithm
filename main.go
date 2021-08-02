package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter the number 1: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Enter the number 2: ")
	fmt.Scan(&b)

	var nod int

	for ;a != 0 && b != 0; {
		if a >= b {
			nod = b
			a = a % b
		}else {
			nod = a
			b = b % a
		}
	}

	fmt.Print(nod)
}