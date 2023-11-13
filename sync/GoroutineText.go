package main

import (
	"fmt"
	"time"
)

// 从 goroutine
func newText() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine i : %d\n", i)
		time.Sleep(1 * time.Second)

	}

}

// 主 goroutine
func main() {
	// 创建 go 程 去执行 newText() 方法
	go newText()
	i := 0
	for {
		i++
		fmt.Printf("main grroutine i: %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
