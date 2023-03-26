package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

const (
	width  = 50
	height = 10

	emptyCell = ' '
	ballCell  = '🏐'

	maxFrames   = 1200
	refreshRate = time.Second / 20
)

func main() {
	var (
		px, py int    // position of the ball
		vx, vy = 1, 1 // velocity of the ball. i.e. where it should move
		board  = make([][]bool, width)
	)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()
	for i := 0; i < maxFrames; i++ {
		calcBallPos(&px, &vx, &py, &vy)
		cleanBoard(&board)
		drawBoardWithBall(&board, &px, &py)
		time.Sleep(refreshRate)
	}

}

func cleanBoard(board *[][]bool) {
	b := *board
	for y := range b[0] {
		for x := range b {
			b[x][y] = false
		}
	}
}

func calcBallPos(px *int, vx *int, py *int, vy *int) {
	*px += *vx
	*py += *vy
	if *px <= 0 || *px >= width-1 {
		*vx *= -1
	}

	if *py <= 0 || *py >= height-1 {
		*vy *= -1
	}
}

func drawBoardWithBall(board *[][]bool, px *int, py *int) {
	b := *board
	buf := make([]rune, 0, width*height)
	buf = buf[:0]
	b[*px][*py] = true
	for y := range b[0] {
		for x := range b {
			cell := emptyCell
			if b[x][y] {
				cell = ballCell
			}
			buf = append(buf, cell, ' ')
		}
		buf = append(buf, '\n')
	}
	screen.MoveTopLeft()
	fmt.Println(string(buf))
}
