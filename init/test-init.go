package main

import "fmt"

var ac = initVar()

func init() {
	fmt.Println("init function...")
}

func initVar() int {
	fmt.Println("initVar")
	return 100
}

func main() {
	fmt.Println("main")
}
