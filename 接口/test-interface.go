package main

import "fmt"

var talker interface {
	talk() string
}

type duck struct {}

func (receiver duck) talk() string {
	return "ga ga ga...."
}

func main() {
	talker = duck{}
	talk:= talker.talk()
	fmt.Println(talk)
}