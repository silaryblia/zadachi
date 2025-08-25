package player

import (
	"battleship/board"
	"math/rand"
	"time"
)

// Player представляет игрока
type Player struct {
	Name  string
	Board *board.Board
}

// NewPlayer создает нового игрока
func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Board: board.NewBoard(),
	}
}

// PlaceShipsRandomly размещает все корабли случайным образом
func (p *Player) PlaceShipsRandomly() {

	ships := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

	for _, size := range ships {
		p.placeShip(size)
	}
}

// placeShip размещает один корабль заданного размера
func (p *Player) placeShip(size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		isHorizontal := r.Intn(2) == 0

		var x, y int
		if isHorizontal {
			x = r.Intn(10)
			y = r.Intn(11 - size)
		} else {
			x = r.Intn(11 - size)
			y = r.Intn(10)
		}

		if p.Board.CanPlaceShip(x, y, size, isHorizontal) {
			p.Board.AddShip(x, y, size, isHorizontal)
			break
		}
	}
}

// Render отрисовывает поле игрока
func (p *Player) Render() {
	p.Board.Render()
}
