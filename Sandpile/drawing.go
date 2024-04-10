package main

import(
  "canvas"
)

func DrawPNGFinal(board Gameboard, cellWidth int, outfile string){

  height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := canvas.CreateNewCanvas(width, height)

	// declare colors
	black := canvas.MakeColor(0, 0, 0)
	darkGray := canvas.MakeColor(85, 85, 85)
  lightGray := canvas.MakeColor(170, 170, 170)
	white := canvas.MakeColor(255, 255, 255)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				c.SetFillColor(black)
			} else if board[i][j] == 1 {
				c.SetFillColor(darkGray)
			} else if board[i][j] == 2 {
				c.SetFillColor(lightGray)
			} else if board[i][j] >= 3 {
				c.SetFillColor(white)
			} else {
				panic("Error: Out of range value in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

  c.SaveToPNG(outfile)
}
