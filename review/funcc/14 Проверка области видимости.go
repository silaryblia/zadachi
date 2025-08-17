package funcc

import "fmt"

var name int = 5

func proverka() {
	// видна (10)
	changeName()
}

func changeName() {

	fmt.Println(name + 5)
}
