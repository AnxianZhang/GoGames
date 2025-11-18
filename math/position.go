package math

import (
	"github.com/AnxianZhang/GoGames/common"
	"math/rand"
)

var (
	UpDirection    = NewGridPosition(0, -1)
	DownDirection  = NewGridPosition(0, 1)
	LeftDirection  = NewGridPosition(-1, 0)
	RightDirection = NewGridPosition(1, 0)
	NoDirection    = NewGridPosition(0, 0)
)

type Position struct {
	x, y int
}

func NewPositionWithOffSet(_x int, _y int, xOffset int, yOffset int) Position {
	// Offsets will be added / removed the total grid boxes
	return Position{_x/common.GRID_SIZE/2 + xOffset, _y/common.GRID_SIZE/2 + yOffset}
}

func NewGridPosition(_x int, _y int) Position {
	return Position{_x, _y}
}

func RandomPosition() Position {
	return NewGridPosition(rand.Intn(common.X_CASE), rand.Intn(common.Y_CASE))
}

func (p Position) GetX() int {
	return p.x
}

func (p Position) GetY() int {
	return p.y
}

func (p Position) IsEqualTo(other Position) bool {
	return p.x == other.GetX() && p.y == other.GetY()
}

func (p *Position) Add(other Position) Position {
	p.x += other.x
	p.y += other.y

	return *p
}

func (p *Position) IsInCollisionWithScreen(entityPosition []Position) bool {
	// is in collision with the borders
	if p.y < 0 || p.y >= common.Y_CASE || p.x < 0 || p.x >= common.X_CASE {
		return true
	}

	// check if it is a self collision
	for _, e := range entityPosition {
		if p.x == e.GetX() && p.y == e.y {
			return true
		}
	}

	return false
}
