package main

import "fmt"

// gofmt로 코드 서식 지정
// 괄호나 들여쓰기 같은 코드의 지정 문제로 개발자 커뮤니티에서 자주 등장하는
// 논쟁거리다. Go의 코드의 스타일을 자동으로 맞춰주는 gofmt 도구를 제공함으로써 이러한 논쟁을 해결함
// gofmt 도구를 사용하면 코드의 스타일을 Go에서 사용하는 스타일 대로 맞춰준다.

func main() {

	type Rect struct {
		width  int // width
		height int // height
	}

	// 스타일에 맞지 않는 코드
	// r := Rect{1, 2}; fmt.Println(r.width * 2 + r.height * 2)

	r := Rect{1, 2}
	// 1 * 2 = 2 + 2 * 2 = 4 = 6
	fmt.Println(r.width*2 + r.height*2)
}
