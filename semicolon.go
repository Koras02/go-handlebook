// 세미콜론 생략 가능 
// go에서는 세미콜론을(;)으로 문장의 끝을 표시,

// go 컴파일러는 세미콜론을 기준으로 문장 단위를 인식,
// 그래서 문장 여러 개를 세미콜론으로 구분해 한 줄에 작성 가능 


// 세미콜론 자동 삽입으로 인한 컴파일 오류의 예
package main

import "fmt"

/* 
  실행 결과
  prog.go:10: syntax error: missing { after for clause prog.go:14: syntax error: unexpected }
*/

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}

	for j := 5; j >= 0; j-- // for문에 세미콜론이 삽입되 컴파일 오류 발생 
	{ 
        fmt.Print(j)
	}
}