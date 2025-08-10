package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth  = 1500
	WindowHeight = 1000
	SegmentSize  = 50
)

func main() {
	rl.InitWindow(int32(WindowWidth), int32(WindowHeight), "Go Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	state := InitState()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		DrawBorder(state.border)

		state.Loop()
		state.DrawState()

		rl.EndDrawing()
	}
}
