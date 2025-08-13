package block1

import "fmt"

func main() {
	countdown(5)
}

func countdown(n int) {
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
}
