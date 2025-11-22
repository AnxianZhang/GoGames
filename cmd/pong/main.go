package main

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"log"
	"math/rand"

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

var _ ebiten.Game = (*PongGame)(nil)

type PongGame struct {
	game.Environment
	score      uint16
	highScore  uint16
	isGameOver bool
}

func NewPongGame() *PongGame {
	return &PongGame{}
}

func (pg *PongGame) Update() error {
	rawPaddle, ok := pg.FindFirstEntity("Paddle")

	if !ok {
		return errors.New("no paddle found")
	}

	paddle := rawPaddle.(*entity.Paddle)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyW):
		paddle.MoveUp(common.PADDLE_SPEED)
	case ebiten.IsKeyPressed(ebiten.KeyS):
		paddle.MoveDown(common.PADDLE_SPEED)
	}

	for _, e := range pg.GetEntites() {
		var status gameStatus.GameStatus = e.Update(pg.Environment)

		switch status {
		case gameStatus.GET_POINT:
			pg.score++
			if pg.score > pg.highScore {
				pg.highScore = pg.score
			}
		case gameStatus.LOSE:
			pg.isGameOver = true
			pg.score = 0
		case gameStatus.CONTINUE:
		default:
			panic("unhandled default case go game status in update func")
		}
	}

	return nil
}

func (pg *PongGame) Draw(screen *ebiten.Image) {
	for _, e := range pg.GetEntites() {
		e.Draw(screen)
	}

	scoreText := fmt.Sprintf("Score: %v", pg.score)
	highText := fmt.Sprintf("High score: %v", pg.highScore)
	drawOptions := &text.DrawOptions{}

	font := &text.GoTextFace{
		Source: fontSource,
		Size:   10,
	}

	drawOptions.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, highText, font, drawOptions)
	drawOptions.GeoM.Translate(0, 20) // the current score will be bellow the high score
	text.Draw(screen, scoreText, font, drawOptions)
}

func (pg *PongGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.SCREEN_WIDTH, common.SCREEN_HEIGHT
}

func initialisePongGame() *PongGame {
	pongGame := NewPongGame()

	pongGame.AddEntity(
		entity.NewPaddle(600, 200, 15, 100),
	)
	//.AddEntity(entity.NewBall(
	//	0, 0,
	//	rand.Intn(common.BALL_SPEED), common.BALL_SPEED,
	//	15, 15,
	//)).AddEntity(entity.NewBall(
	//	50, 50,
	//	common.BALL_SPEED, common.BALL_SPEED,
	//	10, 10,
	//))

	for range 50 {
		pongGame.AddEntity(entity.NewBall(
			rand.Intn(common.SCREEN_WIDTH), rand.Intn(common.SCREEN_HEIGHT),
			geometry.GetRandomDirection()*common.BALL_SPEED, geometry.GetRandomDirection()*common.BALL_SPEED,
			10, 10,
		))
	}

	return pongGame
}

func initialiseFaceSource() {
	// Police d'écriture importer. MPlus1 est une police par défaut fournis par ebiten
	font, err := text.NewGoTextFaceSource(bytes.NewBuffer(fonts.MPlus1pRegular_ttf))

	if err != nil {
		log.Fatal(err)
	}

	fontSource = font
}

func main() {
	initialiseFaceSource()
	pongGame := initialisePongGame()

	ebiten.SetWindowTitle("Pong game")
	ebiten.SetWindowSize(common.SCREEN_WIDTH, common.SCREEN_HEIGHT)

	if err := ebiten.RunGame(pongGame); err != nil {
		log.Fatal(err)
	}
}
