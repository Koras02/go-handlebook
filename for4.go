// 모호한 요소 제거 
// ++ --같은 증감 연산자에는 후차연산으로만 사용할 수 있고, 증감 연산은 반환값이 없다.
// 즉 i = i++ 또는 ++i 같은 코드는 허용되지 않는다. 

// 1. 증감연산은 반환 값이 없다.

/* 
 실행 결과
 prog.go:8: syntax error: unexpected ++, expecting semicolon or newline or }
*/

// package main

// import "fmt"

// func main() {
//     sum, i := 0, 0
//     for i < 10 {
// 		sum += i++ // 증감 연산은 반환 값이 없음. 컴파일 에러 발생
// 	}
// 	fmt.Println(sum)
// }


//  2. 증감 연산자는 전치 연산은 허용하지 않는다.
package main 

import "fmt"


/** 
prog.go:9: syntax error: unexpected ++
**/
func main() {
	sum, i := 0,0 
	for i < 10 {
		sum += i
		++i // 증감 연산자에는 전치 연산은 허용하지 않는다. 컴파일 에러 발생
	}
	fmt.Println(sum)
}