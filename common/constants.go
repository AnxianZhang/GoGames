package common

// BALL_SPEED and PADDLE_SPEED are, how many pixel the object will move per game update
// by default ebiten update the game state 60 times per second.
// each update of the game is called a tick
const (
	BALL_SPEED   = 3
	PADDLE_SPEED = 6 // = move 6 pixels per tick

	PLAYER_SPEED = 2
)
