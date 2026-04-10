package main

import "fmt"

type ID interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 // ~ 表示接受这个类型的别名类型
}

func GetInfo[T ID](id T) T {
	return id
}

func main() {
	fmt.Println(GetInfo[int64](123))
}
