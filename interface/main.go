package main

import "fmt"

func main() {
	o := Output{"output"}
	fmt.Printf("Output.Output():")
	o.Output()
	fmt.Printf("Output.Do():")
	o.Do()

	p := OutputMock{"outputmock"}
	fmt.Printf("OutputMock.Output():")
	p.Output()
	fmt.Printf("OutputMock.Do():")
	p.Do()
}
