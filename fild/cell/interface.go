package cell

type ICell interface {
	Init(value int8, access bool) Cell
	GetCell() (value int8, access bool)
	SetCell(x, y, value int8) (access bool)
}
