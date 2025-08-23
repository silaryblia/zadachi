package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
}

type Storage struct {
	Products map[int]Product
}

type Warehouse interface {
	AddProduct(product Product) error
	UpdateQuantity(productID int, quantity int) error
}

// метод добавления нового товара на склад
func (st *Storage) AddProduct(product Product) error {
	if st.Products == nil {
		st.Products = make(map[int]Product)
	}

	if _, ok := st.Products[product.ID]; ok {
		return fmt.Errorf("товар с ID %d уже существует на складе", product.ID)
	}

	st.Products[product.ID] = product
	return nil
}

// метод изменения количества товара на складе
func (st *Storage) UpdateQuantity(productID int, quantity int) error {
	product, ok := st.Products[productID]
	if !ok {
		return fmt.Errorf("товар с ID %d не найден на складе", productID)
	}

	if product.Quantity+quantity < 0 {
		return fmt.Errorf("количество товара не может быть отрицательным")
	}

	product.Quantity += quantity
	st.Products[productID] = product
	return nil
}

/////////////
func main() {
	// создание содержимого склада
	storage := &Storage{
		Products: make(map[int]Product),
	}

	// добавление товаров на склад
	err := storage.AddProduct(Product{ID: 1, Name: "Ноутбук", Quantity: 10})
	if err != nil {
		fmt.Print("Ошибка при добавлении товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 10}

	// изменение количества товаров на складе
	err = storage.UpdateQuantity(1, -5) // продали 5 ноутбуков
	if err != nil {
		fmt.Print("Ошибка при изменении количества товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 5}
}
