package ship

// CellRef представляет ссылку на клетку (координаты)
type CellRef struct {
	X, Y int
}

// Ship представляет корабль
type Ship struct {
	Size  int
	Cells []*CellRef
	Sunk  bool
}
