package main 

import "fmt"

func main() {
	sum, i := 0, 0

	// for 문에는 조건식 생략 
	for {
		if i >= 10 {
			// 10보다 크거나 같을때가찌 
			break;
		}
		sum += i
		i++ 
	}
	fmt.Println(sum)
}