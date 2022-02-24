package main

import "fmt"

const (
	// iota로 비트 연산의 결과값을 상수로 선언
	Running = 1 << iota // 1 << 0 == 1
	Wating              // 1 << 1 == 2
	Send                // 1 << 2 == 4
	Receive             // 1 << 3 == 8
)

func main() {

	// OR 연산자로(|)로 stat 변수 생성
	stat := Running | Send
	display(stat)
}

func display(stat int) {
	// AND 연산자(&)로 stat에 포함된 비트 값의 상태 출력
	// 이 예제에서는 상수 4개 (Running, Wating, Send, Receive)를 그룹으로 선언
	// 상수 그룹에는 비트 연산의 결과가 할당되게 하고, 각 비트 연산의 피연산자로 iota를

	// 사용해서 0부터 1씩 증가하게 했다. main 함수에는 OR연산자 (|)로 stat변수를 생성하고
	// display 함수에서는 AND 연산자 (&)로 stat에 포함된 비트 값의 상태를 출력했다.

	if stat&Running == Running {
		fmt.Println("Running")
	}

	if stat&Wating == Wating {
		fmt.Println("Wating")
	}

	if stat&Send == Send {
		fmt.Println("Send")
	}

	if stat&Receive == Receive {
		fmt.Println("Receive")
	}

}
