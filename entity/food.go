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

// explicitly tell go that Food should implement Entity
var _ generic.Entity = (*Food)(nil)

type Food struct {
	position *math.Position
}

func NewFood() *Food {
	return &Food{math.RandomPosition()}
}

func (f Food) GetPosition() *math.Position {
	return f.position
}

func (f Food) Update(environment generic.WorldView) gameStatus.GameStatus {
	return gameStatus.CONTINUE
}

func (f Food) Draw(screen *ebiten.Image) {
	// drow food
	vector.FillRect(
		screen,
		float32(f.position.GetX()*common.GRID_SIZE), // * la taille de la grille pour avoir la position initial du premiere pixel,
		float32(f.position.GetY()*common.GRID_SIZE),
		common.GRID_SIZE, common.GRID_SIZE, // fournis pour dire que chaque x,y forme enemble un bloque de 20px * 20px
		color.RGBA{R: 255, A: 255},
		true,
	)
}

func (f Food) Tag() string {
	return "food"
}

func (f *Food) Respwan() {
	f.position = math.RandomPosition()
}
