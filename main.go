package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FullWindowWidth  = 1020
	FullWindowHeight = 670
	SegmentSize      = 50
	ScoreBoredHeight = 50
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(int32(FullWindowWidth), int32(FullWindowHeight), "Go Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ticker := time.NewTicker(time.Millisecond * 200)

	engine := NewEngine()
	go engine.ProcessInput()

	for !rl.WindowShouldClose() {
		select {
		case <-ticker.C:
			engine.RunCycle()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		engine.Draw()
		rl.EndDrawing()
	}
}
