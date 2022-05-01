package main

import "fmt"

var talker interface {
	talk() string
}

type duck struct{}

func (receiver duck) talk() string {
	return "ga ga ga...."
}

type Computer struct {
	Name string
}

func (c *Computer) Read() string {
	return c.Name
}

func (c *Computer) Write(str string) {
	c.Name = str
}

func main() {
	talker = duck{}
	talk := talker.talk()
	fmt.Println(talk)

	var usb Usber
	usb = new(Computer)
	usb.Write("chuan")
	fmt.Println(usb.Read())
}
