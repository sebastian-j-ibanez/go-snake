package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var SegmentColor = rl.NewColor(0, 173, 216, 255)
var FoodColor = rl.NewColor(255, 117, 20, 255)

func (state *State) DrawState() {
	DrawSnake(state.snake)
	DrawFood(state.food)
	// DrawBorder(state.border)
}

func DrawSnake(snake *Snake) {
	for current := snake.head; current != nil; current = current.next {
		rl.DrawRectangle(
			int32(current.x),
			int32(current.y),
			SegmentSize,
			SegmentSize,
			SegmentColor,
		)
	}
}

func DrawFood(f *Food) {
	if f != nil {
		rl.DrawRectangle(int32(f.x), int32(f.y), SegmentSize, SegmentSize, FoodColor)
	}
}

func DrawBorder(border *Border) {
	width := border.x2 - border.x1
	height := border.y2 - border.y1
	rl.DrawRectangleLines(int32(border.x1), int32(border.y1), int32(width), int32(height), rl.Black)
}
