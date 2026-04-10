package main

import "fmt"

// 1.接口->结构体：只看类型是否相同（因为结构体赋值给接口的时候已经检查过方法了）
// 2.接口->接口：只看方法是否相同

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
	Age  int
}

func (d Dog) Speak() string {
	return "Dog"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Cat"
}

func SpeakerTest(s Speaker) {
	if v, ok := s.(Cat); ok {
		fmt.Println(v.Speak())
	}
}

func main() {
	d := &Dog{"Dog", 30}
	SpeakerTest(d)
}
