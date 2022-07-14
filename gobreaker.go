package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type brick struct {
	height int32
	width  int32
	xPos   int32
	yPos   int32
	alive  bool
}

type player struct {
	playerPos int32
	speed     int32
	height    int32
	width     int32
}

var (
	windowName   string = "Go Breaker"
	windowWidth  int32  = 800
	windowHeight int32  = 450

	fps int32 = 60

	playerPos    int32 = 10
	playerSpeed  int32 = 700
	playerHeight int32 = 10
	playerWidth  int32 = 100

	ballX     float32 = 50
	ballY     float32 = 50
	ballXVel  float32 = 1
	ballYVel  float32 = 1
	ballSpeed float32 = 250

	targets = []brick{
		{
			height: 20,
			width:  50,
			xPos:   0,
			yPos:   0,
			alive:  true,
		},
		{
			height: 20,
			width:  50,
			xPos:   55,
			yPos:   0,
			alive:  true,
		},
		{
			height: 20,
			width:  50,
			xPos:   110,
			yPos:   0,
			alive:  true,
		},
		{
			height: 20,
			width:  50,
			xPos:   165,
			yPos:   0,
			alive:  true,
		},
	}
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowName)

	rl.SetTargetFPS(int32(fps))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		update()
		render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func update() {
	updateBall()
	updatePlayer()
}

func render() {
	renderPlayer()
	renderBall()
	renderTargets()
}

func updatePlayer() {
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

func renderTargets() {
	for _, target := range targets {
		if target.alive == true {
			rl.DrawRectangle(target.xPos, target.yPos, target.width, target.height, rl.Green)
		}
	}
}

func updateBall() {
	ballX += ballXVel * ballSpeed * rl.GetFrameTime()
	ballY += ballYVel * ballSpeed * rl.GetFrameTime()

	if ballX >= float32(windowWidth) || ballX <= float32(0) {
		ballXVel = -ballXVel
	}

	if ballY <= float32(0) {
		ballYVel = -ballYVel
	}

	playerRec := rl.Rectangle{
		Width:  float32(playerWidth),
		Height: float32(playerHeight),
		X:      float32(playerPos),
		Y:      float32(windowHeight - playerHeight),
	}

	ballRec := rl.Rectangle{
		Width:  25,
		Height: 25,
		X:      ballX,
		Y:      ballY,
	}

	if ballY >= float32(windowHeight) {
		rl.CloseWindow()
	}

	if rl.CheckCollisionRecs(playerRec, ballRec) {
		ballYVel = -ballYVel
	}

	for i, target := range targets {
		targetRec := rl.Rectangle{
			Width:  25,
			Height: 25,
			X:      float32(target.xPos),
			Y:      float32(target.yPos),
		}

		if rl.CheckCollisionRecs(targetRec, ballRec) && target.alive == true {
			ballYVel = -ballYVel
			targets[i].alive = false
		}
	}

}
