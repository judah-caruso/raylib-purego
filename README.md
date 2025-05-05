# Raylib Purego

Partially handwritten [purego](https://github.com/ebitengine/purego) bindings for [raylib](https://www.raylib.com/).

## Getting Started

Run `go get github.com/judah-caruso/raylib-purego@latest` and have fun.

```go
package main

import (
	rl "github.com/judah-caruso/raylib-purego"
)

func main() {
	rl.InitWindow(960, 540, "Hello, World")
	defer rl.CloseWindow()

	pos := rl.Vector2{}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{ 0, 0, 0, 255 })
		rl.EndDrawing()
	}
}
```

## Why

I'd like to use Raylib in Go without having to fiddle with globally installed libraries or CGO interop. Luckily for me [purego](https://github.com/ebitengine/purego) and [libffi](https://github.com/judah-caruso/ffi-embeded) exist and allow me to do so without much hassle.

## How

At initialzation time raylib is loaded, then each procedure is bound using libffi which allows us to call it without going through CGo. It additionally allows us to pass/return structs by value which is [not (easily) possible](https://github.com/ebitengine/purego/issues/237) with purego.

Because we're using libffi, we need to describe each procedure and type in raylib before we're able to call into it. This is where `bindings/generate.go` comes in. It simply parses the files that describe the procedures and types within raylib, and generates ffi types and Go wrappers.

Note: This is mostly a proof of concept to see how feasible my idea was. Because of that, most of raylib isn't bound and the binding generation code is just bad. However, with the current systme in place, it shouldn't be very difficult to complete this library.

## Q&A

- **Q: How does the performance compare to [raylib-go](github.com/gen2brain/raylib-go)?**
- A: I have not tested this and have no intentions to do so (unless it ends up being dramatically slower). This is because call overhead will exist in both libraries, and the overhead introduced by CGo/purego will not generally be a bottleneck in your game.
- **Q: How about raymath?**
- A: I will most likely reimplement that part of the library in Go so the compiler can optimize Vector, Matrix, etc. operations (something it can't do when using CGo/purego).
- **Q: Stuff is missing! How do I add it myself?**
- A: If a type is missing, add it to `raylib.go` and re-run `go generate`. If a procedure is missing, add it to the specific raylib module file (`rcore.go`, `rshapes.go`, etc.) and re-run `go generate`. I've made the `prep` procedure take something that looks roughly like a C procedure declaration to make it easier to bind things. The comments after the calls to `prep` are "magic" in that they will be attached to the generated Go wrapper. The comma-separated list at the start of the comment (before `::`) will be used to name the arguments in the Go wrapper. If you add anything, please submit a pull request!
- **Q: Will WebAssembly be supported?**
- A: Probably not, but I won't rule it out completely. The library would most likely have to change quite a bit. If WebAssembly is required for your project, you could use `raylib-purego` as a base and do something like [this](https://github.com/gen2brain/raylib-go/issues/356#issuecomment-1967521031). If you end up getting this to work, let me know! It'd be cool to see.
