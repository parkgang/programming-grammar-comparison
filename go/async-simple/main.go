package main

import (
	"fmt"
	"time"
)

func WorkDo(channel chan string, sec time.Duration, content string) {
	time.Sleep(sec * time.Millisecond)
	channel <- content
}

func main() {
	message := make(chan string)

	fmt.Println("go 비동기 작업 시작")

	go WorkDo(message, 200, "1번째 비동기 작업")
	go WorkDo(message, 1000, "2번째 비동기 작업")
	go WorkDo(message, 10, "3번째 비동기 작업")
	go WorkDo(message, 500, "4번째 비동기 작업")
	go WorkDo(message, 600, "5번째 비동기 작업")

	fmt.Printf("\t%s\n", <-message)
	fmt.Printf("\t%s\n", <-message)
	fmt.Printf("\t%s\n", <-message)
	fmt.Printf("\t%s\n", <-message)
	fmt.Printf("\t%s\n", <-message)

	fmt.Println("비동기 작업 완료! 프로그램을 종료합니다.")
}
