package generic

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image *ebiten.Image
	*Object
}

func NewSprite(img *ebiten.Image, x, y int) *Sprite {
	rect := img.Bounds()
	fmt.Println(rect)
	fmt.Println(rect.Dx(), rect.Dy())
	return &Sprite{img, NewObject(x, y, rect.Dx(), rect.Dy())}
}

func (s Sprite) GetImage() *ebiten.Image {
	return s.image
}

// IsInCollisionWith does not work as expected
func (o Object) IsInCollisionWith(other *Object) bool {
	return o.GetX() > other.GetX() && o.GetX() < other.GetX()+other.width &&
		o.GetY() > other.GetY() && o.GetY() < other.GetY()+other.height
}
