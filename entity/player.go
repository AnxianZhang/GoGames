package entity

import (
	"image"
	"strconv"

	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ generic.Entity = (*Player)(nil)

type Player struct {
	*generic.Sprite
	health uint8
}

func NewPlayer(x, y int, img *ebiten.Image, health uint8) *Player {
	return &Player{generic.NewSprite(img, x, y), health}
}

func (p *Player) Update(environment generic.WorldView) gameStatus.GameStatus {
	return gameStatus.CONTINUE
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.Itoa(int(p.health)))

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.GetX()), float64(p.GetY()))

	screen.DrawImage(p.GetImage().SubImage(
		image.Rect(0, 0, 16, 16),
	).(*ebiten.Image), &options)
}

func (p *Player) Tag() string {
	return "Player"
}

func (p *Player) Heal(amount uint8) {
	p.health += amount
}

func (p *Player) Damage(amount uint8) {
	if p.health-amount <= 0 {
		p.health = 0
	} else {
		p.health -= amount
	}
}
