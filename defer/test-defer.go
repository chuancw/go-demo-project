package main

import "fmt"

func main() {
	defer fmt.Println("step2")
	defer fmt.Println("step1")
	fmt.Println("end")
}
