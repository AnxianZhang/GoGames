package entity

import (
	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/math"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// this syntax is a convention in ebiten, implicitly tel go that Snake should implement Entity interface
var _ Entity = (*Snake)(nil)

type Snake struct {
	body      []math.Position
	direction math.Position // can have an x and y acceleration, the snake will move along this direction
}

func NewSnake(start, direction math.Position) *Snake {
	return &Snake{[]math.Position{start}, direction}
}

func (s *Snake) Update(environment WorldView) bool {
	// update stake state
	head := s.getHead()
	newHeadPosition := head.Add(s.direction)

	// if it is the case, lose the game
	if head.IsInCollisionWithScreen(s.body[1:]) {
		return true
	}

	var needToGrow = false

	for _, foodEntity := range environment.SearchEntities("food") {
		food := foodEntity.(*Food)

		if head.IsEqualTo(food.GetPosition()) {
			needToGrow = true
			food.Respwan()
		}
	}

	if needToGrow {
		s.body = append([]math.Position{newHeadPosition}, s.body[:len(s.body)]...)
	} else {
		s.body = append([]math.Position{newHeadPosition}, s.body[:len(s.body)-1]...)
	}

	return false
}

func (s Snake) Draw(screen *ebiten.Image) {
	// draw snake head
	vector.FillRect(
		screen,
		float32(s.getHead().GetX()*common.GRID_SIZE),
		float32(s.getHead().GetY()*common.GRID_SIZE),
		common.GRID_SIZE, common.GRID_SIZE,
		color.RGBA{R: 0, G: 255},
		true,
	)

	// draw snake
	for _, p := range s.body[1:] {
		vector.FillRect(
			screen,
			float32(p.GetX()*common.GRID_SIZE),
			float32(p.GetY()*common.GRID_SIZE),
			common.GRID_SIZE, common.GRID_SIZE,
			color.White,
			true,
		)
	}
}

func (s Snake) Tag() string {
	return "Snake"
}

func (s *Snake) SetDirection(newDirection math.Position) {
	s.direction = newDirection
}

func (s Snake) getHead() math.Position {
	return s.body[0]
}
