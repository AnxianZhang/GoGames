package entity

import (
	"image"
	"log"

	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ generic.Entity = (*Enemy)(nil)

type Enemy struct {
	*generic.Sprite
	canAttack bool
	camera    *Camera
}

func NewEnemy(x, y int, img *ebiten.Image, canAttack bool, camera *Camera) *Enemy {
	return &Enemy{generic.NewSprite(img, x, y), canAttack, camera}
}

func (e *Enemy) Update(environment generic.WorldView) gameStatus.GameStatus {
	rawPlayer, ok := environment.FindFirstEntity("Player")

	if !ok {
		log.Fatal("No player entity was found in enemy update")
	}

	player := rawPlayer.(*Player)

	if e.canAttack {
		if e.GetX() < player.GetX() {
			e.MoveRight(1)
		} else if e.GetX() > player.GetX() {
			e.MoveLeft(1)
		}

		if e.GetY() < player.GetY() {
			e.MoveDown(1)
		} else if e.GetY() > player.GetY() {
			e.MoveUp(1)
		}
	}

	return gameStatus.CONTINUE
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(e.GetX()), float64(e.GetY()))
	options.GeoM.Translate(float64(e.camera.GetX()), float64(e.camera.GetY()))

	screen.DrawImage(e.GetImage().SubImage(
		// 16 and 16 is the dimension of one tiles
		image.Rect(0, 0, 16, 16),
	).(*ebiten.Image), &options)
}

func (e *Enemy) Tag() string {
	return "Enemy"
}
