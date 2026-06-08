package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FullWindowWidth  = 1020
	FullWindowHeight = 670
	SegmentSize      = 50
	ScoreBoardHeight = 50
	BorderWidth      = 10
	FontSize         = 40
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(int32(FullWindowWidth), int32(FullWindowHeight), "Go Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	ticker := time.NewTicker(time.Millisecond * 200)

	engine := NewEngine()

	go engine.ProcessInput()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		engine.Draw()
		rl.EndDrawing()

		if engine.run {
			select {
			case <-ticker.C:
				engine.RunCycle()
			default:
			}
		} else if rl.IsKeyPressed(rl.KeyEnter) {
			break
		}
	}
}
