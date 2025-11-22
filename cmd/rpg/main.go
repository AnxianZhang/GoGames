package main

import (
	"image"
	"image/color"
	"log"

	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/common/tiles"
	"github.com/AnxianZhang/GoGames/entity"
	"github.com/AnxianZhang/GoGames/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*RPG)(nil)

type RPG struct {
	*game.Environment
	tilemap      *tiles.TilemapJSON
	tilemapImage *ebiten.Image
	camera       *entity.Camera
}

func NewRpg(_tilemap *tiles.TilemapJSON, tilemapImage *ebiten.Image, camera *entity.Camera) *RPG {
	return &RPG{game.NewEnvironment(), _tilemap, tilemapImage, camera}
}

// Update is fixed to 60 ticks per second
func (r *RPG) Update() error {
	rawPlayer, ok := r.FindFirstEntity("Player")

	if !ok {
		log.Fatal("No player entity was found")
	}

	player := rawPlayer.(*entity.Player)

	speed := common.PLAYER_SPEED

	// /!\ using switch isn't a bad idea, but switch only allow one solution
	// so that player can only move in horizontal or vertical.
	// left and right
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		player.MoveLeft(speed)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		player.MoveRight(speed)
	}

	// up and down
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		player.MoveUp(speed)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		player.MoveDown(speed)
	}

	for _, s := range r.GetEntites() {
		s.Update(r.Environment)
	}

	r.camera.LimitToBorder(
		25*16,
		25*16,
		r.Environment,
	)

	return nil
}

func (r *RPG) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 120, G: 180, B: 255, A: 255})

	options := ebiten.DrawImageOptions{}

	rawCamera, ok := r.FindFirstEntity("Camera")
	if !ok {
		log.Fatal("No camera was found in main update function")
	}

	camera := rawCamera.(*entity.Camera)

	// loop over the layer
	for _, layer := range r.tilemap.Layers {
		// loop over the tiles
		for idx, num := range layer.Data {
			// get x, y position of a tile, then convert the position in pixels
			x := (idx % layer.Width) * 16
			y := (idx / layer.Width) * 16

			// get the x, y of the actual tile image, then convert it to pixels
			srcX := ((num - 1) % 22) * 16 // 22 items on a row
			srcY := ((num - 1) / 22) * 16 // 22 items on a row

			options.GeoM.Translate(float64(x), float64(y))
			options.GeoM.Translate(float64(camera.GetX()), float64(camera.GetY()))

			// Draw the actual tile
			screen.DrawImage(
				r.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&options,
			)

			options.GeoM.Reset()
		}
	}

	// draw image
	for _, s := range r.GetEntites() {
		s.Draw(screen)
	}
}

func (r *RPG) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return common.RPG_WIDTH_LAYOUT, common.RPG_HEIGHT_LAYOUT
}

func initialiseGameEntities() *RPG {
	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/images/hunter.png")
	skeletonImg, _, err := ebitenutil.NewImageFromFile("./assets/images/skeleton.png")
	potionImg, _, err := ebitenutil.NewImageFromFile("./assets/images/lifePot.png")

	tilemapImage, _, err := ebitenutil.NewImageFromFile("./assets/images/tilesetFloor.png")
	tilemap, err := tiles.NewTilemapJSON("./assets/maps/spawn.json")

	camera := entity.NewCamera()

	rpgGame := NewRpg(tilemap, tilemapImage, camera)

	if err != nil {
		log.Fatal(err)
	}

	rpgGame.AddEntity(camera)

	rpgGame.AddEntity(
		entity.NewPotion(300, 50, 2, potionImg, camera),
	).AddEntity(
		entity.NewPlayer(100, 200, playerImg, 5, camera),
	)

	// Add enemies
	rpgGame.AddEntity(
		entity.NewEnemy(10, 10, skeletonImg, true, camera),
	).AddEntity(
		entity.NewEnemy(200, 188, skeletonImg, false, camera),
	).AddEntity(
		entity.NewEnemy(59, 153, skeletonImg, false, camera),
	)

	return rpgGame
}

func main() {
	rpgGame := initialiseGameEntities()

	ebiten.SetWindowTitle("RPG")
	ebiten.SetWindowSize(common.SCREEN_WIDTH, common.SCREEN_HEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(rpgGame); err != nil {
		log.Fatal(err)
	}
}
