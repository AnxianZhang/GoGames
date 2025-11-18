package entity

import "github.com/hajimehoshi/ebiten/v2"

// An entity is a thing that might be alive in our game
type Entity interface {
	Update(environment WorldView) bool
	Draw(screen *ebiten.Image)

	// Tag is used to identify the identity type
	Tag() string
}
