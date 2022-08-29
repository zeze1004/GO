package main

import "fmt"

// 사각형 구조체
type Rectangle struct {
	width, height int
}

func main() {
	var rect1 Rectangle                   // 구조체 인스턴스를 생성
	var rect2 *Rectangle = new(Rectangle) // 구조체 포인터 선언 후 메모리 할당

	rect1.height = 20 // 구조체 인스턴스 필드에 접근할 때 . 사용
	rect2.height = 62 // 구조체 포인터에 접근할 때도 . 사용

	fmt.Println(rect1, rect2)
}
