package entity

import (
	"image/color"

	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var _ generic.Entity = (*Paddle)(nil)

type Paddle struct {
	generic.Object
}

func NewPaddle(x, y, width, height int) *Paddle {
	return &Paddle{*generic.NewObject(x, y, width, height)}
}

func (p Paddle) Update(environment generic.WorldView) gameStatus.GameStatus {
	return gameStatus.CONTINUE
}

func (p Paddle) Draw(screen *ebiten.Image) {
	vector.FillRect(screen,
		float32(p.GetX()), float32(p.GetY()),
		float32(p.GetWidth()), float32(p.GetHeight()),
		color.White, true,
	)
}

func (p Paddle) Tag() string {
	return "Paddle"
}
