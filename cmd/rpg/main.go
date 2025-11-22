package main

import (
	"image/color"
	"log"

	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/entity"
	"github.com/AnxianZhang/GoGames/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*RPG)(nil)

type RPG struct {
	*game.Environment
}

func NewRpg() *RPG {
	return &RPG{game.NewEnvironment()}
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

	return nil
}

func (r *RPG) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 120, G: 180, B: 255, A: 255})

	// draw image
	for _, s := range r.GetEntites() {
		s.Draw(screen)
	}
}

func (r *RPG) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func initialiseGameEntities() *RPG {
	rpgGame := NewRpg()

	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/images/hunter.png")
	skeletonImg, _, err := ebitenutil.NewImageFromFile("./assets/images/skeleton.png")
	potionImg, _, err := ebitenutil.NewImageFromFile("./assets/images/lifePot.png")

	if err != nil {
		log.Fatal(err)
	}

	rpgGame.AddEntity(
		entity.NewPotion(300, 50, 2, potionImg),
	).AddEntity(
		entity.NewPlayer(100, 200, playerImg, 5),
	)

	// Add enemies
	rpgGame.AddEntity(
		entity.NewEnemy(10, 10, skeletonImg, true),
	).AddEntity(
		entity.NewEnemy(200, 188, skeletonImg, false),
	).AddEntity(
		entity.NewEnemy(59, 153, skeletonImg, false),
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
