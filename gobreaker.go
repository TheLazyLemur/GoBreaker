package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	windowName   string = "Go Breaker"
	windowWidth  int32  = 800
	windowHeight int32  = 450

	fps int32 = 60

	playerPos    int32 = 10
	playerSpeed  int32 = 500
	playerHeight int32 = 50
	playerWidth  int32 = 100

	ballX     float32 = 10
	ballY     float32 = 10
	ballXVel  float32 = 1
	ballYVel  float32 = 1
	ballSpeed float32 = 500
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowName)

	rl.SetTargetFPS(int32(fps))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		renderPlayer()

		renderBall()

		updateBall()

		getPlayerInput()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func getPlayerInput() {
	if rl.IsKeyDown(rl.KeyRight) {
		playerPos += int32(float32(playerSpeed) * rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		playerPos += int32(-float32(playerSpeed) * rl.GetFrameTime())
	}
}

func renderPlayer() {
	rl.DrawRectangle(int32(playerPos), windowHeight-playerHeight, playerWidth, playerHeight, rl.Red)
}

func renderBall() {
	rl.DrawRectangle(int32(ballX), int32(ballY), 25, 25, rl.Purple)
}

func updateBall() {
	ballX += ballXVel * ballSpeed * rl.GetFrameTime()
	ballY += ballYVel * ballSpeed * rl.GetFrameTime()

	if ballX >= float32(windowWidth) || ballX <= float32(0) {
		ballXVel = -ballXVel
	}

	if ballY >= float32(windowHeight) || ballY <= float32(0) {
		ballYVel = -ballYVel
	}

}
