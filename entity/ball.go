package entity

import (
	"image/color"

	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/AnxianZhang/GoGames/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var _ generic.Entity = (*Ball)(nil)

type Ball struct {
	generic.Object
	velocity *math.Velocity // per tick
}

func NewBall(x, y, xVelocity, yVelocity, width, height int) *Ball {
	return &Ball{*generic.NewObject(x, y, width, height),
		math.NewVelocity(xVelocity, yVelocity),
	}
}

func (b *Ball) Update(environment generic.WorldView) gameStatus.GameStatus {
	switch {
	case b.Object.GetX() >= common.SCREEN_WIDTH:
		b.ResetPosition()
		return gameStatus.LOSE
	case b.Object.GetX() <= 0:
		b.velocity.SetX(common.BALL_SPEED)
	case b.Object.GetY() >= common.SCREEN_HEIGHT:
		b.velocity.SetY(common.BALL_SPEED)
	case b.Object.GetY() <= 0:
		b.velocity.SetY(-common.BALL_SPEED)
	}

	b.Object.MoveUp(b.velocity.GetY())
	b.Object.MoveRight(b.velocity.GetX())

	rawPaddle, _ := environment.FindFirstEntity("Paddle")
	paddle := rawPaddle.(*Paddle)

	if b.Object.GetX() >= paddle.GetX() &&
		b.Object.GetY() >= paddle.GetY() && b.Object.GetY() <= paddle.GetY()+paddle.GetHeight() {
		b.velocity.SetX(-common.BALL_SPEED)
		return gameStatus.GET_POINT
	}

	return gameStatus.CONTINUE
}

func (b Ball) Draw(screen *ebiten.Image) {
	vector.FillRect(screen,
		float32(b.GetX()), float32(b.GetY()),
		float32(b.GetWidth()), float32(b.GetHeight()),
		color.White, true,
	)
}

func (b Ball) Tag() string {
	return "Ball"
}

func (b *Ball) ResetPosition() {
	b.Object.Position = math.NewGridPosition(0, 0)
}
