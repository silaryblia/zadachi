package main

import "fmt"

func main() {
	scopedVar()
}

func scopedVar() {
	x := 5
	fmt.Println(x)
}
