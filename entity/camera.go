package entity

import "github.com/AnxianZhang/GoGames/math"

type Camera struct {
	*math.Position
}

func NewCamera(x, y int) *Camera {
	return &Camera{math.NewGridPosition(x, y)}
}
