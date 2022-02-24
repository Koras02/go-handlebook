package main

import "fmt"

func main() {
	c := "A"
	switch {
	// case에 조건식 사용
	case "0" <= c && c <= "9":
		// 숫자를 판단하는 case
		fmt.Printf("%c은(는) 숫자입니다", c)
	case "a" <= c && c <= "z":
		// 소문자 a에서 z인지를 판단
		fmt.Printf("%c은(는) 소문자입니다", c)
	case "A" <= c && c <= "Z":
		// 대문자 a에서 Z인지를 판다
		fmt.Printf("%c은(는) 대문자입니다", c)

	}
}
