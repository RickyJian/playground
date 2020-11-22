package main

import (
	"fmt"
	"time"
)

func main() {
	// done is the signal for work when it cancel or done
	done := make(chan interface{})
	terminated := work(done, generate(done, "1", "2", "3", "4", "5", "6", "7", "8"))

	// cancel work
	go func() {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("cancel work")
		close(done)
	}()

	// block main goroutine here
	<-terminated
	fmt.Println("shutdown")
	// results:
	// 			1
	// 			cancel work
	// 			work done
	// 			shutdown
}

func work(done <-chan interface{}, texts <-chan string) <-chan interface{} {
	terminated := make(chan interface{})

	go func() {
		defer fmt.Println("work done")
		defer close(terminated)
		for {
			select {
			case text := <-texts:
				fmt.Println(text)
				time.Sleep(1 * time.Second)
			case <-done:
				return
			}
		}
	}()

	return terminated
}

func generate(done <-chan interface{}, texts ...string) <-chan string {
	textChan := make(chan string)

	go func() {
		defer close(textChan)
		for _, text := range texts {
			select {
			case textChan <- text:
			case <-done:
				return
			}
		}
	}()

	return textChan
}
