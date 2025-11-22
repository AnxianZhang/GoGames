package main

import (
	"bytes"
	"errors"
	"image/color"
	"log"
	"time"

	"github.com/AnxianZhang/GoGames/common"
	"github.com/AnxianZhang/GoGames/common/gameStatus"
	"github.com/AnxianZhang/GoGames/entity"
	"github.com/AnxianZhang/GoGames/game"
	"github.com/AnxianZhang/GoGames/geometry"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	fontSource *text.GoTextFaceSource
)

var _ ebiten.Game = (*SnakeGame)(nil)

type SnakeGame struct {
	environment game.Environment
	lastUpdate  time.Time
	isGameOver  bool
}

func NewSnakeGame(environment game.Environment) *SnakeGame {
	return &SnakeGame{
		environment: environment,
		lastUpdate:  time.Now(),
	}
}

// Update function can have a diff update rate thant the Draw function
// This function is used to update the element on the screen
func (sg *SnakeGame) Update() error {
	if sg.isGameOver {
		return nil
	}

	rawSnake, ok := sg.environment.FindFirstEntity("Snake")

	if !ok {
		return errors.New("snake entity not found")
	}

	snake := rawSnake.(*entity.Snake)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyW):
		snake.SetDirection(geometry.UpDirection)
	case ebiten.IsKeyPressed(ebiten.KeyS):
		snake.SetDirection(geometry.DownDirection)
	case ebiten.IsKeyPressed(ebiten.KeyA):
		snake.SetDirection(geometry.LeftDirection)
	case ebiten.IsKeyPressed(ebiten.KeyD):
		snake.SetDirection(geometry.RightDirection)
		//case ebiten.IsKeyPressed(ebiten.KeyR):
		//	sg.isGameOver = false
		//	 = []geometry.Position{geometry.NewPositionWithOffSet(common.SCREEN_WIDTH, common.SCREEN_HEIGHT, 0, 0)}
		//	snake.SetDirection(geometry.NewGridPosition(0, 0)
	}
	if time.Since(sg.lastUpdate) < common.GAME_SPEED {
		return nil
	}

	// reset the timer ==> the refresh time rate
	sg.lastUpdate = time.Now()

	for _, e := range sg.environment.GetEntites() {
		var status gameStatus.GameStatus = e.Update(sg.environment)
		if status == gameStatus.LOSE {
			sg.isGameOver = true
			return nil
		}
	}

	return nil
}

func (sg *SnakeGame) showGameOverScreen(screen *ebiten.Image) {
	// le texte a afficher
	gameOverText := "Game over !"
	// definition de la taille de la police
	font := &text.GoTextFace{
		Source: fontSource,
		Size:   50, // tail tu text
	}

	// calcule les dimention du texte par rapport à la police
	w, h := text.Measure(gameOverText, font, font.Size)
	options := &text.DrawOptions{}

	options.GeoM.Translate((common.SCREEN_WIDTH-w)/2, (common.SCREEN_HEIGHT-h)/2)
	options.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, gameOverText, font, options)
}

// Draw function is used to represent UI on the screen
func (sg *SnakeGame) Draw(screen *ebiten.Image) {
	// l'ordre des représentations ici définit aussi l'ordre des layers
	// le premier s'affichera au 1e plan, le 2e au 2e plan, ...
	for _, e := range sg.environment.GetEntites() {
		e.Draw(screen)
	}

	if sg.isGameOver {
		sg.showGameOverScreen(screen)
	}
}

func (sg *SnakeGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.SCREEN_WIDTH, common.SCREEN_HEIGHT
}

func initialiseFaceSource() {
	// Police d'écriture importer. MPlus1 est une police par défaut fournis par ebiten
	font, err := text.NewGoTextFaceSource(bytes.NewBuffer(fonts.MPlus1pRegular_ttf))

	if err != nil {
		log.Fatal(err)
	}

	fontSource = font
}

func initialiseGame() *SnakeGame {
	environment := game.NewEnvironment()
	environment.AddEntity(
		entity.NewSnake(
			geometry.NewPositionWithOffSet(common.SCREEN_WIDTH, common.SCREEN_HEIGHT, 0, 0),
			geometry.NoDirection,
		),
	)

	for range 5 {
		environment.AddEntity(entity.NewFood())
	}

	return NewSnakeGame(*environment)
}

func main() {
	initialiseFaceSource()
	theGame := initialiseGame()

	ebiten.SetWindowSize(common.SCREEN_WIDTH, common.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Snake Game")
	//ebiten.SetTPS(1) // good for games based on grid

	if err := ebiten.RunGame(theGame); err != nil {
		log.Fatal(err)
	}
}
