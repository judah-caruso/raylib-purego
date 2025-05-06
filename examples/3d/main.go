package main

import (
	rl "github.com/judah-caruso/raylib-purego"
)

func main() {
	rl.InitWindow(960, 540, "3D")
	defer rl.CloseWindow()

	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 10, Y: 10, Z: 10},
		Up:         rl.Vector3{Y: 1},
		FovY:       45.0,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)
		{
			rl.DrawGrid(50, 10)
			rl.DrawLine3D(rl.Vector3{X: 0, Y: 0, Z: 0}, rl.Vector3{X: 100, Y: 100, Z: 100}, rl.Red)
			rl.DrawLine3D(rl.Vector3{X: 100, Y: 100, Z: 100}, rl.Vector3{X: 100, Y: 0, Z: 100}, rl.Red)
			rl.DrawLine3D(rl.Vector3{X: 100, Y: 0, Z: 100}, rl.Vector3{X: 0, Y: 0, Z: 0}, rl.Red)
		}
		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
