package main

type Usber interface {
	Read() string
	Write(str string)
}
