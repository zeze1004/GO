package main

import "fmt"

func main() {
	var numPtr *int = new(int)

	numPtr++               // 컴파일 에러, 포인터 연산 허용x
	numPtr = 0x1400011c008 // 컴파일 에러, 메모리 주소를 직접 대입 x

	fmt.Println(numPtr)
}
