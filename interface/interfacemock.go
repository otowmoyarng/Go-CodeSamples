package main

import "fmt"

type OutputMock struct {
	Message string
}

func (o OutputMock) Do() {
	fmt.Println("(Mock)OutputMessage Do()")
}

func (o OutputMock) Output() {
	fmt.Println("(Mock)" + o.Message)
}
