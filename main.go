package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func div2(a int, b int) int {
	return a / b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 22))
	fmt.Println(mul(1, 2))
	fmt.Println(div(1, 2))
	fmt.Println(div2(1, 2))
}
