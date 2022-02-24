package main 

import "fmt"

func main() {
	// 초기값 설정 
	sum, i := 0, 0
    // for문에 조건식만을 사용한다.
	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum)
}