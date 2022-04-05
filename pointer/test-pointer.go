package main

import "fmt"

func main() {
	var p *int
	num := 10
	p = &num
	fmt.Println(*p)

	*p = 20
	fmt.Println("num is :", num)
	fmt.Println(*p)
}
