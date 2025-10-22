package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Represents the game engine.
type Engine struct {
	snake  *Snake
	food   *Food
	border *Border
	score  int
}

// Initialize a new game engine.
func NewEngine() *Engine {
	// Window & Snake
	borderWidth := 10
	windowCenterX := WindowWidth / 2
	windowCenterY := (WindowHeight - BannerHeight) / 2
	// Align snake starting position to the same grid as food
	startingX := borderWidth + ((windowCenterX - borderWidth) / SegmentSize * SegmentSize)
	startingY := borderWidth + ((windowCenterY - borderWidth) / SegmentSize * SegmentSize)
	snake := NewSnake(startingX, startingY, 0, 0)

	// Border
	topBorderY := borderWidth + BannerHeight
	bottomBorderY := WindowHeight - borderWidth
	leftBorderX := borderWidth
	rightBorderX := WindowWidth - borderWidth
	border := NewBorder(leftBorderX, topBorderY, rightBorderX, bottomBorderY)

	// Input channel.

	engine := Engine{
		&snake,
		nil,
		border,
		0,
	}
	return &engine
}

// Draw the engine entities.
func (engine *Engine) Draw() {
	engine.snake.Draw()
	engine.food.Draw()
	engine.border.Draw()
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
