package main

/*
import "fmt"

var name int = 5

func gg() {
	changeName()
}

func changeName() {
	fmt.Println(name + 5)
}

/* ?????????????????????????????????
import "fmt"

func main() {
	double := makeMultiplier(2)
	fmt.Println(double(5)) // 10
}

// ??????????????????????????

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

/*
import "fmt"

func main() {
	add := func(a, b int) int { return a + b }
	result := add(3, 5)
	fmt.Println(result)
}

/*
import "fmt"

func main() {
	result1 := square(5) + square(4)
	result2 := square(6) + square(5)
	result3 := square(7) + square(6)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func square(n int) int {
	return n * n
}

/*
import "fmt"

func main() {
	scopedVar()
}

func scopedVar() {
	x := 5
	fmt.Println(x)
}

/*
import "fmt"

var counter int = 5

func main() {
	increment()
	increment()
	increment()

}

func increment() {
	fmt.Println(counter)
}

/*
import "fmt"

func main() {
	fmt.Println(max(5, 4))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
import "fmt"

func main() {
	lenght("hello")
}

func lenght(s string) {
	fmt.Println(len(s))
}

/*
import "fmt"

func main() {
	fmt.Println(isEven(4))
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

/*
import "fmt"

func main() {
	fmt.Println(swap(7, 9))
	// или
	a, b := swap(3, 5)
	fmt.Println(a, b)

}

func swap(a int, b int) (int, int) {
	return b, a
}

/*
import "fmt"

func main() {
	fmt.Println(getPi())
}

func getPi() float64 {
	return 3.14
}

/*
import "fmt"

func main() {
	greet("Stan")
	greet("John")
}

func greet(name string) {
	fmt.Printf("Привет, %s!\n", name)
}

/*
import "fmt"

func main() {
	result1 := sum(10, 5)
	result2 := sum(11, 6)
	result3 := sum(2, 4)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func sum(a int, b int) int {
	return a + b
}

/*
import "fmt"

func main() {
	var a, b, c string
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	b = string(a[len(a)-2])
	c = string(a[len(a)-1])
	fmt.Println("Две последние цифры числа:", b+c)
}

/*
import "fmt"

func main() {
	var a float64
	fmt.Println("Введите вашу зарплату:")
	fmt.Scan(&a)
	a = a / 100 * 87
	fmt.Println("Сумма после вычера 13% налога:", a)
}

/*
import "fmt"

func main() {
	var a, f float64
	fmt.Println("Введите температуру в Цельсиях:")
	fmt.Scan(&a)
	f = a*1.8 + 32
	fmt.Println("Температура в Фаренгейтах:", f)
}

/*
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	if a-b < 0 {
		fmt.Println("Разность чисел равна:", (a-b)*(-1))
	} else {
		fmt.Println("Разность чисел равна:", (a - b))

	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scan(&a)
	switch {
	case a%2 == 0:
		fmt.Println("0")
	default:
		fmt.Println(a % 2)
	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите ваш возраст")
	fmt.Scan(&a)
	if a < 7 {
		fmt.Println("Ты малыш")
	} else {
		fmt.Println("Ты взрослый")
	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scan(&a)
	if a%3 == 0 {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делится на 3")
	}
}

/*
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("Первое больше")
	} else if a < b {
		fmt.Println("Второе больше")
	} else {
		fmt.Println("Числа равны")
	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	if a>100{
	fmt.Println("Много")
	} else {
		fmt.Println("Норм")
	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	if a == 0 {
		fmt.Println("Число равно нулю")
	} else {
		fmt.Println("Число не равно нулю")
	}
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите сторону квадрата:")
	fmt.Scan(&a)
	fmt.Println("Площадь квадрата:", a*a)
}

/*
import "fmt"

func main() {
	var a int
	fmt.Println("Введите ваш возраст:")
	fmt.Scan(&a)
	fmt.Println("Примерное количество прожитых дней:", a*365)
}

/*
import "fmt"

func main() {
	var a, b, c float64
	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Среднее арифместическое чисел:", (a+b+c)/3)
}

/*
import "fmt"

func main() {
	var a, last string
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	if len(a) > 0 {
		last = string(a[len(a)-1])
	} else {
		fmt.Println("Вы не ввели число")
	}
	fmt.Println("Последняя цифра числа:", last)
}

/*
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа")
	fmt.Scan(&a, &b)
	fmt.Println("Остаток от деления:", a%b)
}

/*
import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите количество минут:")
	fmt.Scan(&a)
	b = a / 60
	c = a % 60
	fmt.Printf("Это %d часов и %d минут\n", b, c)
}

/*
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите цену товара:")
	fmt.Scan(&a)
	fmt.Println("Введите количество товаров")
	fmt.Scan(&b)
	fmt.Println("Общая стоимость:", a*b)
}

/*
import "fmt"

func main() {
	var a float64
	fmt.Println("Введите длину в сантиметрах")
	fmt.Scan(&a)
	fmt.Println("Длина в метрах: ", a/100)

}

/*
import "fmt"

func main() {
	var a, b float64
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	fmt.Println("Деление этих чисел равно: ", a/b)
}

/*
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	fmt.Println("Сложение этих чисел равно: ", a+b)
}

/*
import "fmt"

func main() {
	var n int
	fmt.Println("Введите число от 0 до 100")
	fmt.Scan(&n)
	switch {
	case n < 50:
		fmt.Println("Неудовлетворительно")
	case n >= 50 && n <= 69:
		fmt.Println("Удовлетворительно")
	case n >= 70 && n <= 89:
		fmt.Println("Хорошо")
	case n >= 90 && n <= 100:
		fmt.Println("Отлично")
	default:
		fmt.Println("Неверное число")
	}
}

/*
import "fmt"

func main() {
	var n int
	fmt.Println("Введите число")
	fmt.Scan(&n)
	switch {
	case n == 0:
		fmt.Println("ноль")
	case n > 0:
		fmt.Println("положительное")
	case n < 0:
		fmt.Println("отрицательное ")
	}
}

/*
import "fmt"

func main() {
	var n, m, max int
	fmt.Scan(&n)
	max = n
	fmt.Scan(&m)
	if m > n {
		max = m
	}
	fmt.Println(max)

}

/*
import "fmt"

func main() {
	var month int
	fmt.Println("Введите номер месяца")
	fmt.Scan(&month)
	switch {
	case month == 12 || month == 1 || month == 2:
		fmt.Println("Зима")
	case month >= 3 || month <= 5:
		fmt.Println("Весна")
	case month >= 6 || month <= 8:
		fmt.Println("Лето")
	case month >= 9 || month <= 11:
		fmt.Println("Осень")
	default:
		fmt.Println("Некорректный месяц")
	}
}

/*
import "fmt"

func main() {
	var n int
	fmt.Println("Введите ваш возраст")
	fmt.Scan(&n)
	if n >= 18 {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
}

/*
import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

}
*/
