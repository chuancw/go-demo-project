package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go sum1(c)
	// 读通道，如果没有数据，则会阻塞
	<-c
	fmt.Println("main end...")
}

func sum1(c chan struct{}) {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	time.Sleep(5 * time.Second)
	// 写通道
	c <- struct{}{}
}
