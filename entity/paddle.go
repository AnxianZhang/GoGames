package entity

import (
	"image/color"

	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var _ Entity = (*Paddle)(nil)

type Paddle struct {
	Object
}

func NewPaddle(x, y, width, height int) *Paddle {
	return &Paddle{*NewObject(x, y, width, height)}
}

func (p Paddle) Update(environment WorldView) gameStatus.GameStatus {
	return gameStatus.CONTINUE
}

func (p Paddle) Draw(screen *ebiten.Image) {
	vector.FillRect(screen,
		float32(p.GetX()), float32(p.GetY()),
		float32(p.getWidth()), float32(p.getHeight()),
		color.White, true,
	)
}

func (p Paddle) Tag() string {
	return "Paddle"
}
