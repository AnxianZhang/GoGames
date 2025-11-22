package entity

import (
	"log"

	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ generic.Entity = (*Potion)(nil)

type Potion struct {
	*generic.Sprite
	healAmount uint8
	camera     *Camera
}

func NewPotion(x, y int, healAmount uint8, img *ebiten.Image, camera *Camera) *Potion {
	return &Potion{generic.NewSprite(img, x, y), healAmount, camera}
}

func (p *Potion) Update(environment generic.WorldView) gameStatus.GameStatus {
	rawPlayer, ok := environment.FindFirstEntity("Player")

	if !ok {
		log.Fatal("No player entity was found in potion")
	}

	player := rawPlayer.(*Player)

	if player.Object.IsInCollisionWith(p.Object) {
		player.Heal(p.healAmount)
	}

	return gameStatus.CONTINUE
}

func (p *Potion) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.GetX()), float64(p.GetY()))
	options.GeoM.Translate(float64(p.camera.GetX()), float64(p.camera.GetY()))

	screen.DrawImage(p.GetImage(), &options)
}

func (p *Potion) Tag() string {
	return "Potion"
}
