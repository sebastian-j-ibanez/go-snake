package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	snake  *Snake
	food   *Food
	border *Border
	score  int
}

func InitState() *State {
	segmentRadius := (SegmentSize / 2)
	windowCenterX := WindowWidth / 2
	windowCenterY := WindowHeight / 2
	startingX := windowCenterX - segmentRadius
	startingY := windowCenterY - segmentRadius
	snake := NewSnake(startingX, startingY, 0, 0)

	borderWidth := 20
	topBorderY := borderWidth
	bottomBorderY := WindowHeight - borderWidth
	leftBorderX := borderWidth
	rightBorderX := WindowWidth - borderWidth
	border := NewBorder(leftBorderX, topBorderY, rightBorderX, bottomBorderY)

	state := State{
		&snake,
		nil,
		border,
		0,
	}
	return &state
}

func (state *State) Loop() {
	state.HandleInput()
	if state.food == nil {
		GenerateFood(state)
	}
	state.snake.Move()
	if Collision(state.snake.head, state.food) {
		state.snake.Grow()
		state.food = nil
	}
}

func GenerateFood(state *State) {
	for {
		x := rand.IntN((state.border.x2-state.border.x1)/SegmentSize)*SegmentSize + state.border.x1
		y := rand.IntN((state.border.y2-state.border.y1)/SegmentSize)*SegmentSize + state.border.y1
		if !state.snake.Occupies(x, y) {
			state.food = &Food{x, y}
			return
		}
	}
}

func Collision(a Entity, b Entity) bool {
	return a.GetX() == b.GetX() && a.GetY() == b.GetY()
}

func (state *State) HandleInput() {
	head := state.snake.head
	if rl.IsKeyPressed(rl.KeyUp) && head.dirY == 0 {
		head.dirX = 0
		head.dirY = -1
	}

	if rl.IsKeyPressed(rl.KeyDown) && head.dirY == 0 {
		head.dirX = 0
		head.dirY = 1
	}
	if rl.IsKeyPressed(rl.KeyLeft) && head.dirX == 0 {
		head.dirX = -1
		head.dirY = 0
	}

	if rl.IsKeyPressed(rl.KeyRight) && head.dirX == 0 {
		head.dirX = 1
		head.dirY = 0
	}
}
