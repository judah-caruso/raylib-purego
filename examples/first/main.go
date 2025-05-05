package main

import (
	rl "github.com/judah-caruso/raylib-purego"
)

func main() {
	rl.InitWindow(960, 540, "First")
	defer rl.CloseWindow()

	pos := rl.Vector2{}
	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		pos.X += 100 * delta
		pos.Y += 100 * delta

		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{})
		rl.DrawRectangle(int(pos.X), int(pos.Y), 50, 50, rl.Color{255, 100, 100, 255})
		rl.EndDrawing()
	}
}
