package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Content string
}

func main() {
	messagePool := sync.Pool{
		New: func() interface{} {
			return new(Message)
		},
	}

	msg := messagePool.Get().(*Message)
	msg.Content = "hello"
	fmt.Println(msg.Content)

	msg.Content = ""
	messagePool.Put(msg)

	msg = messagePool.Get().(*Message)
	msg.Content = "world"
	fmt.Println(msg.Content)
}
