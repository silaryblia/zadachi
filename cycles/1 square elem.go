package cycles

import "fmt"

// Вывод квадратов натуральных чисел
func square() {
	n := 50
	for i := 1; i < n; i++ {
		if i*i < n {
			fmt.Print(i*i, " ")
		}
	}
}
