package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func(c chan int) {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}


	testSelect()

}

func testSelect() {
	c1 := make(chan int, 0)
	c2 := make(chan int, 0)
	go produce1(c1)
	go produce2(c2)
	go consume(c1, c2)
	time.Sleep(20 * time.Second)
}

func produce1(ch chan int)  {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func produce2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2 + 1
	}
}

func consume(c1, c2 chan int) {
	for {
		select {
		case v := <- c1:
			fmt.Printf("data from chanel 1 is %d\n", v)
		case v := <- c2:
			fmt.Printf("data from chanel 2 is %d\n", v)
		}
	}
}