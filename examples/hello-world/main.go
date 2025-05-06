package main

import (
	rl "github.com/judah-caruso/raylib-purego"
)

func main() {
	rl.InitWindow(960, 540, "Hello, World")
	defer rl.CloseWindow()

	pos := rl.Vector2{}
	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()

		pos.X += 100 * delta
		pos.Y += 100 * delta

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangleV(pos, rl.Vector2{X: 50, Y: 50}, rl.Red)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
