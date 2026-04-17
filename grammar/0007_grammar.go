package main

import (
	"fmt"
	"time"
)

// 总结：
// 1.在新版本 1.22+，for 以及 range 都会在被闭包捕获计数器的时候，每次迭代都创建新变量
// 2.range 里的 v 是副本，相当于在每轮初始化里进行了一次赋值

func main() {
	strs := []string{"one", "two", "three"}
	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(s)
		}()
	}
	time.Sleep(3 * time.Second)
}
