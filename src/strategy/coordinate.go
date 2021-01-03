package strategy

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) GetX() int {
	return c.x
}

func (c *Coordinate) SetX(x int) {
	c.x = x
}

func (c *Coordinate) GetY() int {
	return c.y
}

func (c *Coordinate) SetY(y int) {
	c.y = y
}
