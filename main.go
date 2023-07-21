package main

import (
	"fmt"
)

func main() {
	first, second, action := interaction()
	switch action {
	case 1:
		result := add(first, second)
		fmt.Printf("%v / %v = %v", first, second, result)
	case 2:
		result := subtract(first, second)
		fmt.Printf("%v / %v = %v", first, second, result)
	case 3:
		result := multiply(first, second)
		fmt.Printf("%v / %v = %v", first, second, result)
	case 4:
		if second != 0 {
			result := divide(first, second)
			fmt.Printf("%v / %v = %v", first, second, result)
		} else {
			fmt.Println("Can not divide by zero")
		}
	}
}

func interaction() (first, second, action int) {
	fmt.Println("enter first number:")
	_, err := fmt.Scan(&first)
	if err != nil {
		panic(err)
	}
	fmt.Println("enter second number:")
	_, err = fmt.Scan(&second)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------")
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide")
	fmt.Println("-----------")
	_, err = fmt.Scan(&action)
	if err != nil {
		panic(err)
	}
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
