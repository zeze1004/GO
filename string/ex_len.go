package main

import "fmt"
import "unicode/utf8"

func main() {
	var s1 string = "한글"
	fmt.Println(utf8.RuneCountInString(s1))
}
