package main

import "fmt"

type OutputMessage interface {
	Do()
}

type Output struct {
	Message string
}

func (o Output) Do() {
	fmt.Println("OutputMessage Do()")
}

func (o Output) Output() {
	fmt.Println(o.Message)
}
