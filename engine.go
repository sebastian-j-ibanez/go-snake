package main

import (
	"fmt"
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
}

// Initialize a new game engine.
func NewEngine() *Engine {
	// Window & Snake
	borderWidth := 10
	windowCenterX := FullWindowWidth / 2
	windowCenterY := (FullWindowHeight - ScoreBoredHeight) / 2
	// Align snake starting position to the same grid as food
	startingX := borderWidth + ((windowCenterX - borderWidth) / SegmentSize * SegmentSize)
	startingY := borderWidth + ((windowCenterY - borderWidth) / SegmentSize * SegmentSize) // Border
	topBorderY := borderWidth + ScoreBoredHeight
	bottomBorderY := FullWindowHeight - borderWidth
	leftBorderX := borderWidth
	rightBorderX := FullWindowWidth - borderWidth
	border := NewBorder(leftBorderX, topBorderY, rightBorderX, bottomBorderY)

	// Score board
	scoreBoard := ScoreBoard{
		originX: 0,
		originY: 0,
		width:   FullWindowWidth,
		height:  ScoreBoredHeight,
		scoreX:  windowCenterX,
		scoreY:  5,
	}

	engine := Engine{
		new(NewSnake(startingX, startingY, 0, 0)),
		nil,
		border,
		0,
		scoreBoard,
	}
	return &engine
}

// Draw the engine entities.
func (engine *Engine) Draw() {
	engine.snake.Draw()
	engine.food.Draw()
	engine.border.Draw()
	engine.scoreBoard.Draw(engine.score)
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
	} else if BorderCollision(engine.border, engine.snake.head) {
		engine.PrintGameOver()
	}
}

// Get input and change snake direction accordingly.
func (s *Engine) ProcessInput() {
	for {
		x := &s.snake.head.dirX
		y := &s.snake.head.dirY
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

// Check collision between entity and border.
func BorderCollision(border *Border, a Entity) bool {
	inXBounds := a.GetX() > border.x1 && a.GetX() < border.x2
	inYBounds := a.GetY() > border.y1 && a.GetY() < border.y2
	return !inXBounds && !inYBounds
}

// Print game over message.
func (engine *Engine) PrintGameOver() {
	text := fmt.Sprintf("Game Over\nScore: %d", engine.score)
	rl.DrawRectangle(int32(engine.scoreBoard.scoreX+20), int32(engine.scoreBoard.scoreY+20), int32(engine.scoreBoard.width), int32(engine.scoreBoard.height), rl.DarkGray)
	rl.DrawText(text, int32(engine.scoreBoard.originX+20), int32(engine.scoreBoard.originY+20), 40, SegmentColor)
}
