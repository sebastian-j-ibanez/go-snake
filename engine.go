package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Represents the game engine.
type Engine struct {
	snake      *Snake
	food       *Food
	border     *Border
	score      int
	scoreBoard ScoreBoard
	run        bool
}

// Initialize a new game engine.
func NewEngine() *Engine {
	// Window & Snake
	windowCenterX := FullWindowWidth / 2
	windowCenterY := (FullWindowHeight - ScoreBoardHeight) / 2
	// Align snake starting position to the same grid as food
	startingX := BorderWidth + ((windowCenterX - BorderWidth) / SegmentSize * SegmentSize)
	startingY := BorderWidth + ((windowCenterY - BorderWidth) / SegmentSize * SegmentSize) // Border
	topBorderY := BorderWidth + ScoreBoardHeight
	bottomBorderY := FullWindowHeight - BorderWidth
	leftBorderX := BorderWidth
	rightBorderX := FullWindowWidth - BorderWidth
	border := NewBorder(leftBorderX, topBorderY, rightBorderX, bottomBorderY)

	// Score board
	scoreBoard := ScoreBoard{
		originX: 0,
		originY: 0,
		width:   FullWindowWidth,
		height:  ScoreBoardHeight,
		scoreX:  100,
		scoreY:  20,
	}

	engine := Engine{
		new(NewSnake(startingX, startingY, 0, 0)),
		nil,
		border,
		0,
		scoreBoard,
		true,
	}
	return &engine
}

// Draw the engine entities.
func (engine *Engine) Draw() {
	engine.snake.Draw()
	engine.food.Draw()
	engine.border.Draw()
	engine.scoreBoard.Draw(engine.score)
	if !engine.run {
		engine.DrawGameOver()
	}
}

// Run one cycle of game logic.
func (engine *Engine) RunCycle() {
	if engine.food == nil {
		GenerateFood(engine)
	}
	engine.snake.Move()
	if Collision(engine.snake.head, engine.food) {
		engine.snake.Grow()
		engine.food = nil
		engine.score += 1
	} else if IsOutOfBounds(engine.border, engine.snake.head) {
		engine.run = false
	}
}

// Get input and change snake direction accordingly.
func (engine *Engine) ProcessInput() {
	for {
		if !engine.run {
			return
		}

		x := &engine.snake.head.dirX
		y := &engine.snake.head.dirY
		if rl.IsKeyPressed(rl.KeyUp) && *y == 0 {
			*x = 0
			*y = -1
		} else if rl.IsKeyPressed(rl.KeyDown) && *y == 0 {
			*x = 0
			*y = 1
		} else if rl.IsKeyPressed(rl.KeyLeft) && *x == 0 {
			*x = -1
			*y = 0
		} else if rl.IsKeyPressed(rl.KeyRight) && *x == 0 {
			*x = 1
			*y = 0
		}
	}
}

// Generate a new food at a random location on screen.
func GenerateFood(engine *Engine) {
	for {
		x := rand.IntN((engine.border.x2-engine.border.x1)/SegmentSize)*SegmentSize + engine.border.x1
		y := rand.IntN((engine.border.y2-engine.border.y1)/SegmentSize)*SegmentSize + engine.border.y1
		if !engine.snake.Occupies(x, y) {
			engine.food = &Food{x, y}
			return
		}
	}
}

// Check collision between entities.
func Collision(a Entity, b Entity) bool {
	return a.GetX() == b.GetX() && a.GetY() == b.GetY()
}

// Check if entity is out of bounds.
func IsOutOfBounds(border *Border, a Entity) bool {
	outOfXBounds := a.GetX() < border.x1 || a.GetX() >= border.x2
	outOfYBounds := a.GetY() < border.y1 || a.GetY() > border.y2-SegmentSize
	return outOfXBounds || outOfYBounds
}

// Draw game over message inside the render loop.
func (engine *Engine) DrawGameOver() {
	text := "Game Over"
	textWidth := rl.MeasureText(text, FontSize)
	msgX := (FullWindowWidth - BorderWidth) - (textWidth + BorderWidth)
	rl.DrawText(text, int32(msgX), int32(engine.scoreBoard.scoreY), FontSize, SegmentColor)
}
