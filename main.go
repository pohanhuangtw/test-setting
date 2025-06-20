package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
}
