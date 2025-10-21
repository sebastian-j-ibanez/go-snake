package main

import (
	"time"

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

	ticker := time.NewTicker(time.Millisecond * 350)

	engine := NewEngine()
	go engine.ProcessInput()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		select {
		case <-ticker.C:
			engine.RunCycle()
		}
		engine.Draw()

		rl.EndDrawing()
	}
}
