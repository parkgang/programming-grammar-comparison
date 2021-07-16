package main

import (
	"fmt"
	"time"
)

func setTimeout() string {
	time.Sleep(1 * time.Millisecond)
	return "Async CallBack!"
}

func main() {
	fmt.Println("비동기 동기화 예제 시작")

	ch := make(chan string)

	go func() {
		ch <- setTimeout()
	}()

Loop:
	for i := 0; i < 20; i++ {
		select {
		case v := <-ch:
			fmt.Printf("%s\n", v)
			// 아래의 break 주석을 해제하면 콜백되면 반복문을 바로 종료헙나다.
			break Loop
		default:
			fmt.Printf("for loop: %d\n", i)
		}
	}

	fmt.Println("프로그램을 종료합니다")
}
