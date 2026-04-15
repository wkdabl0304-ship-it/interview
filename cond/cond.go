package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	readyCount := 0

	tasks := []string{"DB", "Cache", "MQ"}
	for _, task := range tasks {
		go func(t string) {
			time.Sleep(time.Duration(1+time.Duration(time.Now().UnixNano()%3)) * time.Second)
			mu.Lock()
			readyCount++
			fmt.Printf("%s is ready\n", t)
			cond.Broadcast()
			mu.Unlock()
		}(task)
	}

	mu.Lock()
	for readyCount < len(tasks) {
		fmt.Printf("Only %d tasks are ready...\n", readyCount)
		cond.Wait()
	}
	mu.Unlock()

	fmt.Println("All tasks done")
}
