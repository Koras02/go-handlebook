// 주석
// go 언어는 다음과 같이 주석을 사용한다.

// 다음에 나타나는 문자에는 모두 주석으로 인식한다.

package main

import "fmt"

func main() {
	fmt.Println("Hello World") // Hello World 출력

	// 줄 내부에서 일부부만 주석으로 만들거나 여러 줄을 주석으로 만들때에는 /* */를 사용

	fmt.Println( /* 인사말 */ "Hello world two")
}
