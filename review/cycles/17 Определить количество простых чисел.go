package cycles

import "fmt"

// Определить количество простых чисел

// C клавиатуры вводятся целые числа до первого числа, которое меньше двух.
//  Написать программу, которая определяет сколько простых чисел было введено.

func natur() {
	count := 0
	fmt.Println("Введите числа: ")
	for {
		var number int
		fmt.Scan(&number)

		if number < 2 {
			break
		}

		isPrime := true
		if number == 2 {
			isPrime = true
		} else {
			for i := 2; i*i <= number; i++ {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
		}
		if isPrime {
			count++
		}
		fmt.Printf("Количество простых чисел: %d\n", count)
	}
}
