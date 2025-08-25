package board

import (
	"fmt"
)

// Board представляет игровое поле
type Board struct {
	Cells [10][10]Cell
	Ships []ShipInfo
}

// ShipInfo информация о корабле
type ShipInfo struct {
	Size  int
	Cells []*Cell
}

// NewBoard создает новое поле
func NewBoard() *Board {
	b := &Board{}

	// Инициализируем клетки
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Cells[i][j] = Cell{X: i, Y: j, HasShip: false}
		}
	}
	return b
}

// Render отрисовывает поле игрока
func (b *Board) Render() {
	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 10; j++ {
			cell := b.Cells[i][j]
			if cell.HasShip {
				fmt.Print("■ ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// AddShip добавляет корабль на поле
func (b *Board) AddShip(x, y, size int, isHorizontal bool) {
	ship := ShipInfo{Size: size}
	for i := 0; i < size; i++ {
		var cell *Cell
		if isHorizontal {
			cell = &b.Cells[x][y+i]
		} else {
			cell = &b.Cells[x+i][y]
		}
		cell.HasShip = true
		ship.Cells = append(ship.Cells, cell)
	}
	b.Ships = append(b.Ships, ship)
}

// CanPlaceShip проверяет, можно ли разместить корабль
func (b *Board) CanPlaceShip(x, y, size int, isHorizontal bool) bool {
	for i := -1; i <= size; i++ {
		for j := -1; j <= 1; j++ {
			var checkX, checkY int

			if isHorizontal {
				checkX = x + j
				checkY = y + i
			} else {
				checkX = x + i
				checkY = y + j
			}

			if checkX < 0 || checkX >= 10 || checkY < 0 || checkY >= 10 {
				continue
			}

			if b.Cells[checkX][checkY].HasShip {
				return false
			}
		}
	}
	return true
}

// IsValidPlacement проверяет корректность расположения всех кораблей
func (b *Board) IsValidPlacement() bool {
	return b.GetShipsCount() == 10
}

// GetShipsCount возвращает количество кораблей
func (b *Board) GetShipsCount() int {
	return len(b.Ships)
}

// GetShipCells возвращает количество клеток кораблей
func (b *Board) GetShipCells() int {
	count := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if b.Cells[i][j].HasShip {
				count++
			}
		}
	}
	return count
}
