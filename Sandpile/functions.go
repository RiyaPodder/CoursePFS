package main

import(
  "math/rand"
  "time"
  "runtime"
)

//Gameboard is a 2d matrix of ints for every board of the sandpile simulation
type Gameboard [][]int

//position contains information about each cell to be updated in every round of simulation, first value is row index, second value is column index and third value is the amount by which we have to update the current number in this cell position
type position [3]int

//SimulateSandpile is the highest level function of our serial code for simulating a sandpile
//Input: The initial gameboard
//Output: The final stable gameboard
func SimulateSandpile(initial_gameboard Gameboard)Gameboard{

  var final_configuration Gameboard

  for i := 0; i > -1; i++{

    next_gameboard := UpdateGameboard(initial_gameboard)

    if GameboardStable(next_gameboard){
      final_configuration = next_gameboard
      break
    }
    initial_gameboard = next_gameboard
  }

  return final_configuration
}

//InitializeGameboard initializes the first gameboard of this simulation according to the size, pile and placement value
//Input: Size of gameboard int, initial number of coins on the gameboard int and the string placement which defines how the initial gameboard is set
//Output: Gameboard
func InitializeGameboard(size, pile int, placement string)Gameboard{

  var initial_gameboard Gameboard

  initial_gameboard = make(Gameboard, size)
  for row := range initial_gameboard{
    initial_gameboard[row] = make([]int, size)
  }

  if placement == "central"{
    initial_gameboard[size/2][size/2] = pile

  }else if placement == "random"{
    //choose 100 random squares
    rand.Seed(time.Now().UnixNano())

    for n := 0; n<100; n++{
      r := rand.Intn(size)
      c := rand.Intn(size)
      initial_gameboard[r][c] += pile/100
    }
  }else{
    panic("Unacceptable placement value!")
  }

  return initial_gameboard
}

//GameboardsEqual checks if two gameboards have identical values in each of their squares or not
//Input: Two gameboards
//Output: True or false boolean value
func GameboardsEqual(board1, board2 Gameboard)bool{

  var flag bool
  flag = true

  if len(board1) != len(board2) || len(board1[0]) != len(board2[0]){
    flag = false
    return flag
  }

  for r := 0; r < len(board1); r++{
    for c := 0; c < len(board1[0]); c++{
      if board1[r][c] == board2[r][c]{
        continue
      }else{
        flag = false
        return flag
      }
    }
  }

  return flag
}

//UpdateGameboard updates the gameboard after one round of simulation
//Input: The current gameboard
//Output: The next gameboard
func UpdateGameboard(currentBoard Gameboard)Gameboard{

  var newBoard Gameboard

  newBoard = make(Gameboard, len(currentBoard))
  for row := range newBoard{
    newBoard[row] = make([]int, len(currentBoard[0]))
  }

  newBoard = CopyBoard(currentBoard, newBoard)
  size := len(newBoard)
  getPositions := make([]position,0)

  for r := 0; r < size; r++{
    for c := 0; c < size; c++{

      if newBoard[r][c] >= 4{
        getPositions = append(getPositions,Topple(r, c, size, 0)...)
      }

    }
  }
  for _,p := range getPositions{
    row_index := p[0]
    col_index := p[1]
    newBoard[row_index][col_index]+= p[2]
  }


  return newBoard
}

//CopyBoard copies the contents of current board to new board
//Input: Current board and new board
//Output: Changed new board
func CopyBoard(currentBoard Gameboard, newBoard Gameboard)Gameboard{

  size := len(currentBoard)

  for r := 0; r < size; r++{
    for c := 0; c < size; c++{

      newBoard[r][c] = currentBoard[r][c]
    }
  }

  return newBoard

}

//GameboardStable checks if the board has reached a stable configuration
//Input: Gameboard
//Output: True or false boolean
func GameboardStable(board Gameboard)bool{

  for r := range board{
    for c := range board[r]{
      //fmt.Println(c)
      if board[r][c] > 3{
        return false
      }
    }
  }
  return true
}

