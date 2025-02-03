package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var num int = 0

func Init() {
	rl.InitWindow(800, 600, "Go Raylib 1")
	rl.SetTargetFPS(60)
}

func Run() {

	for !rl.WindowShouldClose() {
		update()
		render()
	}
	defer rl.CloseWindow()
}

func update() {
	num += 1
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Color{R: 135, G: 26, B: 250, A: 255})
	rl.DrawText("Hello, Raylib in Go! " + string(num), 200, 250, 20, rl.White)

	rl.EndDrawing()
}

