/*
package main

import (
	"fmt"
)

func main() {


	fmt.Println("Let's hack the Gray-Scott model!")

	fmt.Println("Calling DiffuseBoardOneParticle()...")
	//DiffuseBoardOneParticle(currentBoard [][]float64, kernel [3][3]float64) [][]float64

	firstBoard := [][]float64{
		{0,0,0,0,0} ,
		{0,0,0,0,0} ,
		{0,0,1,0,0} ,
		{0,0,0,0,0} ,
		{0,0,0,0,0} ,
	}

	k := [3][3]float64{
		{0.05,0.2,0.05} ,
		{0.2,-1,0.2} ,
		{0.05,0.2,0.05} ,
	}

	secondBoard := DiffuseBoardOneParticle(firstBoard, 0.1, k)
	fmt.Println("second board: ", secondBoard)


	fmt.Println(SumCells([2]float64{-4.5,1},[2]float64{2,-16.31}))

	cB := Board{
		{{0.2,0.7},{0,0},{0.1,0.4},} ,
		{{0.9,0.6},{1,1},{0,1},} ,
		{{1,0},{0.5,0.2},{0.3,0.3},},
	}


	fmt.Println("Change due to diffusion: ",ChangeDueToDiffusion(cB, 1, 1, 0.2, 0.1, k))


	cb2 := Board{
		{{1,0},{1,0},{1,0},{1,0},{1,0},} ,
		{{1,0},{1,0},{1,0},{1,0},{1,0},} ,
		{{1,0},{1,0},{1,1},{1,0},{1,0},},
		{{1,0},{1,0},{1,0},{1,0},{1,0},},
		{{1,0},{1,0},{1,0},{1,0},{1,0},},
	}
	fmt.Println("Update board: ",UpdateBoard(cb2, 0.3, 0.4, 0.2, 0.1, k))


	mycheck := SimulateGrayScott(cb2, 50, 0.3, 0.4, 0.2, 0.1, k)
	fmt.Println("Gray scott last: ",mycheck[len(mycheck)-1])

}
*/
package main

import (
	"fmt"
	"gifhelper"
)

func main() {
	//fmt.Println("hello!")
	numRows := 250
	numCols := 250

	initialBoard := InitializeBoard(numRows, numCols)

	frac := 0.05 // tells us what percentage of interior cells to color with predators

	// how many predator rows and columns are there?
	predRows := frac * float64(numRows)
	predCols := frac * float64(numCols)

	midRow := numRows / 2
	midCol := numCols / 2

	// a little for loop to fill predators
	for r := midRow - int(predRows/2); r < midRow+int(predRows/2); r++ {
		for c := midCol - int(predCols/2); c < midCol+int(predCols/2); c++ {
			initialBoard[r][c][1] = 1.0
		}
	}

	// make prey concentration 1 at every cell
	for i := range initialBoard {
		for j := range initialBoard[i] {
			initialBoard[i][j][0] = 1.0
		}
	}

	// let's set some parameters too
	numGens := 20000 // number of iterations
	feedRate := 0.034
	killRate := 0.095

	preyDiffusionRate := 0.2 // prey are twice as fast at running
	predatorDiffusionRate := 0.1

	// let's declare kernel
	var kernel [3][3]float64
	kernel[0][0] = .05
	kernel[0][1] = .2
	kernel[0][2] = .05
	kernel[1][0] = .2
	kernel[1][1] = -1.0
	kernel[1][2] = .2
	kernel[2][0] = .05
	kernel[2][1] = .2
	kernel[2][2] = .05

	//fmt.Println("Initial board", initialBoard)
	//fmt.Println("Kernel: ",kernel)
	// let's simulate Gray-Scott!
	// result will be a collection of Boards corresponding to each generation.
	boards := SimulateGrayScott(initialBoard, numGens, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)

	fmt.Println("Done with simulation!")

	// we will draw what we have generated.
	fmt.Println("Drawing boards to file.")

	//for the visualization, we are only going to draw every nth board to be more efficient
	n := 100

	cellWidth := 1 // each cell is 1 pixel

	imageList := DrawBoards(boards, cellWidth, n)
	fmt.Println("Boards drawn! Now draw GIF.")

	outFile := "Gray-Scott"
	gifhelper.ImagesToGIF(imageList, outFile) // code is given
	fmt.Println("GIF drawn!")
}
