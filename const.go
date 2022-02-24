package main

import "fmt"

// 상수
// 처음 선언한 이후 변하지 않는 변수는 상수(constant)로 선언
// 변수는 var 키워드로 선언했듯 상수는 const 키워드로 선언, 상수는 불(boolean), 숫자, 문자열 타입으로만 선언 가능

func main() {

	// const limit = 64
	// 상수를 선언할 시 타입은 표기하지 않아도 된다. 할당 되는 값의 타입에 따라서 컴파일러가
	// 자동적으로 상수의 타입을 결정 명시적인 타입을 지정해주려면 다음과 같이 타입을 표기해주자.
	// const max uint64 = 1024

	// 상수는 컴파일할 때 ㅏ값이 정해진다. 특정 계산식의 결과를 상수로 지정할 수 있는데, 이런 계산식은
	// 커파일 시 연산할 수 있어야 함

	// const max = 1024 * 1024 // 유효함
	// // const c = getNumber() // 유효하지 않음

	// fmt.Println(max)

	// 변수 선언과 마찬가지로 상수 여러개를 한꺼번에 선언시 ()로 묶어 표기
	const (
		RED    = 0
		ORANGE = 1
		YELLO  = 2
	)

	fmt.Println(RED)

	// RED = 2
}
