package main

import "fmt"

// 사각형 구조체
type Rectangle struct {
	width, height int
}

func main() {
	var rect1 Rectangle = Rectangle{10, 20} // 구조체 인스턴스를 생성하면서 값 초기화

	rect2 := Rectangle{45, 62} // var 키워드 생략 구조체 인스턴스를 생성하면서 값 초기화

	rect3 := Rectangle{width: 30} // 구조체 필드를 지정하여 값 초기화

	fmt.Println(rect1, rect2, rect3)
}
