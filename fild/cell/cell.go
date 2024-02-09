package cell

type Cell struct {
	value  int8
	access bool
}

func (c Cell) Init(value int8, access bool) Cell {
	c.value = value
	c.access = access
	return c
}

func (c *Cell) GetCell() (value int8, access bool) {
	return c.value, c.access
}

func (c *Cell) SetCell(value int8, access bool) {
	c.value = value
	c.access = access
}
