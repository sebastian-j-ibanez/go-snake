package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth  = 1020
	WindowHeight = 670
	SegmentSize  = 50
	BannerHeight = 50
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(int32(WindowWidth), int32(WindowHeight), "Go Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ticker := time.NewTicker(time.Millisecond * 200)

	engine := NewEngine()
	go engine.ProcessInput()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		engine.Draw()
		select {
		case <-ticker.C:
			engine.RunCycle()
		}
		engine.Draw()

		rl.EndDrawing()
	}
}
