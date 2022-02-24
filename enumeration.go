// 열거형
// 열거형(enumeration)은 차례로 1씩 증가하는 상수의 묶음
package main

import "fmt"

// Go에서는 다른 언어에서 기본으로 사용하는 개념을 사용하지 않는 경우가 많다, 열거형도 그중 하나

func main() {

	// const (
	// 	Sunday   = 0
	// 	Monday   = 1
	// 	Tuesday  = 2
	// 	Thursday = 3
	// 	Friday   = 4
	// 	Saturday = 5
	// )

	// fmt.Println(Sunday)

	// 상수를 열겨형으로 선언할 때는 iota 예약어를 사용하면 편리하다. 상수를 그룹으로
	// 묶어 선언할때에는 const 그룹에서 iota의 값은 0이고, 이후로 1씩 증가한다, iota로 위 코드를
	// 다음과 같이 정의
	// iota로 1씩 증가하는 값은 새로운 const 그룹을 만나면 0으로 초기화
	const (
		Sunday   = iota // 0
		Monday          // 1
		Tuesday         // 2
		Thursday        // 3
		Friday          // 4
		Saturday        // 5
	)

	type Color int
	const (
		RGB    Color = iota // 0
		ORANGE              // 1
		YELLOW              // 2
		GREEN               // 3
	)
	fmt.Println(ORANGE)
}
