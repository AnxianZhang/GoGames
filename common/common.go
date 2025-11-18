package common

import "time"

const (
	GAME_SPEED    = time.Second / 6
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
	GRID_SIZE     = 20
	X_CASE        = SCREEN_WIDTH / GRID_SIZE
	Y_CASE        = SCREEN_HEIGHT / GRID_SIZE
)
