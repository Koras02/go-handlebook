package main

func main() {

	// 변수 여러개 한번에 선언
	// 같은 타입인 변수 여러개를 콤마(,) 로 구분해 한꺼번에 선언
	var name, id, address string

	// 타입이 다른 변수 여러개를 한꺼번에 선언할 때 소괄호 (()) 로 묶어 표기

	var (
		name   string
		age    int
		weight float
	)
	// 변수 타입 생략
	// 변수 선언과 동시에 값을 할당할 때 타입 생략 가능
	var c = true
	// 타입을 생략할때 변수의 타입은 초기값의 타입으로 정해진다, 타입을 생략후 변수
	// 선언과 동시에 값을 할당할때 다음과 같이 특정 타입으로 변환된 값을 할당할 수 있다.

	var size = uint16(1024)

}
