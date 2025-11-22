package entity

import (
	"log"
	"math"

	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity/generic"
	"github.com/AnxianZhang/GoGames/game"
	"github.com/AnxianZhang/GoGames/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ generic.Entity = (*Camera)(nil)

type Camera struct {
	*geometry.Position
}

func NewCamera() *Camera {
	return &Camera{geometry.NewGridPosition(0, 0)}
}

func (c *Camera) Update(environment generic.WorldView) gameStatus.GameStatus {
	return gameStatus.CONTINUE
}

func (c *Camera) Draw(screen *ebiten.Image) {
}

func (c *Camera) Tag() string {
	return "Camera"
}

func (c *Camera) LimitToBorder(tileMapWidth, tileMapHeight int, environment *game.Environment) {
	rawPlayer, ok := environment.FindFirstEntity("Player")
	if !ok {
		log.Fatal("Error when trying to search Player entity in camera")
	}

	player := rawPlayer.(*Player)

	// 8 is added as offsets to center the camera to the middle of the player
	// initially the origin is at the left corner
	c.SetX(-player.GetX() + 8 + common.RPG_WIDTH_LAYOUT/2)
	c.SetY(-player.GetY() + 8 + common.RPG_HEIGHT_LAYOUT/2)

	c.SetX(int(math.Min(float64(c.GetX()), 0.0)))
	c.SetY(int(math.Min(float64(c.GetY()), 0.0)))

	c.SetX(int(math.Max(float64(c.GetX()), float64(common.RPG_WIDTH_LAYOUT-tileMapWidth))))
	c.SetY(int(math.Max(float64(c.GetY()), float64(common.RPG_HEIGHT_LAYOUT-tileMapHeight))))
}
