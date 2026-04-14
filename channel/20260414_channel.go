package main

import (
	"fmt"
	"time"
)

// 1.channel 关闭以及缓冲区为空，v, ok := <-ch 中的 ok 才为 false
// 2.case 里的语句是否运行取决于是否阻塞，与通道是否开关没有关系

func worker(id int, ch <-chan int, done chan<- struct{}) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("worker %d is closed\n", id)
				done <- struct{}{}
				return
			}
			fmt.Printf("worker %d got %d\n", id, v)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	done := make(chan struct{})
	for i := 0; i <= 3; i++ {
		go worker(i, ch, done)
	}
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
	fmt.Println("main: channel closed")
	for i := 1; i <= 3; i++ {
		<-done
	}
	fmt.Println("main: all workers have finished")
}
