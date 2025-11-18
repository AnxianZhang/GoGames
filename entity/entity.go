package entity

import (
	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/math"
	"github.com/hajimehoshi/ebiten/v2"
)

// An Entity is a thing that might be alive in our game
// with some action do to
type Entity interface {
	Update(environment WorldView) gameStatus.GameStatus
	Draw(screen *ebiten.Image)

	// Tag is used to identify the identity type
	Tag() string
}

// Object is alongside with the entity, both need to be implemented / inherited
type Object struct {
	math.Position
	width  int
	height int
}

func NewObject(_x, _y, _w, _h int) *Object {
	return &Object{
		math.NewGridPosition(_x, _y),
		_w, _h,
	}
}

func (o Object) getWidth() int {
	return o.width
}

func (o Object) getHeight() int {
	return o.height
}
