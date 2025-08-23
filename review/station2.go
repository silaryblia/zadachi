package main

import "fmt"

// Авто
type Car struct {
	Name     string
	Fuel     float64
	MaxFuel  float64
	FuelType string
}

// Станция
type GasStation struct {
	Name        string
	Fuel95      float64
	Fuel98      float64
	FuelDiesel  float64
	Price95     float64
	Price98     float64
	PriceDiesel float64
}

// Метод заправки
func (gs *GasStation) Refuel(car *Car, amount float64, money float64) bool {
	fmt.Printf("Попытка заправки %s на %s\n", car.Name, gs.Name)

	// Проверка совместимости топлива
	if car.FuelType != "95" && car.FuelType != "98" && car.FuelType != "diesel" {
		fmt.Println("Ошибка: неверный тип топлива")
		return false
	}

	// Проверяем, есть ли нужное топливо на заправке
	var availableFuel *float64
	var price float64

	switch car.FuelType {
	case "95":
		availableFuel = &gs.Fuel95
		price = gs.Price95
	case "98":
		availableFuel = &gs.Fuel98
		price = gs.Price98
	case "diesel":
		availableFuel = &gs.FuelDiesel
		price = gs.PriceDiesel
	}

	if *availableFuel <= 0 {
		fmt.Println("На заправке закончилось топливо этого типа")
		return false
	}

	// Проверяем достаточно ли денег
	cost := amount * price
	if money < cost {
		fmt.Printf("Недостаточно денег. Нужно: %.2f, есть %.2f\n", cost, money)
		return false
	}

	// Проверяем, полный ли бак
	if car.Fuel >= car.MaxFuel {
		fmt.Println("Бак полный!")
		return false
	}

	// Рассчитываем сколько можно заправить
	needed := car.MaxFuel - car.Fuel
	if amount > needed {
		amount = needed
		cost = amount * price
		fmt.Printf("Заправляем только %.2f л (до полного бака)\n", amount)
	}

	// Проверяем хватит ли топлива на заправке
	if amount > *availableFuel {
		amount = *availableFuel
		cost = amount * price
		fmt.Printf("Заправляем только %.2f л (ограничение заправки)\n", amount)
	}

	// Производим заправку
	car.Fuel += amount
	*availableFuel -= amount
	change := money - cost

	fmt.Printf("Успешно заправлено: %.2f л %s\n", amount, car.FuelType)
	fmt.Printf("Стоимость: %.2f руб.\n", cost)
	fmt.Printf("Сдача: %.2f руб.\n", change)
	fmt.Printf("В баке теперь: %.2f/%.2f л\n\n", car.Fuel, car.MaxFuel)

	return true
}

func main() {
	// add auto
	myCar := Car{
		Name:     "BMW 3",
		Fuel:     20.0,
		MaxFuel:  60.0,
		FuelType: "95",
	}

	// create station
	station := GasStation{
		Name:        "Shell",
		Fuel95:      1000.0,
		Fuel98:      800.0,
		FuelDiesel:  1200.0,
		Price95:     50.0,
		Price98:     55.0,
		PriceDiesel: 48.0,
	}

	// показываем начальное состояние
	fmt.Printf("Машина: %s\n", myCar.Name)
	fmt.Printf("Топливо: %.2f/%.2f л\n", myCar.Fuel, myCar.MaxFuel)
	fmt.Printf("Заправка %s: 95-%.2fл, 98-%.2fл, дизель-%.2fл\n\n",
		station.Name, station.Fuel95, station.Fuel98, station.FuelDiesel)

	// пробуем заправиться
	station.Refuel(&myCar, 50.0, 3000.0) // заправка на 50л с 3000 руб

	// пробуем заправиться еще раз (почти полный бак)
	station.Refuel(&myCar, 50.0, 3000.0)

	// пробуем заправиться неправильным топливом
	wrongCar := Car{Name: "Lada", Fuel: 10, MaxFuel: 40, FuelType: "92"}
	station.Refuel(&wrongCar, 20.0, 1000.0)
}