//InField checks if the input row or column index is within the gameboard boundaries
//Input: Row or column int, min and max size of gameboard int
//Output: True or false bool
func InField(index, start, stop int)bool{

  if index >= start && index < stop{
    return true
  }else{
    return false
  }
}

//Topple returns a slice of "positions" which is a slice of 3 ints, row index, column index and the value by which we need to update the number in that particular cell by in the current round of simulation
//Input: 4 ints, row and column index of current cell, max rows in the board, startIndex of board needed when a process has a subslice of the gameboard
//Output: Slice of positions for all cells to be updated
func Topple(row, col, size, startIndex int)[]position{

  sendPositions := make([]position,0)

  var trp position
  trp[0] = row + startIndex
  trp[1] = col
  trp[2] = -4
  sendPositions = append(sendPositions, trp)

  if InField(startIndex+row-1, 0, size){
    var pos position
    pos[0] = row-1+startIndex
    pos[1] = col
    pos[2] = 1
    sendPositions = append(sendPositions, pos)
  }
  if InField(startIndex+row+1, 0, size){
    var pos position
    pos[0] = row+1+startIndex
    pos[1] = col
    pos[2] = 1
    sendPositions = append(sendPositions, pos)
  }
  if InField(col-1, 0, size){
    var pos position
    pos[0] = row+startIndex
    pos[1] = col-1
    pos[2] = 1
    sendPositions = append(sendPositions, pos)

  }
  if InField(col+1, 0, size){
    var pos position
    pos[0] = row+startIndex
    pos[1] = col+1
    pos[2] = 1
    sendPositions = append(sendPositions, pos)
  }

  return sendPositions

}


//SimulateSandpileParallel is the highest level function of our parallel code for simulating a sandpile
//Input: The initial gameboard
//Output: The final stable gameboard
func SimulateSandpileParallel(initial_gameboard Gameboard)Gameboard{

    var final_configuration Gameboard
    numProcs := runtime.NumCPU()

    for i := 0; i > -1; i++{

      next_gameboard := SandpileMultiprocs(initial_gameboard, numProcs)

      if GameboardStable(next_gameboard){
        final_configuration = next_gameboard
        break
      }
      initial_gameboard = next_gameboard
    }

    return final_configuration
}

//SandpileMultiprocs updates the gameboard after one round of simulation
//Input: The current gameboard
//Output: The next gameboard
func SandpileMultiprocs(currentBoard Gameboard, numProcs int)Gameboard{

  size := len(currentBoard)

  var board Gameboard

  board = make(Gameboard, size)
  for row := range board{
    board[row] = make([]int, size)
  }

  board = CopyBoard(currentBoard, board)

  c := make(chan []position, numProcs)

  for i := 0; i < numProcs; i++{
    startIndex := i * (size/numProcs)
    endIndex := (i+1) * (size/numProcs)
    if i < numProcs-1{
      go SandpileSingleproc(board[startIndex:endIndex], size, startIndex, c)
    }else{
      go SandpileSingleproc(board[startIndex: ], size, startIndex, c)
    }
  }

  for i := 0; i < numProcs; i++{
    var updatePositions []position
    updatePositions = <- c
    for _,p := range updatePositions{
      row_index := p[0]
      col_index := p[1]
      board[row_index][col_index]+= p[2]
    }
  }

  return board
}

//SandpileSingleproc checks the cells where the coins topple for each subslice of the gameboard that is passed to the various processors
//Input: The current gameboard, max rows in the board, startIndex is the index of the first row of this subslice in the original board, and the buffered channel to pass the message, the cells that need to be updated
//Output: No output as it is a goroutine, instead channel is used for passing message between processors
func SandpileSingleproc(board Gameboard, size, startIndex int, c chan []position){

  getPositions := make([]position,0)

  for row := 0; row < len(board); row++{
    for col := 0; col < size; col++{

      if board[row][col] >= 4{

        getPositions = append(getPositions, Topple(row, col, size, startIndex)...)
      }

    }
  }

  c <- getPositions
}
