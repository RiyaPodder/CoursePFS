package main

import(
  "fmt"
)

//place your functions from the assignment here.

type GameBoard [][]float64

func DiffuseBoardOneParticle(currentBoard [][]float64, diffusionRate float64, kernel [3][3]float64) [][]float64 {

  numRows := CountRows1(currentBoard)
  numCols := CountCols1(currentBoard)
  newBoard := InitializeBoard1(numRows, numCols)

  for i:=0; i<numRows; i++{
    for j:=0; j<numCols; j++{
      mtx := MooreMatrix1(currentBoard,i,j) //gives a 3x3 matrix mtx with center cell currentBoard[i][j]
      change_in_val := ConvolveMatrix1(mtx,kernel) //gives change in value for this cell
      new_val := currentBoard[i][j] + change_in_val*diffusionRate
      newBoard[i][j] = new_val
    }
  }

  return newBoard

}

func CountRows1(board GameBoard) int{
  return len(board)
}

func CountCols1(board GameBoard) int{
  return len(board[0])
}

func InitializeBoard1(r,c int) GameBoard{
  n_board := make([][]float64,r)
  for rows := range n_board{
    n_board[rows] = make([]float64,c)
  }
  return n_board
}

func MooreMatrix1(board GameBoard, row, col int) [3][3]float64{
  moore_mtx := [3][3]float64{}

  for i:=row-1; i<=row+1; i++{
    for j:=col-1; j<=col+1; j++{
      if InField1(board,i,j){
        moore_mtx[i+1-row][j+1-col] = board[i][j]
      } else{
        moore_mtx[i+1-row][j+1-col] = 0
      }
    }
  }

  return moore_mtx
}

func InField1(board GameBoard, r,c int)bool{
  numRows := CountRows1(board)
  numCols := CountCols1(board)

  if r<0 || r>=numRows || c<0 || c>=numCols{
    return false
  }

  return true
}

func ConvolveMatrix1(matrix, kernel [3][3]float64) float64{

  sum := 0.0

  for i:=0; i<3; i++{
    for j:=0; j<3; j++{

      sum += matrix[i][j]*kernel[i][j]

    }
  }

  return sum
}


func SumCells(cells ...Cell) Cell {

  var sum_of_cells Cell

  for _,val := range cells{
    sum_of_cells[0] += val[0]
    sum_of_cells[1] += val[1]
  }

  return sum_of_cells

}

func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {

  var after_change_cell Cell

  reproduction_value := currentCell[0]*currentCell[1]*currentCell[1]

  after_change_cell[0] = feedRate*(1-currentCell[0]) - reproduction_value
  after_change_cell[1] = -killRate*currentCell[1] + reproduction_value

  return after_change_cell

}


func ChangeDueToDiffusion1(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {

// Disjoining the currentBoard to give 2 matrices, each containing cell concentration of only one type of molecule
numRows := len(currentBoard)
numCols := len(currentBoard[0])
matA := InitializeBoard1(numRows, numCols)
matB := InitializeBoard1(numRows, numCols)

for i := range currentBoard{
  for j:= range currentBoard[i]{
    matA[i][j] = currentBoard[i][j][0]
    matB[i][j] = currentBoard[i][j][1]
  }
}

mtxA := MooreMatrix1(matA,row,col)
change_in_val_A := ConvolveMatrix1(mtxA,kernel) * preyDiffusionRate

mtxB := MooreMatrix1(matB,row,col)
change_in_val_B := ConvolveMatrix1(mtxB,kernel) * predatorDiffusionRate

var diffusion_change Cell

diffusion_change[0] = change_in_val_A
diffusion_change[1] = change_in_val_B

return diffusion_change
}

/*
UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
    currentCell ← currentBoard[row][col]
    diffusionValues ← ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
    reactionValues ← ChangeDueToReactions(currentCell, feedRate, killRate)
    return SumCells(currentCell, diffusionValues, reactionValues)
*/

func UpdateCell(currentBoard Board, row, col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell{

  currentCell := currentBoard[row][col]
  diffusionValues := ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
  reactionValues := ChangeDueToReactions(currentCell, feedRate, killRate)

  updated_cell := SumCells(currentCell, diffusionValues, reactionValues)

  return updated_cell
}

/*
UpdateBoard(currentBoard, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	numRows ← CountRows(currentBoard)
	numCols ← CountColumns(currentBoard)
	newBoard ← InitializeBoard(numRows, numCols)
	for row ← 0 to numRows – 1
		for col ← 0 to numCols – 1
			newBoard[row][col] ← UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	return newBoard
*/

func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board{

  numRows := len(currentBoard)
  numCols := len(currentBoard[0])
  newBoard := InitializeBoard(numRows, numCols)

  for row:=0; row<numRows; row++{
    for col:=0; col<numCols; col++{
      newBoard[row][col] = UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
    }
  }
  return newBoard

}

func InitializeBoard(r, c int)Board{

  n_board := make(Board,r) //([][]Cell) ([][][2]float64)
  //fmt.Println(len(n_board))


  for row := range n_board{
    n_board[row] = make([]Cell,c) //([][2]float64)
  }
  //fmt.Println(len(n_board[0]))

  return n_board

}


/*
SimulateGrayScott(initialBoard, numGens, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	boards ← array of numGens + 1 Boards
	boards[0] ← initialBoard
	for i ← 1 to numGens
		boards[i] ← UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	return boards

  */


func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64)[]Board{

  boards := make([]Board, numGens+1)
  boards[0] = initialBoard

  for i:=1; i<=numGens; i++{
    fmt.Println("Generation started: ",i)
    boards[i] = UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
    fmt.Println("Generation finished: ",i)
  }

  return boards

}

func ChangeDueToDiffusion(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {


mtx := MooreMatrix(currentBoard,row,col)

var diffusion_change Cell
diffusion_change = ConvolveMatrix(mtx,kernel)
diffusion_change[0] *= preyDiffusionRate
diffusion_change[1] *= predatorDiffusionRate


return diffusion_change
}

func MooreMatrix(board Board, row, col int) [3][3][2]float64{
  moore_mtx := [3][3][2]float64{}

  for i:=row-1; i<=row+1; i++{
    for j:=col-1; j<=col+1; j++{
      if InField(board,i,j){
        moore_mtx[i+1-row][j+1-col][0] = board[i][j][0]
        moore_mtx[i+1-row][j+1-col][1] = board[i][j][1]
      } else{
        moore_mtx[i+1-row][j+1-col][0] = 0
        moore_mtx[i+1-row][j+1-col][1] = 0
      }
    }
  }

  return moore_mtx
}

func InField(board Board, r,c int)bool{
  numRows := CountRows(board)
  numCols := CountCols(board)

  if r<0 || r>=numRows || c<0 || c>=numCols{
    return false
  }

  return true
}


func CountRows(board Board) int{
  return len(board)
}

func CountCols(board Board) int{
  return len(board[0])
}

func ConvolveMatrix(matrix [3][3][2]float64, kernel [3][3]float64) Cell{

  var sum Cell

  for i:=0; i<3; i++{
    for j:=0; j<3; j++{

      sum[0] += matrix[i][j][0]*kernel[i][j]
      sum[1] += matrix[i][j][1]*kernel[i][j]

    }
  }

  return sum
}
