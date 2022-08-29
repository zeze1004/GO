package main

import "fmt"

// 사각형 구조체
type Rectangle struct {
	width, height int
}

func main() {
	var rect = Rectangle // 구조체 인스턴스 선언

	var rect1 = *Rectangle // 구조체 포인터 선언
	rect1 = new(Rectangle)

	rect2 := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당

	fmt.Println(rect, rect1, rect2)
}
