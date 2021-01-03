package models

import "rover/src/strategy"

type Plateau struct {
	leftBottom *strategy.Coordinate
	rightTop   *strategy.Coordinate
}

func (p *Plateau) GetLeftBottom() *strategy.Coordinate {
	return p.leftBottom
}

func (p *Plateau) SetLeftBottom(leftBottom *strategy.Coordinate) {
	p.leftBottom = leftBottom
}

func (p *Plateau) GetRightTop() *strategy.Coordinate {
	return p.rightTop
}

func (p *Plateau) SetRightTop(rightTop *strategy.Coordinate) {
	p.rightTop = rightTop
}
