package main

import "fmt"

func main() {
	str := "hello안녕"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(runes) = %d\n", len(runes))
}
