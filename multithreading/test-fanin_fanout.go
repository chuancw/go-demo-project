package main

import (
	"fmt"
	"math/rand"
)

/*
扇入和扇出
扇入：将多路通道聚合到一条通道中处理
扇出：将一条通道发散到多条通道中处理
*/

func main() {
	//intChan := gen(1, 2, 3)
	//seqResult := seq(intChan)
	//fmt.Println(<- seqResult)
	//fmt.Println(<- seqResult)
	//fmt.Println(<- seqResult)

	//
	//
	//fmt.Println(<- seqResult)

	//ch := genIntA()
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	////fmt.Println(<-ch)
	//fmt.Println(len(ch))

	ch := GenInt()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}

func gen(numbers ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, number := range numbers {
			out <- number
		}
		fmt.Println("gen goroutine end")
	}()
	fmt.Println("gen end")
	return out
}

func seq(input chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range input {
			out <- i * i
		}
		fmt.Println("seq goroutine end")
	}()
	fmt.Println("seq end")
	return out
}

func GenInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for true {
			select {
			case ch <- <-genIntA():
			case ch <- <-genIntB():

			}
		}
	}()
	return ch
}

func genIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func genIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for true {
			ch <- rand.Int()
		}
	}()
	return ch
}
