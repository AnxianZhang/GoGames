package geometry

import "math/rand/v2"

type Velocity struct {
	direction *Position
}

func NewVelocity(x, y int) *Velocity {
	return &Velocity{
		NewGridPosition(x, y),
	}
}

func GetRandomDirection() int {
	if rand.IntN(2) == 0 {
		return -1
	}
	return 1
}

func (v Velocity) GetX() int {
	return v.direction.GetX()
}

func (v Velocity) GetY() int {
	return v.direction.GetY()
}

func (v *Velocity) SetX(_x int) {
	v.direction.SetX(_x)
}

func (v *Velocity) SetY(_y int) {
	v.direction.SetY(_y)
}
